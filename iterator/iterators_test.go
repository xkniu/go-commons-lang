package iterator

import (
	"github.com/xkniu/go-commons-lang/commutil"
	"reflect"
	"testing"
	"time"
)

func TestSliceIterator(t *testing.T) {
	testCases := []struct {
		elements []interface{}
	}{
		{elements: []interface{}{1, 2, 3, 4, 5}},
		{elements: []interface{}{5, 4, 3, 2, 1}},
		{elements: []interface{}{"a", "b", "c"}},
	}

	for i, tc := range testCases {
		iter := Iterators.NewSliceIterator(tc.elements)
		index := 0
		for iter.HasNext() {
			elem := iter.Next()
			if elem != tc.elements[index] {
				t.Errorf("case %d-%d: expected: %v, got: %v", i, index, tc.elements[index], elem)
			}
			index++
		}
		if index != len(tc.elements) {
			t.Errorf("case %d: expected: %v, got: %v", i, len(tc.elements), index)
		}
	}
}

func TestChainedIterator(t *testing.T) {
	testCases := []struct {
		iterators   []Iterator
		expectSlice []interface{}
	}{
		{
			iterators:   []Iterator{Iterators.NewSliceIterator([]interface{}{1, 2, 3}), Iterators.NewSliceIterator([]interface{}{}), Iterators.NewSliceIterator([]interface{}{4, 5})},
			expectSlice: []interface{}{1, 2, 3, 4, 5},
		},
	}

	for i, tc := range testCases {
		iter := Iterators.NewChainedIterator(tc.iterators)
		result := ToSlice(iter)
		if !reflect.DeepEqual(result, tc.expectSlice) {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expectSlice, result)
		}
	}
}

func TestClosedDayRangeIterator(t *testing.T) {
	testCases := []struct {
		startInclusive time.Time
		endInclusive   time.Time
		expectSize     int
	}{
		{startInclusive: commutil.TimeUtils.OfYmd(2020, 1, 1), endInclusive: commutil.TimeUtils.OfYmd(2019, 12, 31), expectSize: 0},
		{startInclusive: commutil.TimeUtils.OfYmd(2020, 1, 1), endInclusive: commutil.TimeUtils.OfYmd(2020, 1, 1), expectSize: 1},
		{startInclusive: commutil.TimeUtils.OfYmd(2020, 1, 1), endInclusive: commutil.TimeUtils.OfYmd(2020, 1, 2), expectSize: 2},
		{startInclusive: commutil.TimeUtils.OfYmd(2020, 1, 1), endInclusive: commutil.TimeUtils.OfYmd(2020, 2, 1), expectSize: 32},
	}

	for i, tc := range testCases {
		iter := Iterators.NewClosedDayRangeIterator(tc.startInclusive, tc.endInclusive)
		result := Size(iter)
		if result != tc.expectSize {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expectSize, result)
		}
	}
}

func TestClosedOpenDayRangeIterator(t *testing.T) {
	testCases := []struct {
		startInclusive time.Time
		endInclusive   time.Time
		expectSize     int
	}{
		{startInclusive: commutil.TimeUtils.OfYmd(2020, 1, 1), endInclusive: commutil.TimeUtils.OfYmd(2019, 12, 31), expectSize: 0},
		{startInclusive: commutil.TimeUtils.OfYmd(2020, 1, 1), endInclusive: commutil.TimeUtils.OfYmd(2020, 1, 1), expectSize: 0},
		{startInclusive: commutil.TimeUtils.OfYmd(2020, 1, 1), endInclusive: commutil.TimeUtils.OfYmd(2020, 1, 2), expectSize: 1},
		{startInclusive: commutil.TimeUtils.OfYmd(2020, 1, 1), endInclusive: commutil.TimeUtils.OfYmd(2020, 2, 1), expectSize: 31},
	}

	for i, tc := range testCases {
		iter := Iterators.NewClosedOpenDayRangeIterator(tc.startInclusive, tc.endInclusive)
		result := Size(iter)
		if result != tc.expectSize {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expectSize, result)
		}
	}
}
