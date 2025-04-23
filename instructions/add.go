package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/registers"
)

type AddShiftedRegister uint32

func NewAddOrAddsShiftedRegister(
	Xd registers.GPRegister,
	Xn registers.GPRegister,
	Xm registers.GPRegister,
	setFlags immediates.SetFlags,
) AddShiftedRegister {
	return AddShiftedRegister(
		0x8B000000 |
			(setFlags.Binary() << 29) |
			(Xm.Binary() << 16) |
			(Xn.Binary() << 5) |
			(Xd.Binary()),
	)
}

func NewAddShiftedRegister(
	Xd registers.GPRegister,
	Xn registers.GPRegister,
	Xm registers.GPRegister,
) AddShiftedRegister {
	return NewAddOrAddsShiftedRegister(Xd, Xn, Xm, immediates.DoNotSetFlags)
}

func NewAddsShiftedRegister(
	Xd registers.GPRegister,
	Xn registers.GPRegister,
	Xm registers.GPRegister,
) AddShiftedRegister {
	return NewAddOrAddsShiftedRegister(Xd, Xn, Xm, immediates.DoSetFlags)
}

func (i AddShiftedRegister) Binary() uint32 {
	return uint32(i)
}

func (i AddShiftedRegister) Xd() registers.GPRegister {
	return registers.GPRegister(i & 0x1F)
}

func (i AddShiftedRegister) Xn() registers.GPRegister {
	return registers.GPRegister((i >> 5) & 0x1F)
}

func (i AddShiftedRegister) Xm() registers.GPRegister {
	return registers.GPRegister((i >> 16) & 0x1F)
}

func (i AddShiftedRegister) SetFlags() immediates.SetFlags {
	return immediates.SetFlagsFromBinary(uint32(i))
}

func (i AddShiftedRegister) String() string {
	return fmt.Sprintf("add%s %s, %s, %s", i.SetFlags(), i.Xd(), i.Xn(), i.Xm())
}

type AddImmediate uint32

func NewAddImmediate(
	Xd registers.GPorSPRegister,
	Xn registers.GPorSPRegister,
	imm immediates.Immediate12,
) AddImmediate {
	return AddImmediate(
		0x91000000 |
			(imm.Binary() << 10) |
			(Xn.Binary() << 5) |
			(Xd.Binary()),
	)
}

func (i AddImmediate) Binary() uint32 {
	return uint32(i)
}

func (i AddImmediate) Xd() registers.GPorSPRegister {
	return registers.GPorSPRegister(i & 0x1F)
}

func (i AddImmediate) Xn() registers.GPorSPRegister {
	return registers.GPorSPRegister((i >> 5) & 0x1F)
}

func (i AddImmediate) Imm() immediates.Immediate12 {
	return immediates.Immediate12((i >> 10) & 0xFFF)
}

func (i AddImmediate) String() string {
	return fmt.Sprintf("add %s, %s, %s", i.Xd(), i.Xn(), i.Imm())
}

type AddsImmediate uint32

func NewAddsImmediate(
	Xd registers.GPRegister,
	Xn registers.GPorSPRegister,
	imm immediates.Immediate12,
) AddsImmediate {
	return AddsImmediate(
		0xb1000000 |
			(imm.Binary() << 10) |
			(Xn.Binary() << 5) |
			(Xd.Binary()),
	)
}

func (i AddsImmediate) Binary() uint32 {
	return uint32(i)
}

func (i AddsImmediate) Xd() registers.GPRegister {
	return registers.GPRegister(i & 0x1F)
}

func (i AddsImmediate) Xn() registers.GPorSPRegister {
	return registers.GPorSPRegister((i >> 5) & 0x1F)
}

func (i AddsImmediate) Immediate() immediates.Immediate12 {
	return immediates.Immediate12((i >> 10) & 0xFFF)
}

func (i AddsImmediate) String() string {
	return fmt.Sprintf("adds %s, %s, %s", i.Xd(), i.Xn(), i.Immediate())
}
