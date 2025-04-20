package registers

import "fmt"

type GPRegister uint8

const (
	GPRegisterX0 GPRegister = iota
	GPRegisterX1
	GPRegisterX2
	GPRegisterX3
	GPRegisterX4
	GPRegisterX5
	GPRegisterX6
	GPRegisterX7
	GPRegisterX8
	GPRegisterX9
	GPRegisterX10
	GPRegisterX11
	GPRegisterX12
	GPRegisterX13
	GPRegisterX14
	GPRegisterX15
	GPRegisterX16
	GPRegisterX17
	GPRegisterX18
	GPRegisterX19
	GPRegisterX20
	GPRegisterX21
	GPRegisterX22
	GPRegisterX23
	GPRegisterX24
	GPRegisterX25
	GPRegisterX26
	GPRegisterX27
	GPRegisterX28
	GPRegisterX29
	GPRegisterX30
	GPRegisterXZR // Zero register (X31)
)

func (r GPRegister) Validate() error {
	if r > GPRegisterXZR {
		return fmt.Errorf("invalid general purpose register: %d", r)
	}
	return nil
}

func (r GPRegister) Binary() uint32 {
	return uint32(r)
}

func (r GPRegister) String() string {
	if r == GPRegisterXZR {
		return "xzr"
	}
	return fmt.Sprintf("x%d", r)
}
