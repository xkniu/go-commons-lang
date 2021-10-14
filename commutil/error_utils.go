package commutil

import (
	"errors"
	"fmt"
	"github.com/xkniu/go-commons-lang/core"
	"runtime"
)

var ErrorUtils *errorUtils

func init() {
	instance := new(errorUtils)
	ErrorUtils = core.InitializeIfNeeded(instance).(*errorUtils)
}

type errorUtils struct {
}

// Panic if error is not nil.
func (t *errorUtils) PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

// Error with file and line info.
// @param err origin error
// @param skip stack to skip. If you build your own utils, you usually need to set this to one or more.
// @return new error with file and line info.
func (t *errorUtils) ErrorWithFileLine(err error, skip ...int) error {
	if err == nil {
		return nil
	}
	toSkip := 1
	if len(skip) > 0 {
		toSkip += skip[0]
	}
	return t.newError(err, toSkip, func(err error, file string, line int) error {
		return errors.New(fmt.Sprintf("%s:%d %v", file, line, err))
	})
}

// New error with runtime info.
// @param err origin error
// @param skip stack to skip. If you build your own utils, you usually need to set this to one or more.
// @param wrapper function to wrap error and file, line info
// @return error wrapped
func (t *errorUtils) newError(err error, skip int, wrapper func(err error, file string, line int) error) error {
	_, file, line, _ := runtime.Caller(skip + 1)
	return wrapper(err, file, line)
}
