package commutil

import (
	"testing"
)

func TestIsNumeric(t *testing.T) {
	testCases := []struct {
		str    string
		expect bool
	}{
		{str: "12345", expect: true},
		{str: "12abc", expect: false},
		{str: "0xff", expect: false},
	}

	for i, tc := range testCases {
		result := StringUtils.IsNumeric(tc.str)
		if result != tc.expect {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, result)
		}
	}
}

func TestSubstring(t *testing.T) {
	testCases := []struct {
		str    string
		start  int
		end    int
		expect string
	}{
		{str: "0123456789", start: 0, end: 5, expect: "01234"},
		{str: "0123456789", start: 1, end: 9, expect: "12345678"},
		{str: "0123456789", start: 1, end: 10, expect: "123456789"},
		{str: "0123456789", start: 1, end: 11, expect: "123456789"},
		{str: "0123456789", start: 5, end: 5, expect: ""},
		{str: "0123456789", start: -2, end: 10, expect: "89"},
		{str: "0123456789", start: -1, end: 5, expect: ""},
		{str: "0123456789", start: -100, end: -90, expect: ""},
	}

	for i, tc := range testCases {
		result := StringUtils.Substring(tc.str, tc.start, tc.end)
		if result != tc.expect {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, result)
		}
	}
}
