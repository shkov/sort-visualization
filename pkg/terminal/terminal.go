package terminal

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/shkov/sort-visualization/pkg/iteration"
)

type Config struct {
	Sorter         Sorter
	RefreshTimeout time.Duration
}

func (cfg Config) Validate() error {
	if cfg.Sorter == nil {
		return errors.New("must provide Sorter")
	}
	if cfg.RefreshTimeout <= 0 {
		return errors.New("must provide RefreshTimeout")
	}

	return nil
}

type Sorter interface {
	Shuffle()
	Step() bool
	Dump() *iteration.ArrayIterator
	String() string
}

type state int

const (
	stateWaiting state = iota
	stateRunning
	stateClosed
)

type Terminal struct {
	state          state
	barChart       *widgets.BarChart
	sorter         Sorter
	refreshTimeout time.Duration
	wg             sync.WaitGroup
	closeCh        chan struct{}
}

func New(cfg Config) (*Terminal, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}

	bc := widgets.NewBarChart()
	bc.Title = strings.Join([]string{cfg.Sorter.String(), "sort"}, " ")
	bc.SetRect(5, 5, 190, 56)
	bc.BarWidth = 2
	bc.BarGap = 2
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	t := &Terminal{
		state:          stateWaiting,
		barChart:       bc,
		sorter:         cfg.Sorter,
		refreshTimeout: cfg.RefreshTimeout,
		wg:             sync.WaitGroup{},
		closeCh:        make(chan struct{}),
	}

	return t, nil
}

func (t *Terminal) RunWidget() error {
	if t.state != stateWaiting {
		return errors.New("invalid terminal state")
	}

	err := ui.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize term ui: %w", err)
	}

	dataset := t.getDataset(t.sorter.Dump())
	colors := t.getColors(dataset, []int{})
	t.barChart.Data = dataset
	t.barChart.BarColors = colors
	ui.Render(t.barChart)

	// Start bar chard background updating.
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()

		ticker := time.NewTicker(t.refreshTimeout)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				t.renderBadChart()

			case <-t.closeCh:
				return
			}
		}
	}()

	t.state = stateRunning

	return nil
}

func (t *Terminal) renderBadChart() {
	ok := t.sorter.Step()
	if !ok {
		t.sorter.Shuffle()
		return
	}

	iter := t.sorter.Dump()
	dataset := t.getDataset(iter)

	// Find updated values.
	updatedIdxs := make([]int, 0)
	for i := range dataset {
		if dataset[i] != t.barChart.Data[i] {
			updatedIdxs = append(updatedIdxs, i)
		}
	}

	if len(updatedIdxs) == 0 {
		return
	}

	colors := t.getColors(dataset, updatedIdxs)

	t.barChart.Data = dataset
	t.barChart.BarColors = colors
	ui.Render(t.barChart)
}

func (t *Terminal) getColors(dataset []float64, updatedIdxs []int) []ui.Color {
	// isUpdated checks if index is in the updated index list.
	isUpdated := func(i int) bool {
		for _, updatedIdx := range updatedIdxs {
			if i == updatedIdx {
				return true
			}
		}

		return false
	}

	// If value was updated it will have another color.
	colors := make([]ui.Color, 0, len(dataset))
	for i := range dataset {
		if isUpdated(i) {
			colors = append(colors, ui.ColorRed)
			continue
		}

		colors = append(colors, ui.ColorWhite)
	}

	return colors
}

func (t *Terminal) getDataset(iter *iteration.ArrayIterator) []float64 {
	dataset := make([]float64, 0)
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}
		dataset = append(dataset, float64(item))
	}

	return dataset
}

func (t *Terminal) WaitExitSignal() {
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

func (t *Terminal) Close() {
	if t.state != stateRunning {
		return
	}

	close(t.closeCh)
	t.wg.Wait()
	ui.Close()

	t.state = stateClosed
}
