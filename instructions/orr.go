package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/registers"
)

// ORR instruction (logical OR)
type Orr uint32

func ORR(
	Xd registers.GPRegister,
	Xn registers.GPRegister,
	Xm registers.GPRegister,
) Orr {
	return Orr(
		0xAA000000 |
			(Xm.Binary() << 16) |
			(Xn.Binary() << 5) |
			(Xd.Binary()),
	)
}

func (i Orr) Binary() uint32 {
	return uint32(i)
}

func (i Orr) Xd() registers.GPRegister {
	return registers.GPRegister(i & 0x1F)
}

func (i Orr) Xn() registers.GPRegister {
	return registers.GPRegister((i >> 5) & 0x1F)
}

func (i Orr) Xm() registers.GPRegister {
	return registers.GPRegister((i >> 16) & 0x1F)
}

func (i Orr) String() string {
	return fmt.Sprintf("ORR %s, %s, %s", i.Xd(), i.Xn(), i.Xm())
}
