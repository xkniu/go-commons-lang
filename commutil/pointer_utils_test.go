package commutil

import "testing"

func TestInt64Ptr(t *testing.T) {
	testCases := []struct {
		val int64
	}{
		{val: 1},
		{val: -1},
		{val: 0},
	}

	for i, tc := range testCases {
		result := PointerUtils.Int64Ptr(tc.val)
		if *result != tc.val {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.val, *result)
		}
	}
}

func TestInt64(t *testing.T) {
	testCases := []struct {
		ptr    *int64
		expect int64
	}{
		{ptr: PointerUtils.Int64Ptr(1), expect: 1},
		{ptr: nil, expect: 0},
	}

	for i, tc := range testCases {
		result := PointerUtils.Int64(tc.ptr)
		if result != tc.expect {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, result)
		}
	}
}

func TestDefaultInt64(t *testing.T) {
	testCases := []struct {
		ptr        *int64
		defaultVal int64
		expect     int64
	}{
		{ptr: PointerUtils.Int64Ptr(1), defaultVal: -1, expect: 1},
		{ptr: nil, defaultVal: -1, expect: -1},
	}

	for i, tc := range testCases {
		result := PointerUtils.DefaultInt64(tc.ptr, tc.defaultVal)
		if result != tc.expect {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expect, result)
		}
	}
}
