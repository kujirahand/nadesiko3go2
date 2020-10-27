package nadesiko4

// VMCodeType is type of VM Opcode
type VMCodeType int

// VMCodeType list [params] [stack] [return]
const (
	// VMCode.Begin
	TypeOpNop       VMCodeType = iota //
	TypeOpConst                       // push(constants[A])
	TypeOpPushInt                     // push(A)
	TypeOpJump                        // jump(+A)
	TypeOpJumpTrue                    // if(stackTop){ jump(+A) }
	TypeOpGetLocal                    // push(local[A])
	TypeOpSetLocal                    // local[A] = pop()
	TypeOpGetGlobal                   // push(global[const[A]])
	TypeOpSetGlobal                   // global[A] = pop()
	TypeOpPlus                        // push(pop() + pop())
	TypeOpMinus
	TypeOpMul
	TypeOpDiv
	TypeOpMod
	TypeOpPlusStr
	TypeOpEq
	TypeOpNtEq
	TypeOpGt
	TypeOpGtEq
	TypeOpLt
	TypeOpLtEq
	TypeOpIncLocal // local[A]++
	TypeOpDecLocal // local[A]--
	// VMCode.End
)

// VMCode strcut
type VMCode struct {
	Type VMCodeType
	A    int
	B    int
	C    int
}

// NewCode returns VMCode
func NewCode(t VMCodeType, A int, B int, C int) VMCode {
	return VMCode{t, A, B, C}
}
