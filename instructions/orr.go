package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/registers"
)

// ORR instruction (logical OR)
type Orr uint32

// ORR creates an ORR instruction with an optional shifted register
func ORR(
	Xd registers.GPRegister,
	Xn registers.GPRegister,
	Xm registers.GPRegister,
	shift immediates.Shift6,
) Orr {
	return Orr(
		0xAA000000 |
			(shift.Type.Binary() << 22) |
			(shift.Amount.Binary() << 10) |
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

// Shift returns the shift details for this instruction
func (i Orr) Shift() immediates.Shift6 {
	return immediates.NewShift6(
		immediates.ShiftType((i>>22)&0x3),
		immediates.Immediate6((i>>10)&0x3F),
	)
}

func (i Orr) String() string {
	s := fmt.Sprintf("ORR %s, %s, %s", i.Xd(), i.Xn(), i.Xm())
	shift := i.Shift()
	if shift.HasShift() {
		s += fmt.Sprintf(", %s", shift)
	}
	return s
}
