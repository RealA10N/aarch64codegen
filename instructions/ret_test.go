package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestRet1(t *testing.T) {
	AssertExpectedInstruction(t, "ret", instructions.RET(registers.GPRegisterX30))
}

func TestRet2(t *testing.T) {
	AssertExpectedInstruction(t, "ret x0", instructions.RET(registers.GPRegisterX0))
}

func TestRet3(t *testing.T) {
	AssertExpectedInstruction(t, "ret x29", instructions.RET(registers.GPRegisterX29))
}

func TestRetWithXZR(t *testing.T) {
	AssertExpectedInstruction(t, "ret xzr", instructions.RET(registers.GPRegisterXZR))
}
