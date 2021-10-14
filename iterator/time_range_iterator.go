package iterator

import (
	"errors"
	"time"
)

// ---------------------------------------------------------------------
// -- A time range iterator impl
// An Iterator to get next element in time range in sequence.
type timeRangeIterator struct {
	startInclusive time.Time // first element
	endExclusive   time.Time
	nextFunc       func(tm time.Time) time.Time // func to cal next element
	cur            time.Time                    // require initialized
}

func (t *timeRangeIterator) Init() {
	t.cur = t.startInclusive
}

func (t *timeRangeIterator) HasNext() bool {
	return t.cur.Before(t.endExclusive)
}

func (t *timeRangeIterator) Next() interface{} {
	if !t.cur.Before(t.endExclusive) {
		panic(errors.New("no more elements"))
	}
	r := t.cur
	t.cur = t.nextFunc(t.cur)
	return r
}
