package iterator

import (
	"github.com/xkniu/go-commons-lang/commutil"
	"github.com/xkniu/go-commons-lang/core"
	"time"
)

var Iterators = new(iterators)

func init() {
	instance := new(iterators)
	Iterators = core.InitializeIfNeeded(instance).(*iterators)
}

type iterators struct {
}

// ---------------------------------------------------------------------
// -- Common iterators
// New sliceIterator instance.
func (t *iterators) NewSliceIterator(elements []interface{}) Iterator {
	iter := new(sliceIterator)
	iter.elements = elements
	core.InitializeIfNeeded(iter)
	return iter
}

// New chainedIterator instance.
func (t *iterators) NewChainedIterator(iterators []Iterator) Iterator {
	iter := new(chainedIterator)
	iter.iterators = iterators
	core.InitializeIfNeeded(iter)
	return iter
}

// ---------------------------------------------------------------------
// -- Common time range iterators
// New Iterator to iter day asc in range [startInclusive, endInclusive]
// Time param would be truncated to day field.
// @param startInclusive start time included, will be truncated to day
// @param endInclusive end time included, will be truncated to day
// @return the Iterator instance
func (t *iterators) NewClosedDayRangeIterator(startInclusive time.Time, endInclusive time.Time) Iterator {
	return t.NewTimeRangeIterator(
		commutil.TimeUtils.TruncateToDay(startInclusive),
		commutil.TimeUtils.TruncateToDay(commutil.TimeUtils.AddDays(endInclusive, 1)),
		func(tm time.Time) time.Time { return commutil.TimeUtils.AddDays(tm, 1) },
	)
}

// New Iterator to iter day asc in range [startInclusive, endExclusive)
// Time param would be truncated to day field.
// @param startInclusive start time included, will be truncated to day
// @param endExclusive end time excluded, will be truncated to day
// @return the Iterator instance
func (t *iterators) NewClosedOpenDayRangeIterator(startInclusive time.Time, endExclusive time.Time) Iterator {
	return t.NewTimeRangeIterator(
		commutil.TimeUtils.TruncateToDay(startInclusive),
		commutil.TimeUtils.TruncateToDay(endExclusive),
		func(tm time.Time) time.Time { return commutil.TimeUtils.AddDays(tm, 1) },
	)
}

// New timeRangeIterator instance.
// @param startInclusive the first element to iter
// @param endExclusive the max limit, not include
// @param nextFunc function to generate next element from previous one
// @return the Iterator instance
func (t *iterators) NewTimeRangeIterator(startInclusive time.Time, endExclusive time.Time, nextFunc func(tm time.Time) time.Time) Iterator {
	iter := new(timeRangeIterator)
	iter.startInclusive = startInclusive
	iter.endExclusive = endExclusive
	iter.nextFunc = nextFunc
	core.InitializeIfNeeded(iter)
	return iter
}
