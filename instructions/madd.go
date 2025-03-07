package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/registers"
)

// MADD (Multiply-Add) instruction: Xd = Xa + Xn * Xm
type Madd uint32

func MADD(
	Xd registers.GPRegister,
	Xn registers.GPRegister,
	Xm registers.GPRegister,
	Xa registers.GPRegister,
) Madd {
	return Madd(
		0x9B000000 |
			(Xm.Binary() << 16) |
			(Xa.Binary() << 10) |
			(Xn.Binary() << 5) |
			(Xd.Binary()),
	)
}

func (i Madd) Binary() uint32 {
	return uint32(i)
}

func (i Madd) Xd() registers.GPRegister {
	return registers.GPRegister(i & 0x1F)
}

func (i Madd) Xn() registers.GPRegister {
	return registers.GPRegister((i >> 5) & 0x1F)
}

func (i Madd) Xa() registers.GPRegister {
	return registers.GPRegister((i >> 10) & 0x1F)
}

func (i Madd) Xm() registers.GPRegister {
	return registers.GPRegister((i >> 16) & 0x1F)
}

func (i Madd) String() string {
	return fmt.Sprintf("MADD %s, %s, %s, %s", i.Xd(), i.Xn(), i.Xm(), i.Xa())
}
