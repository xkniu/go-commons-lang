package commutil

import (
	"testing"
	"time"
)

func TestIsSameDay(t *testing.T) {
	testCases := []struct {
		tm1  time.Time
		tm2    time.Time
		expect bool
	}{
		{tm1: TimeUtils.OfYmd(2020, 1, 1), tm2: TimeUtils.OfYmd(2020, 1, 1), expect: true},
		{tm1: TimeUtils.OfYmdh(2020, 1, 1, 1), tm2: TimeUtils.OfYmdh(2020, 1, 1, 2), expect: true},
		{tm1: TimeUtils.OfYmd(2020, 1, 1), tm2: TimeUtils.OfYmd(2021, 1, 1), expect: false},
		{tm1: TimeUtils.OfYmd(2020, 1, 1), tm2: TimeUtils.OfYmd(2020, 2, 1), expect: false},
		{tm1: TimeUtils.OfYmd(2020, 1, 1), tm2: TimeUtils.OfYmd(2021, 2, 2), expect: false},
	}

	for i, tc := range testCases {
		result := TimeUtils.IsSameDay(tc.tm1, tc.tm2)
		if result != tc.expect {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, result)
		}
	}
}

func TestMonthsBetweenInclusive(t *testing.T) {
	testCases := []struct {
		start  time.Time
		end    time.Time
		expect int
	}{
		{start: TimeUtils.OfYmd(2020, 1, 1), end: TimeUtils.OfYmd(2020, 2, 1), expect: 1},
		{start: TimeUtils.OfYmdh(2020, 1, 1, 9), end: TimeUtils.OfYmdh(2020, 2, 1, 0), expect: 0},
		{start: TimeUtils.OfYmdh(2020, 1, 1, 9), end: TimeUtils.OfYmdh(2020, 2, 1, 9), expect: 1},
		{start: TimeUtils.OfYmd(2020, 1, 1), end: TimeUtils.OfYmd(2021, 2, 1), expect: 13},
		{start: TimeUtils.OfYmd(2021, 2, 1), end: TimeUtils.OfYmd(2020, 1, 1), expect: -13},
	}

	for i, tc := range testCases {
		result := TimeUtils.MonthsBetweenInclusive(tc.start, tc.end)
		if result != tc.expect {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, result)
		}
	}
}

func TestDaysBetweenInclusive(t *testing.T) {
	testCases := []struct {
		start  time.Time
		end    time.Time
		expect int
	}{
		{start: TimeUtils.OfYmd(2020, 1, 1), end: TimeUtils.OfYmd(2020, 1, 2), expect: 1},
		{start: TimeUtils.OfYmd(2020, 1, 1), end: TimeUtils.OfYmd(2021, 1, 1), expect: 366},
		{start: TimeUtils.OfYmdh(2020, 1, 1, 9), end: TimeUtils.OfYmdh(2020, 1, 2, 0), expect: 0},
		{start: TimeUtils.OfYmd(2020, 1, 2), end: TimeUtils.OfYmd(2020, 1, 1), expect: -1},
	}

	for i, tc := range testCases {
		result := TimeUtils.DaysBetweenInclusive(tc.start, tc.end)
		if result != tc.expect {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, result)
		}
	}
}
