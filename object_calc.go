package nadesiko4

// ObjectAdd return a + b
func ObjectAdd(a Object, b Object) Object {
	c := NewNumber(a.Number() + b.Number())
	return c
}

// ObjectAddStr return a & b
func ObjectAddStr(a Object, b Object) Object {
	s := a.String() + b.String()
	return &String{Value: s}
}

// ObjectSub return a - b
func ObjectSub(a Object, b Object) Object {
	c := NewNumber(a.Number() - b.Number())
	return c
}

// ObjectMul return a * b
func ObjectMul(a Object, b Object) Object {
	c := NewNumber(a.Number() * b.Number())
	return c
}

// ObjectDiv return a / b
func ObjectDiv(a Object, b Object) Object {
	c := NewNumber(a.Number() / b.Number())
	return c
}

// ObjectDivInt return a / b
func ObjectDivInt(a Object, b Object) Object {
	c := NewInt(a.Int() / b.Int())
	return c
}

// ObjectMod return a % b
func ObjectMod(a Object, b Object) Object {
	c := NewInt(a.Int() % b.Int())
	return c
}
