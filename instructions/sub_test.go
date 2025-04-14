package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestSub(t *testing.T) {
	// SUB X3, X4, X5
	AssertExpectedInstruction(t, "SUB X3, X4, X5", instructions.SUB(
		registers.GPRegisterX3,
		registers.GPRegisterX4,
		registers.GPRegisterX5,
	))

	// SUB X27, X26, X25
	AssertExpectedInstruction(t, "SUB X27, X26, X25", instructions.SUB(
		registers.GPRegisterX27,
		registers.GPRegisterX26,
		registers.GPRegisterX25,
	))
}

func TestSubWithXZR(t *testing.T) {
	// Test SUB with XZR as destination
	AssertExpectedInstruction(t, "SUB XZR, X4, X5", instructions.SUB(
		registers.GPRegisterXZR,
		registers.GPRegisterX4,
		registers.GPRegisterX5,
	))

	// Test SUB with XZR as first source
	AssertExpectedInstruction(t, "SUB X3, XZR, X5", instructions.SUB(
		registers.GPRegisterX3,
		registers.GPRegisterXZR,
		registers.GPRegisterX5,
	))

	// Test SUB with XZR as second source
	AssertExpectedInstruction(t, "SUB X3, X4, XZR", instructions.SUB(
		registers.GPRegisterX3,
		registers.GPRegisterX4,
		registers.GPRegisterXZR,
	))

	// Multiple XZR registers
	AssertExpectedInstruction(t, "SUB XZR, XZR, XZR", instructions.SUB(
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
	))
}

func TestSubImm(t *testing.T) {
	// SUB X3, X4, #10
	AssertExpectedInstruction(t, "SUB X3, X4, #10", instructions.SUBI(
		registers.GPRegisterX3,
		registers.GPorSPRegisterX4,
		immediates.Immediate12(10),
	))

	// SUB X20, SP, #1000
	AssertExpectedInstruction(t, "SUB X20, SP, #1000", instructions.SUBI(
		registers.GPRegisterX20,
		registers.GPorSPRegisterSP,
		immediates.Immediate12(1000),
	))
}
