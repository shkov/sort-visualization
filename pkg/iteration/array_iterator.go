package iteration

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
