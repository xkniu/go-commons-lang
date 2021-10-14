package iterator

import errors "errors"

// A Iterator to iterate through multiple iterators one by one.
type chainedIterator struct {
	iterators []Iterator
	cursor    int
}

func (t *chainedIterator) HasNext() bool {
	for ; t.cursor < len(t.iterators); {
		curIter := t.iterators[t.cursor]
		if curIter.HasNext() {
			return true
		}
		t.cursor++
	}
	return false
}

func (t *chainedIterator) Next() interface{} {
	if !t.HasNext() {
		panic(errors.New("no more elements"))
	}
	return t.iterators[t.cursor].Next()
}
