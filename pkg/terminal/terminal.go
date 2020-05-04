package terminal

import (
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
	return nil
}

type Sorter interface {
	Shuffle()
	Step() bool
	Dump() *iteration.ArrayIterator
	String() string
}

type Terminal struct {
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
		barChart:       bc,
		sorter:         cfg.Sorter,
		refreshTimeout: cfg.RefreshTimeout,
		wg:             sync.WaitGroup{},
		closeCh:        make(chan struct{}),
	}

	return t, nil
}

func (t *Terminal) RunWidget() error {
	err := ui.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize termui: %w", err)
	}

	t.drawBarChart()

	return nil
}

func (t *Terminal) drawBarChart() {
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()

		ticker := time.NewTicker(t.refreshTimeout)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				ok := t.sorter.Step()
				if !ok {
					t.sorter.Shuffle()
					break
				}

				iter := t.sorter.Dump()
				dataset, colors, ok := t.getDataset(iter)
				if !ok {
					break
				}

				t.barChart.Data = dataset
				t.barChart.BarColors = colors
				ui.Render(t.barChart)

			case <-t.closeCh:
				return
			}
		}
	}()
}

func (t *Terminal) getDataset(iter *iteration.ArrayIterator) ([]float64, []ui.Color, bool) {
	dataset := make([]float64, 0)
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}
		dataset = append(dataset, float64(item))
	}

	if len(t.barChart.Data) == 0 {
		return dataset, []ui.Color{ui.ColorWhite}, true
	}

	updatedIdxs := make([]int, 0)
	for i := range dataset {
		if dataset[i] != t.barChart.Data[i] {
			updatedIdxs = append(updatedIdxs, i)
		}
	}
	if len(updatedIdxs) == 0 {
		return nil, nil, false
	}

	colors := make([]ui.Color, 0, len(dataset))
	for i := range dataset {
		for _, updatedIdx := range updatedIdxs {
			if i == updatedIdx {
				colors = append(colors, ui.ColorRed)
				break
			}
		}
		colors = append(colors, ui.ColorWhite)
	}

	return dataset, colors, true
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

func (t *Terminal) Cancel() {
	close(t.closeCh)
	t.wg.Wait()
	ui.Close()
}
