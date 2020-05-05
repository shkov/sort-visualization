package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSorter_Step(t *testing.T) {
	testCases := []struct {
		name    string
		arr     []int
		wantArr []int
	}{
		{
			name:    "normal response",
			arr:     []int{44, 28, 43, 25, 1, 11, 29, 40, 30, 26, 18, 16, 12, 8, 19, 17, 4, 13, 20, 24, 10, 7, 27, 6, 21, 41, 9, 3, 5, 2, 45, 14, 15, 23, 42, 22},
			wantArr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 40, 41, 42, 43, 44, 45},
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
			s := NewInsertionSorter(tc.arr)

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
