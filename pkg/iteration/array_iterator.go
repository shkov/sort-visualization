package iteration

import (
	"sync/atomic"
)

type ArrayIterator struct {
	arr     []int
	nextIdx int
}

func NewArrayIterator(a []int) *ArrayIterator {
	return &ArrayIterator{arr: a, nextIdx: 0}
}

func (iter *ArrayIterator) Reset() {
	iter.nextIdx = 0
}

func (iter *ArrayIterator) Next() (int, bool) {
	if iter.nextIdx > len(iter.arr)-1 {
		return 0, false
	}

	item := iter.arr[iter.nextIdx]
	iter.nextIdx++

	return item, true
}

type Stat struct {
	comparisonsTotal   uint64
	arrayAccessesTotal uint64
}

func (s *Stat) Reset() {
	atomic.StoreUint64(&s.comparisonsTotal, 0)
	atomic.StoreUint64(&s.arrayAccessesTotal, 0)
}

func (s *Stat) GetComparisonsTotal() uint64 {
	return atomic.LoadUint64(&s.comparisonsTotal)
}

func (s *Stat) GetArrayAccessesTotal() uint64 {
	return atomic.LoadUint64(&s.arrayAccessesTotal)
}

func (s *Stat) OnComparison() {
	atomic.AddUint64(&s.comparisonsTotal, 1)
}

func (s *Stat) OnArrayAccess() {
	atomic.AddUint64(&s.arrayAccessesTotal, 1)
}
