package sort

import (
	"math/rand"

	"github.com/shkov/sort-visualization/pkg/iteration"
)

type BubbleSorter struct {
	arr  []int
	stat iteration.Stat

	maxNotSortedIdx int
	maxValueIdx     int
}

func NewBubbleSorter(arr []int) Sorter {
	return &BubbleSorter{
		arr:             arr,
		stat:            iteration.Stat{},
		maxNotSortedIdx: len(arr) - 1,
		maxValueIdx:     0,
	}
}

func (bs *BubbleSorter) String() string {
	return "bubble"
}

func (bs *BubbleSorter) Step() (*iteration.Stat, bool) {
	if bs.maxNotSortedIdx < 0 {
		return nil, false
	}

	if bs.maxValueIdx == 0 {
		for i := 0; i < bs.maxNotSortedIdx; i++ {
			bs.stat.OnComparison()
			if bs.arr[i] > bs.arr[bs.maxValueIdx] {
				bs.stat.OnArrayAccess()
				bs.maxValueIdx = i
			}
		}
	}

	bs.stat.OnComparison()
	if bs.maxValueIdx != bs.maxNotSortedIdx && bs.arr[bs.maxValueIdx] > bs.arr[bs.maxNotSortedIdx] {
		bs.stat.OnArrayAccess()
		bs.stat.OnArrayAccess()
		bs.arr[bs.maxValueIdx], bs.arr[bs.maxValueIdx+1] = bs.arr[bs.maxValueIdx+1], bs.arr[bs.maxValueIdx]
		bs.maxValueIdx++
	} else {
		bs.maxValueIdx = 0
		bs.maxNotSortedIdx--
	}

	return &bs.stat, true
}

func (bs *BubbleSorter) Dump() *iteration.ArrayIterator {
	return iteration.NewArrayIterator(bs.arr)
}

func (bs *BubbleSorter) Shuffle() {
	rand.Shuffle(len(bs.arr), func(i, j int) { bs.arr[i], bs.arr[j] = bs.arr[j], bs.arr[i] })
	bs.stat = iteration.Stat{}
	bs.maxNotSortedIdx = len(bs.arr) - 1
	bs.maxValueIdx = 0
}
