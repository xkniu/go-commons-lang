package commutil

import (
	"errors"
	"github.com/xkniu/go-commons-lang/core"
)

var AssertUtils *assertUtils

func init() {
	instance := new(assertUtils)
	AssertUtils = core.InitializeIfNeeded(instance).(*assertUtils)
}

type assertUtils struct {
}

func (t *assertUtils) AssertPanic(call func()) {
	err := CallUtils.Quietly(call)
	if err == nil {
		panic(errors.New("expect call with panic but it does not"))
	}
}
