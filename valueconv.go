package nadesiko4

// ToBool convert to bool
func ToBool(v Value) bool {
	switch v.Type() {
	case TypeBool:
		b := v.(*TBool).value
		return b
	case TypeNumber:
		num := v.Number()
		if num == 0 {
			return false
		}
		return true
	case TypeString:
		str := v.String()
		if str == "" {
			return false
		}
	}
	return false
}

// ToBoolInt returns bool number true(1) or false(0)
func ToBoolInt(v Value) int {
	b := ToBool(v)
	if b {
		return 1
	}
	return 0
}

// ToFloat64 convert to float64
func ToFloat64(v Value) float64 {
	return v.Number()
}

// ToInt convert to int
func ToInt(v Value) int {
	return int(v.Number())
}

// ToString convert to string
func ToString(v Value) string {
	return v.String()
}

// ToValue From native type to Value
func ToValue(v interface{}) Value {
	switch v2 := v.(type) {
	case int:
		return &TNumber{value: float64(v2)}
	case int8:
		return &TNumber{value: float64(v2)}
	case int16:
		return &TNumber{value: float64(v2)}
	case int32:
		return &TNumber{value: float64(v2)}
	case int64:
		return &TNumber{value: float64(v2)}
	case uint:
		return &TNumber{value: float64(v2)}
	case uint32:
		return &TNumber{value: float64(v2)}
	case uint64:
		return &TNumber{value: float64(v2)}
	case float32:
		return &TNumber{value: float64(v2)}
	case float64:
		return &TNumber{value: v2}
	case string:
		return &TString{value: v2}
		// TODO: hash array etc
	}
	return NilValue
}
