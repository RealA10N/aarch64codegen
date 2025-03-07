package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/registers"
)

// MUL instruction (note: MUL is actually an alias for MADD with XZR)
type Mul Madd

func MUL(
	Xd registers.GPRegister,
	Xn registers.GPRegister,
	Xm registers.GPRegister,
) Mul {
	return Mul(MADD(Xd, Xn, Xm, registers.XZR))
}

func (i Mul) Binary() uint32 {
	return uint32(i)
}

func (i Mul) Xd() registers.GPRegister {
	return Madd(i).Xd()
}

func (i Mul) Xn() registers.GPRegister {
	return Madd(i).Xn()
}

func (i Mul) Xm() registers.GPRegister {
	return Madd(i).Xm()
}

func (i Mul) String() string {
	return fmt.Sprintf("MUL %s, %s, %s", i.Xd(), i.Xn(), i.Xm())
}
