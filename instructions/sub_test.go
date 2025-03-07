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
		registers.X3,
		registers.X4,
		registers.X5,
	))

	// SUB X27, X26, X25
	AssertExpectedInstruction(t, "SUB X27, X26, X25", instructions.SUB(
		registers.X27,
		registers.X26,
		registers.X25,
	))
}

func TestSubImm(t *testing.T) {
	// SUB X3, X4, #10
	AssertExpectedInstruction(t, "SUB X3, X4, #10", instructions.SUBI(
		registers.X3,
		registers.GPorSPRegister(registers.X4),
		immediates.Immediate12(10),
	))

	// SUB X20, SP, #1000
	AssertExpectedInstruction(t, "SUB X20, SP, #1000", instructions.SUBI(
		registers.X20,
		registers.SP,
		immediates.Immediate12(1000),
	))
}

func TestSubWithXZR(t *testing.T) {
	// Test SUB with XZR as destination
	AssertExpectedInstruction(t, "SUB XZR, X4, X5", instructions.SUB(
		registers.XZR,
		registers.X4,
		registers.X5,
	))

	// Test SUB with XZR as first source
	AssertExpectedInstruction(t, "SUB X3, XZR, X5", instructions.SUB(
		registers.X3,
		registers.XZR,
		registers.X5,
	))

	// Test SUB with XZR as second source
	AssertExpectedInstruction(t, "SUB X3, X4, XZR", instructions.SUB(
		registers.X3,
		registers.X4,
		registers.XZR,
	))
}
