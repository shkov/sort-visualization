package sort

import (
	"math/rand"

	"github.com/shkov/sort-visualization/pkg/iteration"
)

type SelectionSorter struct {
	arr  []int
	stat iteration.Stat

	minNotSortedIdx int
}

func NewSelectionSorter(arr []int) Sorter {
	return &SelectionSorter{
		arr:             arr,
		stat:            iteration.Stat{},
		minNotSortedIdx: 0,
	}
}

func (ss *SelectionSorter) Step() (*iteration.Stat, bool) {
	if ss.minNotSortedIdx > len(ss.arr)-1 {
		return nil, false
	}

	minValueIdx := ss.minNotSortedIdx
	for i := ss.minNotSortedIdx; i < len(ss.arr); i++ {
		ss.stat.OnComparison()
		if ss.arr[i] < ss.arr[minValueIdx] {
			ss.stat.OnArrayAccess()
			minValueIdx = i
		}
	}

	ss.stat.OnComparison()
	if ss.arr[minValueIdx] < ss.arr[ss.minNotSortedIdx] {
		ss.stat.OnArrayAccess()
		ss.stat.OnArrayAccess()
		ss.arr[minValueIdx], ss.arr[ss.minNotSortedIdx] = ss.arr[ss.minNotSortedIdx], ss.arr[minValueIdx]
	}

	ss.minNotSortedIdx++

	return &ss.stat, true
}

func (ss *SelectionSorter) Shuffle() {
	rand.Shuffle(len(ss.arr), func(i, j int) { ss.arr[i], ss.arr[j] = ss.arr[j], ss.arr[i] })
	ss.stat = iteration.Stat{}
	ss.minNotSortedIdx = 0
}

func (ss *SelectionSorter) Dump() *iteration.ArrayIterator {
	return iteration.NewArrayIterator(ss.arr)
}

func (ss *SelectionSorter) String() string {
	return "selection"
}
