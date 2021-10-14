package commutil

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestOfString(t *testing.T) {
	testCases := []struct {
		value       string
		expect      decimal.Decimal
		expectError bool
	}{
		{value: "123", expect: decimal.NewFromInt(123), expectError: false},
		{value: "by", expect: decimal.Zero, expectError: true},
	}

	for i, tc := range testCases {
		if tc.expectError {
			AssertUtils.AssertPanic(func() {
				OfString(tc.value)
			})
		} else {
			result := OfString(tc.value)
			if !result.Equal(tc.expect) {
				t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, result)
			}
		}
	}
}

func TestToInt64(t *testing.T) {
	testCases := []struct {
		value       decimal.Decimal
		expect      int64
		expectError bool
	}{
		{value: OfString("123.45"), expect: 123, expectError: true},
		{value: OfString("100"), expect: 100, expectError: false},
		{value: OfString("100.0"), expect: 100, expectError: false},
		{value: OfString("-100"), expect: -100, expectError: false},
		{value: decimal.New(1, 100), expect: 0, expectError: true},
	}
	for i, tc := range testCases {
		if tc.expectError {
			AssertUtils.AssertPanic(func() {
				ToInt64(tc.value)
			})
		} else {
			result := ToInt64(tc.value)
			if result != tc.expect {
				t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, result)
			}
		}
	}
}

func TestToInt64IgnoreFraction(t *testing.T) {
	testCases := []struct {
		value  decimal.Decimal
		expect int64
	}{
		{value: OfString("123.88"), expect: 123},
		{value: OfString("123.45"), expect: 123},
		{value: OfString("100"), expect: 100},
		{value: OfString("100.0"), expect: 100},
		{value: OfString("-100"), expect: -100},
	}
	for i, tc := range testCases {
		result := ToInt64IgnoreFraction(tc.value)
		if result != tc.expect {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, result)
		}
	}
}

func TestToUint64(t *testing.T) {
	testCases := []struct {
		value       decimal.Decimal
		expect      uint64
		expectError bool
	}{
		{value: OfString("123.88"), expect: 0, expectError: true},
		{value: OfString("123.45"), expect: 0, expectError: true},
		{value: OfString("100"), expect: 100, expectError: false},
		{value: OfString("100.0"), expect: 100, expectError: false},
		{value: OfString("-100"), expect: 0, expectError: true},
		{value: decimal.New(1, 100), expect: 0, expectError: true},
	}
	for i, tc := range testCases {
		if tc.expectError {
			AssertUtils.AssertPanic(func() {
				ToUint64(tc.value)
			})
		} else {
			result := ToUint64(tc.value)
			if result != tc.expect {
				t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, result)
			}
		}
	}
}

func TestToUint64IgnoreFraction(t *testing.T) {
	testCases := []struct {
		value       decimal.Decimal
		expect      uint64
		expectError bool
	}{
		{value: OfString("123.88"), expect: 123, expectError: false},
		{value: OfString("123.45"), expect: 123, expectError: false},
		{value: OfString("100"), expect: 100, expectError: false},
		{value: OfString("100.0"), expect: 100, expectError: false},
		{value: OfString("-100"), expect: 0, expectError: true},
	}
	for i, tc := range testCases {
		if tc.expectError {
			AssertUtils.AssertPanic(func() {
				ToUint64IgnoreFraction(tc.value)
			})
		} else {
			result := ToUint64IgnoreFraction(tc.value)
			if result != tc.expect {
				t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, result)
			}
		}
	}
}
