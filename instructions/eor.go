package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/registers"
)

// EOR instruction (logical XOR)
type Eor uint32

// EOR creates an EOR instruction with a shifted register
func EOR(
	Xd registers.GPRegister,
	Xn registers.GPRegister,
	Xm registers.GPRegister,
	shift immediates.Shift6,
) Eor {
	return Eor(
		0xCA000000 |
			(shift.Type.Binary() << 22) |
			(shift.Amount.Binary() << 10) |
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

// Shift returns the shift details for this instruction
func (i Eor) Shift() immediates.Shift6 {
	return immediates.NewShift6(
		immediates.ShiftType((i>>22)&0x3),
		immediates.Immediate6((i>>10)&0x3F),
	)
}

func (i Eor) String() string {
	s := fmt.Sprintf("eor %s, %s, %s", i.Xd(), i.Xn(), i.Xm())
	shift := i.Shift()
	if shift.HasShift() {
		s += fmt.Sprintf(", %s", shift)
	}
	return s
}
