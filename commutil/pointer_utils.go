package commutil

import (
	"github.com/xkniu/go-commons-lang/core"
	"github.com/shopspring/decimal"
	"time"
)

var PointerUtils *pointerUtils

func init() {
	instance := new(pointerUtils)
	PointerUtils = core.InitializeIfNeeded(instance).(*pointerUtils)
}

type pointerUtils struct {
}

// ---------------------------------------------------------------------
// -- Pointer from value
// Return a pointer to the value copy.
func (t *pointerUtils) BoolPtr(val bool) *bool {
	return &val
}

// Return a pointer to the value copy.
func (t *pointerUtils) IntPtr(val int) *int {
	return &val
}

// Return a pointer to the value copy.
func (t *pointerUtils) Int32Ptr(val int32) *int32 {
	return &val
}

// Return a pointer to the value copy.
func (t *pointerUtils) Int64Ptr(val int64) *int64 {
	return &val
}

// Return a pointer to the value copy.
func (t *pointerUtils) StringPtr(val string) *string {
	return &val
}

// Return a pointer to the value copy.
func (t *pointerUtils) TimePtr(val time.Time) *time.Time {
	return &val
}

// Return a pointer to the value copy.
func (t *pointerUtils) DecimalPtr(val decimal.Decimal) *decimal.Decimal {
	return &val
}

// ---------------------------------------------------------------------
// -- Safe getter
// Return the value of the pointer, or default value if pointer is nil.
func (t *pointerUtils) Bool(ptr *bool) bool {
	return t.DefaultBool(ptr, false)
}

// Return the value of the pointer, or default value if pointer is nil.
func (t *pointerUtils) Int(ptr *int) int {
	return t.DefaultInt(ptr, 0)
}

// Return the value of the pointer, or default value if pointer is nil.
func (t *pointerUtils) Int32(ptr *int32) int32 {
	return t.DefaultInt32(ptr, 0)
}

// Return the value of the pointer, or default value if pointer is nil.
func (t *pointerUtils) Int64(ptr *int64) int64 {
	return t.DefaultInt64(ptr, 0)
}

// Return the value of the pointer, or default value if pointer is nil.
func (t *pointerUtils) String(ptr *string) string {
	return t.DefaultString(ptr, "")
}

// Return the value of the pointer, or default value if pointer is nil.
func (t *pointerUtils) Time(ptr *time.Time) time.Time {
	return t.DefaultTime(ptr, time.Time{})
}

// Return the value of the pointer, or default value if pointer is nil.
func (t *pointerUtils) Decimal(ptr *decimal.Decimal) decimal.Decimal {
	return t.DefaultDecimal(ptr, decimal.Zero)
}

// ---------------------------------------------------------------------
// -- Default pointer getter
// Return the value of the pointer, or specific default value if pointer is nil.
func (t *pointerUtils) DefaultBool(ptr *bool, defaultVal bool) bool {
	if ptr != nil {
		return *ptr
	}
	return defaultVal
}

// Return the value of the pointer, or specific default value if pointer is nil.
func (t *pointerUtils) DefaultInt(ptr *int, defaultVal int) int {
	if ptr != nil {
		return *ptr
	}
	return defaultVal
}

// Return the value of the pointer, or specific default value if pointer is nil.
func (t *pointerUtils) DefaultInt32(ptr *int32, defaultVal int32) int32 {
	if ptr != nil {
		return *ptr
	}
	return defaultVal
}

// Return the value of the pointer, or specific default value if pointer is nil.
func (t *pointerUtils) DefaultInt64(ptr *int64, defaultVal int64) int64 {
	if ptr != nil {
		return *ptr
	}
	return defaultVal
}

// Return the value of the pointer, or specific default value if pointer is nil.
func (t *pointerUtils) DefaultString(ptr *string, defaultVal string) string {
	if ptr != nil {
		return *ptr
	}
	return defaultVal
}

// Return the value of the pointer, or specific default value if pointer is nil.
func (t *pointerUtils) DefaultTime(ptr *time.Time, defaultVal time.Time) time.Time {
	if ptr != nil {
		return *ptr
	}
	return defaultVal
}

// Return the value of the pointer, or specific default value if pointer is nil.
func (t *pointerUtils) DefaultDecimal(ptr *decimal.Decimal, defaultVal decimal.Decimal) decimal.Decimal {
	if ptr != nil {
		return *ptr
	}
	return defaultVal
}
