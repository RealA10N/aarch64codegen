package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/registers"
)

// SUB instruction with register operands
type Sub uint32

func SUB(
	Xd registers.GPRegister,
	Xn registers.GPRegister,
	Xm registers.GPRegister,
) Sub {
	return Sub(
		0xCB000000 |
			(Xm.Binary() << 16) |
			(Xn.Binary() << 5) |
			(Xd.Binary()),
	)
}

func (i Sub) Binary() uint32 {
	return uint32(i)
}

func (i Sub) Xd() registers.GPRegister {
	return registers.GPRegister(i & 0x1F)
}

func (i Sub) Xn() registers.GPRegister {
	return registers.GPRegister((i >> 5) & 0x1F)
}

func (i Sub) Xm() registers.GPRegister {
	return registers.GPRegister((i >> 16) & 0x1F)
}

func (i Sub) String() string {
	return fmt.Sprintf("sub %s, %s, %s", i.Xd(), i.Xn(), i.Xm())
}

// SUBI instruction (SUB with immediate)
type SubImm uint32

func SUBI(
	Xd registers.GPRegister,
	Xn registers.GPorSPRegister,
	imm immediates.Immediate12,
) SubImm {
	return SubImm(
		0xD1000000 |
			(imm.Binary() << 10) |
			(Xn.Binary() << 5) |
			(Xd.Binary()),
	)
}

func (i SubImm) Binary() uint32 {
	return uint32(i)
}

func (i SubImm) Xd() registers.GPRegister {
	return registers.GPRegister(i & 0x1F)
}

func (i SubImm) Xn() registers.GPorSPRegister {
	return registers.GPorSPRegister((i >> 5) & 0x1F)
}

func (i SubImm) Imm() immediates.Immediate12 {
	return immediates.Immediate12((i >> 10) & 0xFFF)
}

func (i SubImm) String() string {
	return fmt.Sprintf("sub %s, %s, %s", i.Xd(), i.Xn(), i.Imm())
}
