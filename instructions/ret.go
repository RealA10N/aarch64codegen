package instructions

import (
	"alon.kr/x/aarch64codegen/registers"
)

type Ret uint32

func RET(Xn registers.GPRegister) Ret {
	return Ret(0xD65F0000 | (Xn.Binary() << 5))
}

func (i Ret) Xn() registers.GPRegister {
	return registers.GPRegister((i >> 5) & 0b11111)
}

func (r Ret) String() string {
	Xn := r.Xn()
	s := "RET"
	if Xn != registers.GPRegisterX30 {
		s += " " + Xn.String()
	}
	return s
}

func (r Ret) Binary() uint32 {
	return uint32(r)
}
