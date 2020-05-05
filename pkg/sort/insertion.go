package sort

import (
	"math/rand"

	"github.com/shkov/sort-visualization/pkg/iteration"
)

type InsertionSorter struct {
	arr  []int
	stat iteration.Stat

	maxSortedIdx  int
	processingIdx int
}

func NewInsertionSorter(arr []int) Sorter {
	return &InsertionSorter{
		arr:           arr,
		stat:          iteration.Stat{},
		maxSortedIdx:  1,
		processingIdx: 1,
	}
}

func (is *InsertionSorter) Step() (*iteration.Stat, bool) {
	if is.processingIdx < 1 {
		is.maxSortedIdx++
		is.processingIdx = is.maxSortedIdx
	}

	if is.maxSortedIdx == len(is.arr) {
		return nil, false
	}

	for is.processingIdx > 0 {
		is.stat.OnComparison()
		is.stat.OnArrayAccess()
		is.stat.OnArrayAccess()
		if is.arr[is.processingIdx-1] > is.arr[is.processingIdx] {
			is.stat.OnArrayAccess()
			is.stat.OnArrayAccess()
			is.arr[is.processingIdx-1], is.arr[is.processingIdx] = is.arr[is.processingIdx], is.arr[is.processingIdx-1]
			break
		}

		is.processingIdx--
	}

	return &is.stat, true
}

func (is *InsertionSorter) Shuffle() {
	rand.Shuffle(len(is.arr), func(i, j int) { is.arr[i], is.arr[j] = is.arr[j], is.arr[i] })
	is.stat = iteration.Stat{}
	is.maxSortedIdx = 1
	is.processingIdx = 1
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
