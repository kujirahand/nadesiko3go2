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

func TestVMPlusMul(t *testing.T) {
	vm := NewVM(nil)
	const1 := vm.AddConst(&TNumber{value: 1})
	const2 := vm.AddConst(&TNumber{value: 2})
	const3 := vm.AddConst(&TNumber{value: 3})
	codes := []VMCode{
		NewCode(TypeOpConst, const2, 0, 0),
		NewCode(TypeOpConst, const3, 0, 0),
		NewCode(TypeOpMul, 0, 0, 0),
		NewCode(TypeOpConst, const1, 0, 0),
		NewCode(TypeOpPlus, 0, 0, 0),
	}
	res, _ := vm.Eval(codes)
	if res.Number() != 7.0 {
		t.Errorf("TestVMMulPlus %s != 7.0", res.String())
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

func TestVMLEq(t *testing.T) {
	vm := NewVM(nil)
	const1 := vm.AddConst(&TNumber{value: 1})
	const2 := vm.AddConst(&TNumber{value: 2})
	const3 := vm.AddConst(&TNumber{value: 3})
	codes := []VMCode{
		NewCode(TypeOpConst, const1, 0, 0),
		NewCode(TypeOpSetLocal, 1, 0, 0),
		NewCode(TypeOpConst, const2, 0, 0),
		NewCode(TypeOpSetLocal, 2, 0, 0),
		NewCode(TypeOpGetLocal, 1, 0, 0),
		NewCode(TypeOpGetLocal, 2, 0, 0),
		NewCode(TypeOpPlus, 0, 0, 0),
		NewCode(TypeOpConst, const3, 0, 0),
		NewCode(TypeOpEq, const3, 0, 0),
	}
	res, _ := vm.Eval(codes)
	if res.Number() != 1.0 {
		t.Errorf("TestVMEq %s != 1", res.String())
	}
}

func TestVMLikeFor(t *testing.T) {
	vm := NewVM(nil)
	const1 := vm.AddConst(&TNumber{value: 1})
	codes := []VMCode{
		// i=1
		NewCode(TypeOpConst, const1, 0, 0),
		NewCode(TypeOpSetLocal, 1, 0, 0),
		// j=1
		NewCode(TypeOpPushInt, 0, 0, 0),
		NewCode(TypeOpSetLocal, 2, 0, 0),
		// j+=i
		NewCode(TypeOpGetLocal, 2, 0, 0),
		NewCode(TypeOpGetLocal, 1, 0, 0),
		NewCode(TypeOpPlus, 0, 0, 0),
		NewCode(TypeOpSetLocal, 2, 0, 0),
		// print j
		NewCode(TypeOpGetLocal, 2, 0, 0),
		NewCode(TypeOpPrint, 0, 0, 0),
		// i++
		NewCode(TypeOpIncLocal, 1, 0, 0),
		// if i <= 10: jump top
		NewCode(TypeOpGetLocal, 1, 0, 0),
		NewCode(TypeOpPushInt, 10, 0, 0),
		NewCode(TypeOpLtEq, 0, 0, 0),
		NewCode(TypeOpJumpAddrTrue, 4, 0, 0),
		// push j
		NewCode(TypeOpGetLocal, 2, 0, 0),
	}
	res, _ := vm.Eval(codes)
	if res.Number() != 55.0 {
		t.Errorf("TestVMFor %s != 55.0", res.String())
	}
}

func _evalVM(t *testing.T, codes []VMCode) {
	env := NewEnv()
	vm := NewVM(env)
	vm.Eval(codes)
}
