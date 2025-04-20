package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/registers"
)

type Ldr uint32

func LDR(
	Xt registers.GPRegister,
	Xn registers.GPorSPRegister,
	imm immediates.Immediate9,
) Ldr {
	return Ldr(
		0xF8400400 |
			(imm.Binary() << 12) |
			(Xn.Binary() << 5) |
			(Xt.Binary()),
	)
}

func (i Ldr) Binary() uint32 {
	return uint32(i)
}

func (i Ldr) Xt() registers.GPRegister {
	return registers.GPRegister(i & 0x1F)
}

func (i Ldr) Xn() registers.GPorSPRegister {
	return registers.GPorSPRegister((i >> 5) & 0x1F)
}

func (i Ldr) Imm() immediates.Immediate9 {
	return immediates.Immediate9((i >> 12) & 0x1FF)
}

func (i Ldr) String() string {
	s := fmt.Sprintf("ldr %s, [%s]", i.Xt(), i.Xn())
	imm := i.Imm()
	if imm != 0 {
		s += ", " + imm.String()
	}
	return s
}
