package iteration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayIterator_Next(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		nextIdx  int
		wantNext int
		wantResp bool
	}{
		{
			name:     "normal response",
			arr:      []int{1, 2, 3, 4, 5, 6},
			nextIdx:  0,
			wantNext: 1,
			wantResp: true,
		},
		{
			name:     "normal response",
			arr:      []int{1, 2, 3, 4, 5, 6},
			nextIdx:  3,
			wantNext: 4,
			wantResp: true,
		},
		{
			name:     "no next item",
			arr:      []int{1, 2, 3, 4, 5, 6},
			nextIdx:  6,
			wantNext: 0,
			wantResp: false,
		},
		{
			name:     "empty array",
			arr:      []int{},
			nextIdx:  0,
			wantNext: 0,
			wantResp: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			iter := &ArrayIterator{
				arr:     tc.arr,
				nextIdx: tc.nextIdx,
			}

			gotItem, gotOk := iter.Next()

			assert.Equal(t, tc.wantNext, gotItem)
			assert.Equal(t, tc.wantResp, gotOk)
		})
	}
}

func TestArrayIterator_Reset(t *testing.T) {
	iter := NewArrayIterator([]int{1, 2, 3})

	item, _ := iter.Next()
	if item != 1 {
		t.Fatal("not first item returned")
	}

	iter.Reset()

	item, _ = iter.Next()
	if item != 1 {
		t.Fatal("not first item returned")
	}
}
