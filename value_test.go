package nadesiko4

import (
	"testing"
)

func TestValueDef(t *testing.T) {
	var s string
	s = GetValueTypeName(NilValue.Type())
	if s != "nil" {
		t.Errorf("nil type name error")
	}
}

func TestValueBool(t *testing.T) {
	v1 := &TBool{value: true}
	if v1.Number() != 1.0 {
		t.Errorf("bool type error")
	}
}

func TestValue(t *testing.T) {
	v := &TNil{}
	if v.Type() != TypeNil {
		t.Errorf("TNil error")
	}
	b := &TBool{value: true}
	bs := b.String()
	if bs != "1" {
		t.Errorf("TBool.String %s != 1", bs)
	}
	n := &TNumber{value: 32}
	ns := n.String()
	if ns != "32" {
		t.Errorf("TNumber.String %s != 32", ns)
	}
	n2 := &TNumber{value: 3.1415}
	n2s := n2.String()
	if n2s != "3.1415" {
		t.Errorf("TNumber String %s != 3.1415", n2s)
	}
}
