package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/registers"
)

// EOR instruction (logical XOR)
type Eor uint32

func EOR(
	Xd registers.GPRegister,
	Xn registers.GPRegister,
	Xm registers.GPRegister,
) Eor {
	return Eor(
		0xCA000000 |
			(Xm.Binary() << 16) |
			(Xn.Binary() << 5) |
			(Xd.Binary()),
	)
}

func (i Eor) Binary() uint32 {
	return uint32(i)
}

func (i Eor) Xd() registers.GPRegister {
	return registers.GPRegister(i & 0x1F)
}

func (i Eor) Xn() registers.GPRegister {
	return registers.GPRegister((i >> 5) & 0x1F)
}

func (i Eor) Xm() registers.GPRegister {
	return registers.GPRegister((i >> 16) & 0x1F)
}

func (i Eor) String() string {
	return fmt.Sprintf("EOR %s, %s, %s", i.Xd(), i.Xn(), i.Xm())
}
