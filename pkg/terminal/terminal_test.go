package terminal

import (
	"errors"
	"testing"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/stretchr/testify/assert"

	"github.com/shkov/sort-visualization/pkg/iteration"
)

type mockSorter struct {
	onStep   func() (*iteration.Stat, bool)
	onDump   func() *iteration.ArrayIterator
	onString func() string
}

func (m mockSorter) Shuffle() {
}

func (m mockSorter) Step() (*iteration.Stat, bool) {
	return m.onStep()
}

func (m mockSorter) Dump() *iteration.ArrayIterator {
	return m.onDump()
}

func (m mockSorter) String() string {
	return m.onString()
}

func TestConfig_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		cfg     Config
		wantErr error
	}{
		{
			name:    "normal response",
			cfg:     *makeConfig(t, nil),
			wantErr: nil,
		},
		{
			name: "invalid Sorter",
			cfg: *makeConfig(t, func(cfg *Config) {
				cfg.Sorter = nil
			}),
			wantErr: errors.New("must provide Sorter"),
		},
		{
			name: "invalid RefreshTimeout",
			cfg: *makeConfig(t, func(cfg *Config) {
				cfg.RefreshTimeout = 0
			}),
			wantErr: errors.New("must provide RefreshTimeout"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotErr := tc.cfg.Validate()
			assert.Equal(t, tc.wantErr, gotErr)
		})
	}
}

func TestNew(t *testing.T) {
	terminal, err := New(*makeConfig(t, nil))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	terminal.Close()

	// check for deadlock
	terminal, err = New(*makeConfig(t, nil))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = terminal.RunWidget()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	terminal.Close()
}

func TestTerminal_renderBadChart(t *testing.T) {
	initialData := []float64{1, 2}
	initialColors := []ui.Color{ui.ColorBlue, ui.ColorGreen}

	testCases := []struct {
		name       string
		gotStep    bool
		newDataset []int
		wantColors []ui.Color
		wantData   []float64
	}{
		{
			name:       "normal response: changed all",
			gotStep:    true,
			newDataset: []int{3, 4},
			wantColors: []ui.Color{ui.ColorRed, ui.ColorRed},
			wantData:   []float64{3, 4},
		},
		{
			name:       "no changes",
			gotStep:    true,
			newDataset: []int{1, 2},
			wantColors: []ui.Color{ui.ColorBlue, ui.ColorGreen},
			wantData:   []float64{1, 2},
		},
		{
			name:       "only first change",
			gotStep:    true,
			newDataset: []int{7, 2},
			wantColors: []ui.Color{ui.ColorRed, ui.ColorWhite},
			wantData:   []float64{7, 2},
		},
		{
			name:       "only second change",
			gotStep:    true,
			newDataset: []int{1, 5},
			wantColors: []ui.Color{ui.ColorWhite, ui.ColorRed},
			wantData:   []float64{1, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bc := widgets.NewBarChart()
			bc.BarColors = initialColors
			bc.Data = initialData

			terminal := &Terminal{
				state:    stateRunning,
				barChart: bc,
				sorter: mockSorter{
					onStep: func() (*iteration.Stat, bool) {
						return &iteration.Stat{}, tc.gotStep
					},
					onDump: func() *iteration.ArrayIterator {
						return iteration.NewArrayIterator(tc.newDataset)
					},
					onString: func() string {
						return "testSorter"
					},
				},
			}

			terminal.renderBadChart()

			assert.Equal(t, tc.wantData, terminal.barChart.Data)
			assert.Equal(t, tc.wantColors, terminal.barChart.BarColors)
		})
	}
}

func makeConfig(_ *testing.T, fn func(cfg *Config)) *Config {
	cfg := &Config{
		Sorter: mockSorter{
			onStep: func() (*iteration.Stat, bool) {
				return &iteration.Stat{}, true
			},
			onDump: func() *iteration.ArrayIterator {
				return &iteration.ArrayIterator{}
			},
			onString: func() string {
				return "testSorter"
			},
		},
		RefreshTimeout: time.Second,
	}

	if fn != nil {
		fn(cfg)
	}

	return cfg
}
