package nadesiko4

// VMCodeType is type of VM Opcode
type VMCodeType int

// VMCodeType list [params] [stack] [return]
const (
	// __begin__
	TypeOpNop          VMCodeType = iota //
	TypeOpConst                          // push(constants[A])
	TypeOpPushInt                        // push(A)
	TypeOpJump                           // jump(ci+A)
	TypeOpJumpTrue                       // if(pop()){ jump(ci+A) }
	TypeOpJumpAddr                       // jump(A)
	TypeOpJumpAddrTrue                   // if(pop()) { jump(A) }
	TypeOpGetLocal                       // push(local[A])
	TypeOpSetLocal                       // local[A] = pop()
	TypeOpGetGlobal                      // push(global[const[A]])
	TypeOpSetGlobal                      // global[A] = pop()
	TypeOpPlus                           // push(pop() + pop())
	TypeOpMinus                          //
	TypeOpMul                            //
	TypeOpDiv                            //
	TypeOpMod                            //
	TypeOpPlusStr                        //
	TypeOpEq                             //
	TypeOpNtEq                           //
	TypeOpGt                             //
	TypeOpGtEq                           //
	TypeOpLt                             //
	TypeOpLtEq                           //
	TypeOpIncLocal                       // local[A]++
	TypeOpDecLocal                       // local[A]--
	TypeOpPrint                          // print(pop())
	// __end__
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
