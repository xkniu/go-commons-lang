package commutil

import (
	"github.com/xkniu/go-commons-lang/core"
	"strings"
	"unicode"
)

var StringUtils *stringUtils

func init() {
	instance := new(stringUtils)
	StringUtils = core.InitializeIfNeeded(instance).(*stringUtils)
}

type stringUtils struct {
	EMPTY string // The empty String ""
	SPACE string // The string for a space character
	LF    string // The string "\n"
	CR    string // The string "\r"
}

func (t *stringUtils) Init() {
	t.EMPTY = ""
	t.SPACE = " "
	t.LF = "\n"
	t.CR = "\r"
}

// Return true if the two strings are the same, ignoring case.
func (t *stringUtils) EqualIgnoreCase(str1, str2 string) bool {
	return strings.EqualFold(str1, str2)
}

// Return the default string if value is empty.
func (t *stringUtils) DefaultIfEmpty(value string, defaultString string) string {
	if len(value) == 0 {
		return defaultString
	}
	return value
}

// Return a pointer of the first non empty value. If all values are empty, return nil.
func (t *stringUtils) FirstNonEmpty(values ...string) *string {
	for _, value := range values {
		if len(values) > 0 {
			return &value
		}
	}
	return nil
}

// Checks if the string contains only Unicode digits.
func (t *stringUtils) IsNumeric(str string) bool {
	for _, c := range str {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// Checks if the string contains only Unicode digits or space.
func (t *stringUtils) IsNumericSpace(str string) bool {
	for _, c := range str {
		if !unicode.IsDigit(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// Return a new string that is a substring of this string without exception.
// @param str origin string to substitute
// @param start include, negative number means start from end of the str
// @pram end exclude, negative number means start from end of the str
func (t *stringUtils) Substring(str string, start int, end int) string {
	if str == "" {
		return ""
	}

	// negative iter recursively
	if start < 0 {
		start += len(str)
	}
	if end < 0 {
		end += len(str)
	}
	if end > len(str) {
		end = len(str)
	}
	if start >= end {
		return ""
	}
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	return str[start:end]
}
