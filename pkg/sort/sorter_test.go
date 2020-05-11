package sort

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shkov/sort-visualization/pkg/iteration"
)

func TestBubbleSorter_Step(t *testing.T) {
	testCases := []struct {
		name      string
		arr       []int
		algorithm Algorithm
	}{
		{
			name:      "AlgorithmBubble",
			arr:       makeArr(t),
			algorithm: AlgorithmBubble,
		},
		{
			name:      "AlgorithmInsertion",
			arr:       makeArr(t),
			algorithm: AlgorithmInsertion,
		},
		{
			name:      "AlgorithmSelection",
			arr:       makeArr(t),
			algorithm: AlgorithmSelection,
		},
		{
			name:      "AlgorithmMerge",
			arr:       makeArr(t),
			algorithm: AlgorithmMerge,
		},
	}

	for _, tc := range testCases {
		copied := make([]int, len(tc.arr))
		copy(copied, tc.arr)
		sort.Slice(copied, func(i, j int) bool { return copied[i] < copied[j] })

		t.Run(tc.name, func(t *testing.T) {
			s, err := NewSorter(tc.arr, string(tc.algorithm))
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}

			for {
				_, ok := s.Step()
				if !ok {
					break
				}
			}

			gotArr := getArr(s.Dump())

			assert.Equal(t, copied, gotArr)
		})
	}
}

func makeArr(_ *testing.T) []int {
	maxLen := rand.Intn(1000)

	arr := make([]int, maxLen)
	for i := 0; i < maxLen; i++ {
		arr[i] = rand.Int()
	}

	return arr
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
