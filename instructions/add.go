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
) Add {
	return Add(
		0x8B000000 |
			(Xm.Binary() << 16) |
			(Xn.Binary() << 5) |
			(Xd.Binary()),
	)
}

func (i Add) Binary() uint32 {
	return uint32(i)
}

func (i Add) String() string {
	Xd := registers.GPRegister(i & 0x1F)
	Xn := registers.GPRegister((i >> 5) & 0x1F)
	Xm := registers.GPRegister((i >> 16) & 0x1F)
	return fmt.Sprintf("ADD %s, %s, %s", Xd, Xn, Xm)
}

// ADDI instruction (ADD with immediate)
type AddImm uint32

func ADDI(
	Xd registers.GPorSPRegister,
	Xn registers.GPorSPRegister,
	imm immediates.Immediate12,
) AddImm {
	return AddImm(
		0x91000000 |
			(imm.Binary() << 10) |
			(Xn.Binary() << 5) |
			(Xd.Binary()),
	)
}

func (i AddImm) Binary() uint32 {
	return uint32(i)
}

func (i AddImm) String() string {
	Xd := registers.GPorSPRegister(i & 0x1F)
	Xn := registers.GPorSPRegister((i >> 5) & 0x1F)
	imm := immediates.Immediate12((i >> 10) & 0xFFF)
	return fmt.Sprintf("ADD %s, %s, %s", Xd, Xn, imm)
}
