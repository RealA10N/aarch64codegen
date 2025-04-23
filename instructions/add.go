// filepath: /Users/alonkr/Developer/aarch64codegen/instructions/add.go
package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/registers"
)

// ADD instruction with register operands
type Add uint32

func ADD(
	Xd registers.GPRegister,
	Xn registers.GPRegister,
	Xm registers.GPRegister,
	setFlags immediates.SetFlags,
) Add {
	return Add(
		0x8B000000 |
			(setFlags.Binary() << 29) |
			(Xm.Binary() << 16) |
			(Xn.Binary() << 5) |
			(Xd.Binary()),
	)
}

func (i Add) Binary() uint32 {
	return uint32(i)
}

func (i Add) Xd() registers.GPRegister {
	return registers.GPRegister(i & 0x1F)
}

func (i Add) Xn() registers.GPRegister {
	return registers.GPRegister((i >> 5) & 0x1F)
}

func (i Add) Xm() registers.GPRegister {
	return registers.GPRegister((i >> 16) & 0x1F)
}

func (i Add) SetFlags() immediates.SetFlags {
	return immediates.SetFlagsFromBinary(uint32(i))
}

func (i Add) String() string {
	return fmt.Sprintf("add%s %s, %s, %s", i.SetFlags(), i.Xd(), i.Xn(), i.Xm())
}

// ADDI instruction (ADD with immediate)
type AddImm uint32

func ADDI(
	Xd registers.GPorSPRegister,
	Xn registers.GPorSPRegister,
	imm immediates.Immediate12,
	setFlags immediates.SetFlags,
) AddImm {
	return AddImm(
		0x91000000 |
			(setFlags.Binary()) |
			(imm.Binary() << 10) |
			(Xn.Binary() << 5) |
			(Xd.Binary()),
	)
}

func (i AddImm) Binary() uint32 {
	return uint32(i)
}

func (i AddImm) Xd() registers.GPorSPRegister {
	return registers.GPorSPRegister(i & 0x1F)
}

func (i AddImm) Xn() registers.GPorSPRegister {
	return registers.GPorSPRegister((i >> 5) & 0x1F)
}

func (i AddImm) Imm() immediates.Immediate12 {
	return immediates.Immediate12((i >> 10) & 0xFFF)
}

func (i AddImm) SetFlags() immediates.SetFlags {
	return immediates.SetFlagsFromBinary(uint32(i))
}

func (i AddImm) String() string {
	return fmt.Sprintf("add%s %s, %s, %s", i.SetFlags(), i.Xd(), i.Xn(), i.Imm())
}
