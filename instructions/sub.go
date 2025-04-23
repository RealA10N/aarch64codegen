package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/registers"
)

// SUB instruction with register operands
type Sub uint32

func NewSubShiftedRegister(
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

func NewSubsShiftedRegister(
	Xd registers.GPRegister,
	Xn registers.GPRegister,
	Xm registers.GPRegister,
) Sub {
	return NewSubShiftedRegister(Xd, Xn, Xm) | (1 << 29)
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

func (i Sub) SetFlags() immediates.SetFlags {
	return immediates.SetFlagsFromBinary(uint32(i))
}

func (i Sub) String() string {
	return fmt.Sprintf("sub%s %s, %s, %s", i.SetFlags(), i.Xd(), i.Xn(), i.Xm())
}

type SubImmediate uint32

func NewSubImmediate(
	Xd registers.GPRegister,
	Xn registers.GPorSPRegister,
	imm immediates.Immediate12,
) SubImmediate {
	return SubImmediate(
		0xD1000000 |
			(imm.Binary() << 10) |
			(Xn.Binary() << 5) |
			(Xd.Binary()),
	)
}

func NewSubsImmediate(
	Xd registers.GPRegister,
	Xn registers.GPorSPRegister,
	imm immediates.Immediate12,
) SubImmediate {
	return NewSubImmediate(Xd, Xn, imm) | (1 << 29)
}

func (i SubImmediate) Binary() uint32 {
	return uint32(i)
}

func (i SubImmediate) Xd() registers.GPRegister {
	return registers.GPRegister(i & 0x1F)
}

func (i SubImmediate) Xn() registers.GPorSPRegister {
	return registers.GPorSPRegister((i >> 5) & 0x1F)
}

func (i SubImmediate) Imm() immediates.Immediate12 {
	return immediates.Immediate12((i >> 10) & 0xFFF)
}

func (i SubImmediate) SetFlags() immediates.SetFlags {
	return immediates.SetFlagsFromBinary(uint32(i))
}

func (i SubImmediate) String() string {
	return fmt.Sprintf("sub%s %s, %s, %s", i.SetFlags(), i.Xd(), i.Xn(), i.Imm())
}
