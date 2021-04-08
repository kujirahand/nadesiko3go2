package nadesiko4

import (
	"fmt"
	"strconv"
)

const (
	ErrNotImplemented = "(未実装です)"
	TypeNameString    = "文字列"
	TypeNameNumber    = "数値"
	TypeNameInt       = "整数"
	TypeNameBool      = "真偽型"
	TypeNameArray     = "配列"
	TypeNameDict      = "辞書型"
	TypeNameNull      = "NULL"
)

// Object 仮想マシンで使われる基本オブジェクトのインターフェイス
type Object interface {
	// TypeName returns the type string
	TypeName() string
	// String returns string value
	String() string
	// Number returns number value
	Number() float64
	// Int returns int value
	Int() int
	// Bool returns bool value
	Bool() bool
	// Clone returns clone object
	Clone() Object
	// GetIndex get value
	GetIndex(index int) Object
	// SetIndex set value
	SetIndex(index int, value Object)
	// GetDict get value
	GetDict(key string) Object
	// SetDict set value
	SetDict(key string, value Object)
	// Length returns size of object
	Length() int
	// StrictEqual retuens true or false
	StrictEqual(target Object) bool
	// Equal retuens true or false
	Equal(target Object) bool
}

var (
	// NullValue is null
	NullValue Object = &Null{}
	// TrueValue is true
	TrueValue Object = &Bool{Value: true}
	// FalseValue is false
	FalseValue Object = &Bool{Value: false}
)

// ------------------------------------------------------------------
// ObjectBase 未実装を検出するためのオブジェクト実装
// ------------------------------------------------------------------
type ObjectBase struct {
	Object
}

// TypeName returns the type string
func (o *ObjectBase) TypeName() string {
	panic(ErrNotImplemented)
}

// String returns string value
func (o *ObjectBase) String() string {
	panic(ErrNotImplemented)
}

// Number returns number value
func (o *ObjectBase) Number() float64 {
	panic(ErrNotImplemented)
}

// Int returns int value
func (o *ObjectBase) Int() int {
	panic(ErrNotImplemented)
}

// Bool returns bool value
func (o *ObjectBase) Bool() bool {
	v := o.Int()
	return (v != 0)
}

// Clone returns clone object
func (o *ObjectBase) Clone() Object {
	panic(ErrNotImplemented)
}

// GetIndex get value
func (o *ObjectBase) GetIndex(index int) Object {
	panic(ErrNotImplemented)
}

// SetIndex set value
func (o *ObjectBase) SetIndex(index int, value Object) {
	panic(ErrNotImplemented)
}

// Length returns size of object
func (o *ObjectBase) Length() int {
	panic(ErrNotImplemented)
}

// GetDict get value
func (o *ObjectBase) GetDict(key string) Object {
	panic(ErrNotImplemented)
}

// SetDict set value
func (o *ObjectBase) SetDict(key string, value Object) {
	panic(ErrNotImplemented)
}

// StrictEqual retuens true or false
func (o *ObjectBase) StrictEqual(target Object) bool {
	return o == target
}

// Equal retuens true or false
func (o *ObjectBase) Equal(target Object) bool {
	return o == target
}

// ------------------------------------------------------------------
// String type 文字列型
// ------------------------------------------------------------------
type String struct {
	ObjectBase
	Value string
}

// TypeName returns the type string
func (o *String) TypeName() string {
	return TypeNameString
}

// String returns string value
func (o *String) String() string {
	return o.Value
}

// Number returns number value
func (o *String) Number() float64 {
	v, err := strconv.ParseFloat(o.Value, 64)
	if err != nil {
		return 0
	}
	return v
}

// Int returns int value
func (o *String) Int() int {
	v, err := strconv.ParseInt(o.Value, 10, 32)
	if err != nil {
		return 0
	}
	return int(v)
}

// Clone returns clone object
func (o *String) Clone() Object {
	return &String{Value: o.Value}
}

// Length returns bytes of String (NOT Character length)
func (o *String) Length() int {
	return len(o.Value)
}

// StrictEqual retuens true or false
func (o *String) StrictEqual(target Object) bool {
	x, ok := target.(*String)
	if !ok {
		return false
	}
	return x.Value == o.Value
}

// Equal retuens true or false
func (o *String) Equal(target Object) bool {
	t := target.String()
	return t == o.Value
}

// ------------------------------------------------------------------
// Number type
// ------------------------------------------------------------------
type Number struct {
	ObjectBase
	Value float64
}

func NewNumber(def float64) Object {
	return &Number{Value: def}
}

// TypeName returns the type string
func (o *Number) TypeName() string {
	return TypeNameNumber
}

// String returns string value
func (o *Number) String() string {
	return fmt.Sprintf("%f", o.Value)
}

// Number returns number value
func (o *Number) Number() float64 {
	return o.Value
}

// Int returns int value
func (o *Number) Int() int {
	return int(o.Value)
}

// Clone returns clone object
func (o *Number) Clone() Object {
	return &Number{Value: o.Value}
}

// Length returns size
func (o *Number) Length() int {
	return 8
}

// StrictEqual retuens true or false
func (o *Number) StrictEqual(target Object) bool {
	x, ok := target.(*Number)
	if !ok {
		return false
	}
	return x.Value == o.Value
}

// Equal retuens true or false
func (o *Number) Equal(target Object) bool {
	t := target.Number()
	return o.Value == t
}

// ------------------------------------------------------------------
// Int type
// ------------------------------------------------------------------
type Int struct {
	ObjectBase
	Value int
}

func NewInt(def int) Object {
	return &Int{Value: def}
}

// TypeName returns the type string
func (o *Int) TypeName() string {
	return TypeNameInt
}

// String returns string value
func (o *Int) String() string {
	return strconv.Itoa(o.Value)
}

// Number returns number value
func (o *Int) Number() float64 {
	return float64(o.Value)
}

// Int returns int value
func (o *Int) Int() int {
	return o.Value
}

// Clone returns clone object
func (o *Int) Clone() Object {
	return &Int{Value: o.Value}
}

// Length returns size
func (o *Int) Length() int {
	return 1
}

// StrictEqual retuens true or false
func (o *Int) StrictEqual(target Object) bool {
	x, ok := target.(*Int)
	if !ok {
		return false
	}
	return x.Value == o.Value
}

// Equal retuens true or false
func (o *Int) Equal(target Object) bool {
	t := target.Int()
	return t == o.Value
}

// ------------------------------------------------------------------
// Bool type
// ------------------------------------------------------------------
type Bool struct {
	ObjectBase
	Value bool
}

// TypeName returns the type string
func (o *Bool) TypeName() string {
	return TypeNameNumber
}

// String returns string value
func (o *Bool) String() string {
	if o.Value {
		return "真"
	} else {
		return "偽"
	}
}

// Number returns number value
func (o *Bool) Number() float64 {
	return float64(o.Int())
}

// Int returns int value
func (o *Bool) Int() int {
	if o.Value {
		return 1
	} else {
		return 0
	}
}

// Clone returns clone object
func (o *Bool) Clone() Object {
	return &Bool{Value: o.Value}
}

// Length returns size
func (o *Bool) Length() int {
	return 1
}

// StrictEqual retuens true or false
func (o *Bool) StrictEqual(target Object) bool {
	x, ok := target.(*Bool)
	if !ok {
		return false
	}
	return x.Value == o.Value
}

// Equal retuens true or false
func (o *Bool) Equal(target Object) bool {
	t := target.Bool()
	return t == o.Value
}

// ------------------------------------------------------------------
// Array type
// ------------------------------------------------------------------
type Array struct {
	ObjectBase
	Value []Object
}

func NewArray() Object {
	a := make([]Object, 0)
	o := &Array{Value: a}
	return o
}

// TypeName returns the type string
func (o *Array) TypeName() string {
	return TypeNameArray
}

// String returns string value
func (o *Array) String() string {
	panic(ErrNotImplemented)
}

// Number returns number value
func (o *Array) Number() float64 {
	return float64(o.Int())
}

// Int returns int value
func (o *Array) Int() int {
	return len(o.Value)
}

// Clone returns clone object
func (o *Array) Clone() Object {
	var v []Object
	for _, el := range o.Value {
		v = append(v, el.Clone())
	}
	return &Array{Value: v}
}

// GetIndex get value
func (o *Array) GetIndex(index int) Object {
	if (index < 0) || (index >= len(o.Value)) {
		return NullValue
	}
	return o.Value[index]
}

// SetIndex set value
func (o *Array) SetIndex(index int, value Object) {
	// 自動的に要素を伸ばす
	for index >= len(o.Value) {
		o.Value = append(o.Value, NullValue)
	}
	o.Value[index] = value
}

// Length returns size
func (o *Array) Length() int {
	return len(o.Value)
}

// StrictEqual retuens true or false
func (o *Array) StrictEqual(target Object) bool {
	x, ok := target.(*Array)
	if !ok {
		return false
	}
	if x.Length() != o.Length() {
		return false
	}
	for i := 0; i < o.Length(); i++ {
		ov := o.GetIndex(i)
		tv := x.GetIndex(i)
		if !ov.StrictEqual(tv) {
			return false
		}
	}
	return true
}

// Equal retuens true or false
func (o *Array) Equal(target Object) bool {
	switch target.(type) {
	case *Array:
		return o.StrictEqual(target)
	case *String:
		//todo: convert to array
		return false
	default:
		return false
	}
}

// ------------------------------------------------------------------
// Dict type
// ------------------------------------------------------------------
type Dict struct {
	ObjectBase
	Value map[string]Object
}

func NewDict() Object {
	v := map[string]Object{}
	return &Dict{Value: v}
}

// TypeName returns the type string
func (o *Dict) TypeName() string {
	return TypeNameDict
}

// String returns string value
func (o *Dict) String() string {
	panic(ErrNotImpl)
}

// Number returns number value
func (o *Dict) Number() float64 {
	return float64(o.Int())
}

// Int returns int value
func (o *Dict) Int() int {
	return len(o.Value)
}

// Clone returns clone object
func (o *Dict) Clone() Object {
	p := make(map[string]Object)
	for k, v := range o.Value {
		p[k] = v.Clone()
	}
	return &Dict{Value: o.Value}
}

// Length returns size
func (o *Dict) Length() int {
	return len(o.Value)
}

// GetDict get value
func (o *Dict) GetDict(key string) Object {
	v, ok := o.Value[key]
	if !ok {
		return NullValue
	}
	return v
}

// SetDict set value
func (o *Dict) SetDict(key string, value Object) {
	o.Value[key] = value
}

// StrictEqual retuens true or false
func (o *Dict) StrictEqual(target Object) bool {
	x, ok := target.(*Dict)
	if !ok {
		return false
	}
	if x.Length() != o.Length() {
		return false
	}
	for k, v := range x.Value {
		if !o.Value[k].StrictEqual(v) {
			return false
		}
	}
	return true
}

// Equal retuens true or false
func (o *Dict) Equal(target Object) bool {
	switch target.(type) {
	case *Dict:
		return o.StrictEqual(target)
	case *String:
		// todo: convert to Dict
		return false
	default:
		return false
	}
}

// ------------------------------------------------------------------
// Null type
// ------------------------------------------------------------------
type Null struct {
	ObjectBase
}

// TypeName returns the type string
func (o *Null) TypeName() string {
	return TypeNameNull
}

// String returns string value
func (o *Null) String() string {
	return ""
}

// Number returns number value
func (o *Null) Number() float64 {
	return 0
}

// Int returns int value
func (o *Null) Int() int {
	return 0
}

// Clone returns clone object
func (o *Null) Clone() Object {
	return &Null{}
}

// Length returns size
func (o *Null) Length() int {
	return 0
}

// StrictEqual retuens true or false
func (o *Null) StrictEqual(target Object) bool {
	_, ok := target.(*Null)
	return ok
}

// Equal retuens true or false
func (o *Null) Equal(target Object) bool {
	return o.StrictEqual(target)
}
