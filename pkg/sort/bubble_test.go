package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shkov/sort-visualization/pkg/iteration"
)

func TestBubbleSorter_Step(t *testing.T) {
	testCases := []struct {
		name    string
		arr     []int
		wantArr []int
	}{
		{
			name:    "normal response",
			arr:     []int{7, 4, 3, 6, 5, 2, 1},
			wantArr: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:    "reverse",
			arr:     []int{7, 6, 5, 4, 3, 2, 1},
			wantArr: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:    "already sorted",
			arr:     []int{1, 2, 3, 4, 5, 6, 7},
			wantArr: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewBubbleSorter(tc.arr)

			for {
				_, ok := s.Step()
				if !ok {
					break
				}
			}

			gotArr := getArr(s.Dump())

			assert.Equal(t, tc.wantArr, gotArr)
		})
	}
}

func getArr(iter *iteration.ArrayIterator) []int {
	arr := make([]int, 0)
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}
		arr = append(arr, item)
	}

	return arr
}
