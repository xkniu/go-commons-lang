package commutil

import (
	"errors"
	"strings"
	"testing"
)

func TestErrorWithFileLine(t *testing.T) {
	testCases := []struct {
		skip           int
		expectFileName string
	}{
		{skip: 0, expectFileName: "error_utils_test.go"},
		{skip: 1, expectFileName: "testing.go"},
	}

	for i, tc := range testCases {
		err := ErrorUtils.ErrorWithFileLine(errors.New("hello, world"), tc.skip)
		if err == nil {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expectFileName, nil)
		} else if !strings.Contains(err.Error(), tc.expectFileName) {
			t.Errorf("case %d: expected: %v, got: %v", i, tc.expectFileName, err.Error())
		}
	}
}
