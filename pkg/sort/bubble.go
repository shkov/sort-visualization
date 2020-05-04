package sort

import (
	"math/rand"

	"github.com/shkov/sort-visualization/pkg/iteration"
)

type BubbleSorter struct {
	arr []int

	i int
	j int
}

func NewBubbleSorter(arr []int) Sorter {
	return &BubbleSorter{
		arr: arr,
		i:   0,
		j:   len(arr) - 1,
	}
}

func (bs *BubbleSorter) String() string {
	return "bubble"
}

func (bs *BubbleSorter) Step() bool {
	if bs.j == 0 {
		if bs.i == len(bs.arr)-1 {
			return false
		}

		bs.i++
		bs.j = len(bs.arr) - 1
	}

	if bs.arr[bs.j] < bs.arr[bs.j-1] {
		bs.arr[bs.j], bs.arr[bs.j-1] = bs.arr[bs.j-1], bs.arr[bs.j]
	}

	bs.j--

	return true
}

func (bs *BubbleSorter) Dump() *iteration.ArrayIterator {
	return iteration.NewArrayIterator(bs.arr)
}

func (bs *BubbleSorter) Shuffle() {
	rand.Shuffle(len(bs.arr), func(i, j int) { bs.arr[i], bs.arr[j] = bs.arr[j], bs.arr[i] })
	bs.i = 0
	bs.j = len(bs.arr) - 1
}

//func bubbleSort(tosort []int) {
//	size := len(tosort)
//	if size < 2 {
//		return
//	}
//
//	for i := 0; i < size; i++ {
//		for j := size - 1; j >= i+1; j-- {
//			if tosort[j] < tosort[j-1] {
//				tosort[j], tosort[j-1] = tosort[j-1], tosort[j]
//			}
//		}
//	}
//}
