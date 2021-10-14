package iterator

import "errors"

// A Iterator based on the slice.
type sliceIterator struct {
	elements []interface{}
	cursor   int
}

func (t *sliceIterator) HasNext() bool {
	return t.cursor < len(t.elements)
}

func (t *sliceIterator) Next() interface{} {
	i := t.cursor
	if i >= len(t.elements) {
		panic(errors.New("no more elements"))
	}
	t.cursor++
	return t.elements[i]
}