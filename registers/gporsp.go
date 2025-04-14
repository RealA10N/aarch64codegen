package registers

import "fmt"

type GPorSPRegister uint8

const (
	GPorSPRegisterX0 GPorSPRegister = iota
	GPorSPRegisterX1
	GPorSPRegisterX2
	GPorSPRegisterX3
	GPorSPRegisterX4
	GPorSPRegisterX5
	GPorSPRegisterX6
	GPorSPRegisterX7
	GPorSPRegisterX8
	GPorSPRegisterX9
	GPorSPRegisterX10
	GPorSPRegisterX11
	GPorSPRegisterX12
	GPorSPRegisterX13
	GPorSPRegisterX14
	GPorSPRegisterX15
	GPorSPRegisterX16
	GPorSPRegisterX17
	GPorSPRegisterX18
	GPorSPRegisterX19
	GPorSPRegisterX20
	GPorSPRegisterX21
	GPorSPRegisterX22
	GPorSPRegisterX23
	GPorSPRegisterX24
	GPorSPRegisterX25
	GPorSPRegisterX26
	GPorSPRegisterX27
	GPorSPRegisterX28
	GPorSPRegisterX29
	GPorSPRegisterX30
	GPorSPRegisterSP
)

func (r GPorSPRegister) Validate() error {
	if r > GPorSPRegisterSP {
		return fmt.Errorf("invalid general purpose or SP register: %d", r)
	}
	return nil
}

func (r GPorSPRegister) Binary() uint32 {
	return uint32(r)
}

func (r GPorSPRegister) String() string {
	if r == GPorSPRegisterSP {
		return "SP"
	}
	return fmt.Sprintf("X%d", r)
}
