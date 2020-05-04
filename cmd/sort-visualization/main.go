package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/shkov/sort-visualization/pkg/sort"
	"github.com/shkov/sort-visualization/pkg/terminal"
)

const (
	maxSortItem = 45
)

type configuration struct {
	TerminalRefreshInterval time.Duration
	SortingAlgorithm        string
}

func parseConfiguration() (*configuration, error) {
	terminalRefreshInterval := flag.Duration("refresh", 5*time.Millisecond, "bar chart refresh interval")
	sortingAlgorithm := flag.String("algorithm", "bubble", "sorting algorithm type")
	flag.Parse()

	return &configuration{
		TerminalRefreshInterval: *terminalRefreshInterval,
		SortingAlgorithm:        *sortingAlgorithm,
	}, nil
}

func main() {
	cfg, err := parseConfiguration()
	if err != nil {
		log.Fatalf("failed to parse configuration: %v", err)
	}

	dataset := make([]int, 0, 30)
	for i := 1; i <= maxSortItem; i++ {
		dataset = append(dataset, i)
	}

	rand.Shuffle(len(dataset), func(i, j int) { dataset[i], dataset[j] = dataset[j], dataset[i] })

	sorter, err := sort.NewSorter(dataset, cfg.SortingAlgorithm)
	if err != nil {
		log.Fatalf("failed to initialize sorter: %v", err)
	}

	t, err := terminal.New(terminal.Config{
		Sorter:         sorter,
		RefreshTimeout: cfg.TerminalRefreshInterval,
	})
	if err != nil {
		log.Fatalf("failed to initialize terminal: %v", err)
	}

	err = t.RunWidget()
	if err != nil {
		log.Fatalf("failed to run widget: %v", err)
	}

	t.WaitExitSignal()

	t.Close()
}
