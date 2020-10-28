package nadesiko4

import (
	"fmt"
	"math"
)

// Frame for function
type Frame struct {
	locals []Value
	names  map[string]int
}

// NewFrame returns new frame
func NewFrame() *Frame {
	f := Frame{}
	f.locals = []Value{NilValue}
	f.names = map[string]int{
		"それ": 0,
	}
	return &f
}

// SetLocal set value to locals
func (f *Frame) SetLocal(index int, v Value) {
	for len(f.locals) <= index {
		f.locals = append(f.locals, NilValue)
	}
	f.locals[index] = v
}

// VM is Virtual Macine struct
type VM struct {
	env        *Env
	codes      []VMCode // bytecodes
	ci         int      // bytecode index
	constants  []Value
	sp         int // stack pointer
	stack      [MaxStack]Value
	farmeIndex int // fame index
	frames     [MaxFrame]*Frame
	curFrame   *Frame
	aborted    bool
}

// NewVM creates VM strcut
func NewVM(env *Env) *VM {
	if env == nil {
		env = NewEnv()
	}
	vm := &VM{
		env:       env,
		sp:        0,
		ci:        0,
		constants: []Value{},
		aborted:   false,
	}
	vm.curFrame = NewFrame()
	vm.frames[0] = vm.curFrame
	return vm
}

// AddConst constants value
func (vm *VM) AddConst(v Value) int {
	index := len(vm.constants)
	vm.constants = append(vm.constants, v)
	return index
}

// Eval execute bytecodes
func (vm *VM) Eval(codes []VMCode) (Value, error) {
	vm.codes = codes
	vm.ci = 0
	return vm.run()
}

// run
func (vm *VM) run() (Value, error) {
	for !vm.aborted {
		if len(vm.codes) <= vm.ci {
			break
		}
		code := vm.codes[vm.ci]
		fmt.Printf("%3d: %s\n", vm.ci, getVMCodeName(code.Type))
		switch code.Type {
		case TypeOpNop:
			// NOP
		case TypeOpJumpAddr:
			vm.ci = code.A
			continue // 自動で＋１しないように
		case TypeOpJumpAddrTrue:
			cond := vm.stack[vm.sp-1]
			vm.sp--
			if cond.Number() > 0 {
				vm.ci = code.A
				continue // 自動で＋１しないように
			}
		case TypeOpJump:
			vm.ci += code.A
			continue // 自動で＋１しないように
		case TypeOpJumpTrue:
			cond := vm.stack[vm.sp-1]
			vm.sp--
			if cond.Number() > 0 {
				vm.ci += code.A
				continue // 自動で＋１しないように
			}
		case TypeOpConst:
			vm.stack[vm.sp] = vm.constants[code.A]
			vm.sp++
		case TypeOpPushInt:
			vm.stack[vm.sp] = &TNumber{value: float64(code.A)}
			vm.sp++
		case TypeOpIncLocal:
			v := vm.curFrame.locals[code.A]
			vm.curFrame.locals[code.A] = &TNumber{value: v.Number() + 1.0}
		case TypeOpDecLocal:
			v := vm.curFrame.locals[code.A]
			vm.curFrame.locals[code.A] = &TNumber{value: v.Number() - 1}
		case TypeOpGetLocal:
			vm.stack[vm.sp] = vm.curFrame.locals[code.A]
			vm.sp++
		case TypeOpSetLocal:
			v := vm.stack[vm.sp-1]
			vm.sp--
			vm.curFrame.SetLocal(code.A, v)
		case TypeOpPlus:
			valR, valL := vm.stack[vm.sp-1], vm.stack[vm.sp-2]
			vm.sp -= 2
			vm.stack[vm.sp] = &TNumber{value: (valL.Number() + valR.Number())}
			vm.sp++
		case TypeOpPlusStr:
			valR, valL := vm.stack[vm.sp-1], vm.stack[vm.sp-2]
			vm.sp -= 2
			vm.stack[vm.sp] = &TString{value: (valL.String() + valR.String())}
			vm.sp++
		case TypeOpMinus:
			valR, valL := vm.stack[vm.sp-1], vm.stack[vm.sp-2]
			vm.sp -= 2
			vm.stack[vm.sp] = &TNumber{value: (valL.Number() - valR.Number())}
			vm.sp++
		case TypeOpMul:
			valR, valL := vm.stack[vm.sp-1], vm.stack[vm.sp-2]
			vm.sp -= 2
			vm.stack[vm.sp] = &TNumber{value: (valL.Number() * valR.Number())}
			vm.sp++
		case TypeOpDiv:
			valR, valL := vm.stack[vm.sp-1], vm.stack[vm.sp-2]
			vm.sp -= 2
			vm.stack[vm.sp] = &TNumber{value: (valL.Number() / valR.Number())}
			vm.sp++
		case TypeOpMod:
			valR, valL := vm.stack[vm.sp-1], vm.stack[vm.sp-2]
			vm.sp -= 2
			vm.stack[vm.sp] = &TNumber{value: math.Mod(valL.Number(), valR.Number())}
			vm.sp++
		case TypeOpEq, TypeOpNtEq, TypeOpGt, TypeOpGtEq, TypeOpLt, TypeOpLtEq:
			valR, valL := vm.stack[vm.sp-1], vm.stack[vm.sp-2]
			vm.sp -= 2
			vm.stack[vm.sp] = &TBool{value: valL.Compare(code.Type, valR)}
			vm.sp++
		case TypeOpPrint:
			v := vm.stack[vm.sp-1]
			vm.sp--
			println(">>> ", v.String())
		default:
			panic(fmt.Errorf("Invalid OpCode : %d", int(code.Type)))
		}
		vm.ci++
	}
	if vm.sp >= 1 {
		res := vm.stack[vm.sp-1]
		vm.sp--
		return res, nil
	}
	return NilValue, nil
}
