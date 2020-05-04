package iteration

type ArrayIterator struct {
	arr     []int
	current int
}

func NewArrayIterator(a []int) *ArrayIterator {
	return &ArrayIterator{arr: a, current: 0}
}

func (iter *ArrayIterator) Reset() {
	iter.current = 0
}

func (iter *ArrayIterator) Next() (int, bool) {
	if iter.current > len(iter.arr)-1 {
		return 0, false
	}

	item := iter.arr[iter.current]
	iter.current++

	return item, true
}
