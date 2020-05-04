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
	AlgorithmBubble Algorithm = "bubble"
)

func ParseAlgorithm(val string) (Algorithm, error) {
	switch Algorithm(val) {
	case AlgorithmBubble:
		return AlgorithmBubble, nil
	}

	return "", errInvalidAlgorithm
}

type Sorter interface {
	Step() bool
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

	default:
		return nil, errInvalidAlgorithm
	}
}
