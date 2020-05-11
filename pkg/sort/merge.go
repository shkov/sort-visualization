package sort

import (
	"math/rand"

	"github.com/shkov/sort-visualization/pkg/iteration"
)

type MergeSorter struct {
	arr  []int
	stat iteration.Stat
	c    int
}

func NewMergeSorter(arr []int) Sorter {
	return &MergeSorter{
		arr: arr,
	}
}

func (s *MergeSorter) Step() (*iteration.Stat, bool) {
	if s.c > 0 {
		return nil, false
	}

	s.arr = mergeSort(s.arr)
	s.c++

	return &s.stat, true
}

func (s *MergeSorter) Shuffle() {
	rand.Shuffle(len(s.arr), func(i, j int) { s.arr[i], s.arr[j] = s.arr[j], s.arr[i] })
}

func (s *MergeSorter) Dump() *iteration.ArrayIterator {
	return iteration.NewArrayIterator(s.arr)
}

func (s *MergeSorter) String() string {
	return "merge"
}

func mergeSort(arr []int) []int {
	var left, right []int

	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2

	left = make([]int, 0, middle)
	for i := 0; i < middle; i++ {
		left = append(left, arr[i])
	}

	right = make([]int, 0, len(arr)-middle)
	for i := middle; i < len(arr); i++ {
		right = append(right, arr[i])
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}

/*
function merge(left,right)
    var list result
    while length(left) > 0 and length(right) > 0
        if first(left) ≤ first(right)
            append first(left) to result
            left = rest(left)
        else
            append first(right) to result
            right = rest(right)
        end if
    while length(left) > 0
        append first(left) to result
        left = rest(left)
    while length(right) > 0
        append first(right) to result
        right = rest(right)
    return result
*/
