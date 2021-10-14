package commutil

import (
	"github.com/xkniu/go-commons-lang/core"
	"strconv"
)

var NumberUtils *numberUtils

func init() {
	instance := new(numberUtils)
	NumberUtils = core.InitializeIfNeeded(instance).(*numberUtils)
}

type numberUtils struct {
}

// Parse the string to int, panic with error.
// @throws error if the string cannot be parsed as a number.
func (t *numberUtils) ParseInt(s string) int {
	parsed, err := strconv.ParseInt(s, 0, 0)
	if err != nil {
		panic(err)
	}
	return int(parsed)
}

// Parse the string to int32, panic with error.
// @throws panic if the string cannot be parsed as a number.
func (t *numberUtils) ParseInt32(s string) int32 {
	parsed, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		panic(err)
	}
	return int32(parsed)
}

// Parse the string to int64, panic with error.
// @throws panic if the string cannot be parsed as a number.
func (t *numberUtils) ParseInt64(s string) int64 {
	parsed, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		panic(err)
	}
	return parsed
}
