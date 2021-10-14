package commutil

import (
	"errors"
	"fmt"
	"github.com/xkniu/go-commons-lang/core"
)

var CallUtils *callUtils

func init() {
	instance := new(callUtils)
	CallUtils = core.InitializeIfNeeded(instance).(*callUtils)
}

type callUtils struct {
}

// Alias for method Quietly.
func (t *callUtils) Safely(call func()) (err error) {
	return t.Quietly(call)
}
// Call a function quietly, safely.
// Wrap as error if function has a panic.
func (t *callUtils) Quietly(call func()) (err error) {
	defer func() {
		if e := recover(); e != nil {
			switch e.(type) {
			case error:
				err = e.(error)
			default:
				err = errors.New(fmt.Sprintf("call with error: %v", e))
			}
		}
	}()
	call()
	return err
}

// Alias for method Quietly2.
func (t *callUtils) Safely2(call func() error) (err error) {
	return t.Quietly2(call)
}

// Call a function quietly, safely.
// Wrap as error if function has a panic, or return the origin error if function has non panic.
func (t *callUtils) Quietly2(call func() error) (err error) {
	defer func() {
		if e := recover(); e != nil {
			switch e.(type) {
			case error:
				err = e.(error)
			default:
				err = errors.New(fmt.Sprintf("call with error: %v", e))
			}
		}
	}()
	err = call()
	return err
}