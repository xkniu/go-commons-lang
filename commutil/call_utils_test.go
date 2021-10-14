package commutil

import (
	"errors"
	"testing"
)

func TestQuietly(t *testing.T) {
	testCases := []struct {
		call   func()
		expect error
	}{
		{call: func() {}, expect: nil},
		{call: func() { panic(errors.New("hello, world")) }, expect: errors.New("hello, world")},
	}

	for i, tc := range testCases {
		err := CallUtils.Quietly(tc.call)
		if err == nil {
			if tc.expect != nil {
				t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, err)
			}
		} else if err.Error() != tc.expect.Error() {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, err)
		}
	}
}

func TestQuietly2(t *testing.T) {
	testCases := []struct {
		call   func() error
		expect error
	}{
		{call: func() error { return nil }, expect: nil},
		{call: func() error { panic(errors.New("hello, world")) }, expect: errors.New("hello, world")},
		{call: func() error { return errors.New("hello, world") }, expect: errors.New("hello, world")},
	}

	for i, tc := range testCases {
		err := CallUtils.Quietly2(tc.call)
		if err == nil {
			if tc.expect != nil {
				t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, err)
			}
		} else if err.Error() != tc.expect.Error() {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, err)
		}
	}
}
