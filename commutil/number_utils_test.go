package commutil

import (
	"testing"
)

func TestParseInt64(t *testing.T) {
	testCases := []struct {
		s      string
		expect int64
	}{
		{s: "123", expect: 123},
		{s: "-123", expect: -123},
		{s: "0xff", expect: 255},
		{s: "0o10", expect: 8},
	}

	for i, tc := range testCases {
		result := NumberUtils.ParseInt64(tc.s)
		if result != tc.expect {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, result)
		}
	}
}

func TestParseInt64WithError(t *testing.T) {
	testCases := []struct {
		s      string
	}{
		{s: "abc"},
		{s: "1231a"},
		{s: "+"},
	}

	for i, tc := range testCases {
		err := CallUtils.Quietly(func() {
			NumberUtils.ParseInt64(tc.s)
		})
		if err == nil {
			t.Errorf("case %d: expected: %v, got: %v", i, "error", nil)
		}
	}
}
