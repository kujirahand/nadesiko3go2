package nadesiko4

import (
	"errors"
	"fmt"
	"strconv"
)

// ValueType is Value type
type ValueType int

// ValueType list
const (
	TypeNil ValueType = iota
	TypeBool
	TypeNumber
	TypeString
	TypeArray
	TypeHash
	TypeFunction
)

var typeNames = map[ValueType]string{
	TypeNil:      "nil",
	TypeBool:     "bool",
	TypeNumber:   "number",
	TypeString:   "string",
	TypeArray:    "array",
	TypeHash:     "hash",
	TypeFunction: "function",
}

// GetValueTypeName returns value type name.
func GetValueTypeName(t ValueType) string {
	v, ok := typeNames[t]
	if ok {
		return v
	}
	return "INVALID_TYPE"
}

// Default Values
var (
	// NilValue is nil value
	NilValue Value = &TNil{}
	// TrueValue is true
	TrueValue Value = &TBool{value: true}
	// FalseValue is false
	FalseValue Value = &TBool{value: false}
)

// Value interfalce
type Value interface {
	//  should return the name of type.
	Type() ValueType

	// String should return value string.
	String() string

	// Number should return number.
	Number() float64

	// Compare should return bool
	Compare(op VMCodeType, val Value) bool
}

// ErrNotImpl is error
var ErrNotImpl = errors.New("未実装")

// ValueImpl implements Value interface
//
type ValueImpl struct {
	Value
}

// Type returns error
func (v *ValueImpl) Type() ValueType {
	panic(ErrNotImpl)
}

// String returns string
func (v *ValueImpl) String() string {
	return ""
}

// Number returns number
func (v *ValueImpl) Number() float64 {
	return 0
}

// Compare returns bool
func (v *ValueImpl) Compare(op VMCodeType, val Value) bool {
	return false
}

// TNil is nil object
type TNil struct {
	ValueImpl
}

// Type returns VTNil
func (v *TNil) Type() ValueType {
	return TypeNil
}

// TBool type
type TBool struct {
	ValueImpl
	value bool
}

// Type returns VTBool
func (v *TBool) Type() ValueType {
	return TypeBool
}

func (v *TBool) String() string {
	if v.value {
		return "真"
	}
	return "偽"
}

// Number returns 0 or 1
func (v *TBool) Number() float64 {
	if v.value {
		return 0
	}
	return 1
}

// Compare returns bool
func (v *TBool) Compare(op VMCodeType, val Value) bool {
	switch op {
	case TypeOpEq:
		return v.value == ToBool(val)
	case TypeOpNtEq:
		return v.value != ToBool(val)
	// Goではboolの比較がエラーになるので整数として処理する
	case TypeOpGt:
		return ToBoolInt(v) > ToBoolInt(val)
	case TypeOpGtEq:
		return ToBoolInt(v) >= ToBoolInt(val)
	case TypeOpLt:
		return ToBoolInt(v) < ToBoolInt(val)
	case TypeOpLtEq:
		return ToBoolInt(v) <= ToBoolInt(val)
	}
	return false
}

// TNumber type
type TNumber struct {
	ValueImpl
	value float64
}

// Type returns VTNumber
func (v *TNumber) Type() ValueType {
	return TypeNumber
}

// Number returns value
func (v *TNumber) Number() float64 {
	return v.value
}

func (v *TNumber) String() string {
	return fmt.Sprint(v.value)
}

// Compare returns bool
func (v *TNumber) Compare(op VMCodeType, val Value) bool {
	switch op {
	case TypeOpEq:
		return v.value == val.Number()
	case TypeOpNtEq:
		return v.value != val.Number()
	case TypeOpGt:
		return v.value > val.Number()
	case TypeOpGtEq:
		return v.value >= val.Number()
	case TypeOpLt:
		return v.value < val.Number()
	case TypeOpLtEq:
		return v.value <= val.Number()
	}
	return false
}

// TString type
type TString struct {
	ValueImpl
	value string
}

// Type returns TypeString
func (v *TString) Type() ValueType {
	return TypeString
}

// Number returns number value
func (v *TString) Number() float64 {
	num, err := strconv.ParseFloat(v.value, 64)
	if err != nil {
		return 0
	}
	return num
}

func (v *TString) String() string {
	return fmt.Sprint(v.value)
}

// Compare returns bool
func (v *TString) Compare(op VMCodeType, val Value) bool {
	switch op {
	case TypeOpEq:
		return v.value == val.String()
	case TypeOpNtEq:
		return v.value != val.String()
	case TypeOpGt:
		return v.value > val.String()
	case TypeOpGtEq:
		return v.value >= val.String()
	case TypeOpLt:
		return v.value < val.String()
	case TypeOpLtEq:
		return v.value <= val.String()
	}
	return false
}
