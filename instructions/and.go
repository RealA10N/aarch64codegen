package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/registers"
)

// AND instruction (logical AND)
type And uint32

func AND(
	Xd registers.GPRegister,
	Xn registers.GPRegister,
	Xm registers.GPRegister,
) And {
	return And(
		0x8A000000 |
			(Xm.Binary() << 16) |
			(Xn.Binary() << 5) |
			(Xd.Binary()),
	)
}

func (i And) Binary() uint32 {
	return uint32(i)
}

func (i And) Xd() registers.GPRegister {
	return registers.GPRegister(i & 0x1F)
}

func (i And) Xn() registers.GPRegister {
	return registers.GPRegister((i >> 5) & 0x1F)
}

func (i And) Xm() registers.GPRegister {
	return registers.GPRegister((i >> 16) & 0x1F)
}

func (i And) String() string {
	return fmt.Sprintf("AND %s, %s, %s", i.Xd(), i.Xn(), i.Xm())
}
