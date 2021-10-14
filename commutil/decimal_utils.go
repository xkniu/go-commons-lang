package commutil

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
)

// New instance from string.
// @throws error if the string cannot be parsed as a decimal.
func OfString(s string) decimal.Decimal {
	d, err := decimal.NewFromString(s)
	if err != nil {
		panic(err)
	}
	return d
}

// Convert decimal to int64, panic if the decimal is float.
// @throws error the decimal cannot be represented a int64.
// @throws error if the decimal has fraction part.
func ToInt64(d decimal.Decimal) int64 {
	checkInteger(d)
	return toInt64IgnoreFraction(d)
}

// Convert decimal to int64, ignore fraction part.
// @throws error the decimal cannot be represented a int64.
func ToInt64IgnoreFraction(d decimal.Decimal) int64 {
	return toInt64IgnoreFraction(d)
}

func toInt64IgnoreFraction(d decimal.Decimal) int64 {
	i := d.BigInt()
	if !i.IsInt64() {
		panic(fmt.Sprintf("decimal %v cannot convert to int64", d))
	}
	return i.Int64()
}

// Convert decimal to uint64, ignore fraction part.
// @throws error if the decimal is negative or the decimal cannot be represented a uint64.
// @throws error if the decimal has fraction part.
func ToUint64(d decimal.Decimal) uint64 {
	checkInteger(d)
	return toUint64IgnoreFraction(d)
}

// Convert decimal to uint64, ignore fraction part.
// @throws error if the decimal is negative or the decimal cannot be represented a uint64.
func ToUint64IgnoreFraction(d decimal.Decimal) uint64 {
	return toUint64IgnoreFraction(d)
}

func toUint64IgnoreFraction(d decimal.Decimal) uint64 {
	if d.IsNegative() {
		panic(errors.New(fmt.Sprintf("decimal %v is negative, cannot convert to uint64", d)))
	}
	i := d.BigInt()
	if !i.IsUint64() {
		panic(fmt.Sprintf("decimal %v cannot convert to uint64", d))
	}
	return i.Uint64()
}

// Check if the decimal is an integer.
// @throws error if the decimal has fraction part.
func checkInteger(d decimal.Decimal) {
	if !d.Equal(d.Truncate(0)) {
		panic(fmt.Sprintf("decimal %v is float, cannot convert to int64", d))
	}
}
