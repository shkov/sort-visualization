package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSorter_Step(t *testing.T) {
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
			s := NewSelectionSorter(tc.arr)

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
