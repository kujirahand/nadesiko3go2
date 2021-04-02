package nadesiko4

import "testing"

func TestVMNop(t *testing.T) {
	vm := NewVM(nil)
	codes := []VMCode{
		NewCode(TypeOpNop, 0, 0, 0),
	}
	vm.Eval(codes)
}

func TestVMPlus(t *testing.T) {
	vm := NewVM(nil)
	const1 := vm.AddConst(&TNumber{value: 1})
	const2 := vm.AddConst(&TNumber{value: 2})
	codes := []VMCode{
		NewCode(TypeOpConst, const1, 0, 0),
		NewCode(TypeOpConst, const2, 0, 0),
		NewCode(TypeOpPlus, 0, 0, 0),
	}
	res, _ := vm.Eval(codes)
	if res.Number() != 3.0 {
		t.Errorf("TestVMPlus %s != 3.0", res.String())
	}
}

func TestVMLocal(t *testing.T) {
	vm := NewVM(nil)
	const1 := vm.AddConst(&TNumber{value: 1})
	const2 := vm.AddConst(&TNumber{value: 2})
	codes := []VMCode{
		NewCode(TypeOpConst, const1, 0, 0),
		NewCode(TypeOpSetLocal, 1, 0, 0),
		NewCode(TypeOpConst, const2, 0, 0),
		NewCode(TypeOpSetLocal, 2, 0, 0),
		NewCode(TypeOpGetLocal, 2, 0, 0),
		NewCode(TypeOpGetLocal, 1, 0, 0),
		NewCode(TypeOpPlus, 0, 0, 0),
	}
	res, _ := vm.Eval(codes)
	if res.Number() != 3.0 {
		t.Errorf("TestVMPlus %s != 3.0", res.String())
	}
}

func TestVMLikeFor(t *testing.T) {
	vm := NewVM(nil)
	const1 := vm.AddConst(&TNumber{value: 1})
	const2 := vm.AddConst(&TNumber{value: 2})
	codes := []VMCode{
		NewCode(TypeOpConst, const1, 0, 0),
		NewCode(TypeOpSetLocal, 1, 0, 0),
		NewCode(TypeOpConst, const2, 0, 0),
		NewCode(TypeOpSetLocal, 2, 0, 0),
		NewCode(TypeOpGetLocal, 2, 0, 0),
		NewCode(TypeOpGetLocal, 1, 0, 0),
		NewCode(TypeOpPlus, 0, 0, 0),
	}
	res, _ := vm.Eval(codes)
	if res.Number() != 3.0 {
		t.Errorf("TestVMPlus %s != 3.0", res.String())
	}
}
