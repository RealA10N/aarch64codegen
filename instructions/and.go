package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/registers"
)

// AND instruction (logical AND)
type And uint32

// AND creates an AND instruction with an optional shifted register
func AND(
	Xd registers.GPRegister,
	Xn registers.GPRegister,
	Xm registers.GPRegister,
	shift immediates.Shift6,
) And {
	return And(
		0x8A000000 |
			(shift.Type.Binary() << 22) |
			(shift.Amount.Binary() << 10) |
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

// Shift returns the shift details for this instruction
func (i And) Shift() immediates.Shift6 {
	return immediates.NewShift6(
		immediates.ShiftType((i>>22)&0x3),
		immediates.Immediate6((i>>10)&0x3F),
	)
}

func (i And) String() string {
	s := fmt.Sprintf("AND %s, %s, %s", i.Xd(), i.Xn(), i.Xm())
	shift := i.Shift()
	if shift.HasShift() {
		s += fmt.Sprintf(", %s", shift)
	}
	return s
}
