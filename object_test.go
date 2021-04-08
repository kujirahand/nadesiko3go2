package nadesiko4_test

import (
	"testing"

	"github.com/kujirahand/nadesiko4"
)

func TestObject1(t *testing.T) {
	s1 := nadesiko4.String{Value: "hoge"}
	if s1.Value != "hoge" {
		t.Error("error")
	}
	s2 := s1.Clone()
	if s1.Value != s2.String() {
		t.Error("error")
	}
	v2 := nadesiko4.Object(s2).(*nadesiko4.String).Value
	if s1.Value != v2 {
		t.Error("error")
	}
	null1 := nadesiko4.NullValue
	null2 := &nadesiko4.Null{}
	if !null1.StrictEqual(null2) {
		t.Error("Nullが一致しません")
	}
}

func TestObjectArray1(t *testing.T) {
	av := []nadesiko4.Object{&nadesiko4.Number{Value: 1.0}, &nadesiko4.Number{Value: 2}, &nadesiko4.Number{Value: 3}}
	a1 := nadesiko4.Array{Value: av}
	if len(a1.Value) != 3 {
		t.Error("array len wrong")
	}
}

func TestObjectArray2(t *testing.T) {
	v := nadesiko4.NewArray()
	v.SetIndex(0, &nadesiko4.Number{Value: 30})
	if v.Length() != 1 {
		t.Error("array len wrong")
	}
}
func TestObjectArray3(t *testing.T) {
	a1 := nadesiko4.NewArray()
	a2 := nadesiko4.NewArray()
	a1.SetIndex(0, nadesiko4.NewNumber(10))
	a1.SetIndex(1, nadesiko4.NewNumber(20))
	a2.SetIndex(0, nadesiko4.NewNumber(10))
	a2.SetIndex(1, nadesiko4.NewNumber(20))
	if !a1.StrictEqual(a2) {
		t.Error("array not mutch")
	}
	a2.SetIndex(2, nadesiko4.NewNumber(3))
	if a1.StrictEqual(a2) {
		t.Error("different array mutch!!")
	}
}

func TestObjectDict1(t *testing.T) {
	dict := nadesiko4.NewDict()
	dict.SetDict("test", nadesiko4.NewNumber(30))
	v := dict.GetDict("test")
	if v == nil || v.Int() != 30 {
		t.Errorf("dict error, %s", v)
	}
}

func TestObjectNumber(t *testing.T) {
	a := &nadesiko4.Number{Value: 3}
	b := &nadesiko4.Number{Value: 3}
	if a.Value != b.Value {
		t.Errorf("number error, %f", a.Value)
	}
	if !a.Equal(b) {
		t.Errorf("number error, %f", b.Value)
	}
	c := nadesiko4.NewNumber(10)
	d := nadesiko4.NewNumber(10)
	if !c.StrictEqual(d) {
		t.Errorf("number error, %f", c.Number())
	}
}

func TestObjectNumberCalc(t *testing.T) {
	a := &nadesiko4.Number{Value: 3}
	b := &nadesiko4.Number{Value: 5}
	c := nadesiko4.ObjectAdd(a, b)
	if c.Number() != 8.0 {
		t.Errorf("number error, %f", a.Value)
	}
	if !c.Equal(nadesiko4.NewNumber(8)) {
		t.Errorf("number error, %f", a.Value)
	}
	a2 := nadesiko4.NewInt(3)
	b2 := nadesiko4.NewInt(5)
	d := nadesiko4.ObjectAddStr(a2, b2)
	if d.String() != "35" {
		t.Errorf("ObjectAddStr error, %s", d.String())
	}
	e := nadesiko4.ObjectMul(a, b)
	if e.Number() != 15 {
		t.Errorf("ObjectMul error, %s", e.String())
	}

}
