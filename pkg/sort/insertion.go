package sort

import (
	"math/rand"

	"github.com/shkov/sort-visualization/pkg/iteration"
)

type InsertionSorter struct {
	arr  []int
	stat iteration.Stat

	maxSortedIdx int
}

func NewInsertionSorter(arr []int) Sorter {
	return &InsertionSorter{
		arr:          arr,
		stat:         iteration.Stat{},
		maxSortedIdx: 1,
	}
}

func (is *InsertionSorter) Step() (*iteration.Stat, bool) {
	if is.maxSortedIdx == len(is.arr) {
		return nil, false
	}

	j := is.maxSortedIdx
	for j > 0 {
		if is.arr[j-1] > is.arr[j] {
			is.arr[j-1], is.arr[j] = is.arr[j], is.arr[j-1]
		}

		j--
	}

	is.maxSortedIdx++

	return &is.stat, true
}

func (is *InsertionSorter) Shuffle() {
	rand.Shuffle(len(is.arr), func(i, j int) { is.arr[i], is.arr[j] = is.arr[j], is.arr[i] })
	is.stat = iteration.Stat{}
	is.maxSortedIdx = 1
}

func (is *InsertionSorter) Dump() *iteration.ArrayIterator {
	return iteration.NewArrayIterator(is.arr)
}

func (is *InsertionSorter) String() string {
	return "insertion"
}

func insertionsort(items []int) {
	var n = len(items)

	for i := 1; i < n; i++ {

		j := i
		for j > 0 {

			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}

			j--
		}
	}
}
