package sort

import (
	"math/rand"

	"github.com/shkov/sort-visualization/pkg/iteration"
)

type BubbleSorter struct {
	arr  []int
	stat iteration.Stat
	i    int
	j    int
}

func NewBubbleSorter(arr []int) Sorter {
	return &BubbleSorter{
		arr:  arr,
		stat: iteration.Stat{},
		i:    0,
		j:    len(arr) - 1,
	}
}

func (bs *BubbleSorter) String() string {
	return "bubble"
}

func (bs *BubbleSorter) Step() (*iteration.Stat, bool) {
	if bs.j == 0 {
		if bs.i == len(bs.arr)-1 {
			return nil, false
		}

		bs.i++
		bs.j = len(bs.arr) - 1
	}

	bs.stat.OnComparison()
	if bs.arr[bs.j] < bs.arr[bs.j-1] {
		bs.stat.OnArrayAccess()
		bs.stat.OnArrayAccess()
		bs.stat.OnArrayAccess()
		bs.stat.OnArrayAccess()
		bs.arr[bs.j], bs.arr[bs.j-1] = bs.arr[bs.j-1], bs.arr[bs.j]
	}

	bs.j--

	return &bs.stat, true
}

func (bs *BubbleSorter) Dump() *iteration.ArrayIterator {
	return iteration.NewArrayIterator(bs.arr)
}

func (bs *BubbleSorter) Shuffle() {
	rand.Shuffle(len(bs.arr), func(i, j int) { bs.arr[i], bs.arr[j] = bs.arr[j], bs.arr[i] })
	bs.i = 0
	bs.j = len(bs.arr) - 1
	bs.stat = iteration.Stat{}
}
