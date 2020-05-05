package sort

import (
	"errors"

	"github.com/shkov/sort-visualization/pkg/iteration"
)

var (
	errInvalidAlgorithm = errors.New("invalid algorithm type")
)

type Algorithm string

const (
	AlgorithmBubble    Algorithm = "bubble"
	AlgorithmSelection Algorithm = "selection"
	AlgorithmInsertion Algorithm = "insertion"
)

func ParseAlgorithm(val string) (Algorithm, error) {
	switch Algorithm(val) {
	case AlgorithmBubble:
		return AlgorithmBubble, nil
	case AlgorithmSelection:
		return AlgorithmSelection, nil
	case AlgorithmInsertion:
		return AlgorithmInsertion, nil
	}

	return "", errInvalidAlgorithm
}

type Sorter interface {
	Step() (*iteration.Stat, bool)
	Shuffle()
	Dump() *iteration.ArrayIterator
	String() string
}

func NewSorter(arr []int, algorithmStr string) (Sorter, error) {
	algorithm, err := ParseAlgorithm(algorithmStr)
	if err != nil {
		return nil, err
	}

	switch algorithm {
	case AlgorithmBubble:
		return NewBubbleSorter(arr), nil

	case AlgorithmSelection:
		return NewSelectionSorter(arr), nil

	case AlgorithmInsertion:
		return NewInsertionSorter(arr), nil

	default:
		return nil, errInvalidAlgorithm
	}
}
