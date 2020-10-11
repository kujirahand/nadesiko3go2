package nadesiko4

// VMCodeType is type of VM Opcode
type VMCodeType int

// VMCodeType list [params] [stack] [return]
const (
	TypeOpNop      VMCodeType = iota //
	TypeOpConst                      // push(constants[A])
	TypeOpJump                       // jump(+A)
	TypeOpJumpTrue                   // if(stackTop){ jump(+A) }
	TypeOpGetLocal                   // push(local[A])
	TypeOpSetLocal                   // local[A] = pop()
	TypeOpPlus                       // push(pop() + pop())
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
