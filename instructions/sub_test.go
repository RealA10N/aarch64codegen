package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestSub(t *testing.T) {
	// SUB X3, X4, X5
	AssertExpectedInstruction(t, "sub x3, x4, x5", instructions.SUB(
		registers.GPRegisterX3,
		registers.GPRegisterX4,
		registers.GPRegisterX5,
		immediates.DoNotSetFlags,
	))

	// SUB X27, X26, X25
	AssertExpectedInstruction(t, "sub x27, x26, x25", instructions.SUB(
		registers.GPRegisterX27,
		registers.GPRegisterX26,
		registers.GPRegisterX25,
		immediates.DoNotSetFlags,
	))
}

func TestSubWithXZR(t *testing.T) {
	// Test SUB with XZR as destination
	AssertExpectedInstruction(t, "sub xzr, x4, x5", instructions.SUB(
		registers.GPRegisterXZR,
		registers.GPRegisterX4,
		registers.GPRegisterX5,
		immediates.DoNotSetFlags,
	))

	// Test SUB with XZR as first source
	AssertExpectedInstruction(t, "sub x3, xzr, x5", instructions.SUB(
		registers.GPRegisterX3,
		registers.GPRegisterXZR,
		registers.GPRegisterX5,
		immediates.DoNotSetFlags,
	))

	// Test SUB with XZR as second source
	AssertExpectedInstruction(t, "sub x3, x4, xzr", instructions.SUB(
		registers.GPRegisterX3,
		registers.GPRegisterX4,
		registers.GPRegisterXZR,
		immediates.DoNotSetFlags,
	))

	// Multiple XZR registers
	AssertExpectedInstruction(t, "sub xzr, xzr, xzr", instructions.SUB(
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
		immediates.DoNotSetFlags,
	))
}

func TestSubImm(t *testing.T) {
	// SUB X3, X4, #10
	AssertExpectedInstruction(t, "sub x3, x4, #10", instructions.SUBI(
		registers.GPRegisterX3,
		registers.GPorSPRegisterX4,
		immediates.Immediate12(10),
	))

	// SUB X20, SP, #1000
	AssertExpectedInstruction(t, "sub x20, sp, #1000", instructions.SUBI(
		registers.GPRegisterX20,
		registers.GPorSPRegisterSP,
		immediates.Immediate12(1000),
	))
}

func TestSubDoSetFlags(t *testing.T) {
	AssertExpectedInstruction(t, "subs x0, x1, x2", instructions.SUB(
		registers.GPRegisterX0,
		registers.GPRegisterX1,
		registers.GPRegisterX2,
		immediates.DoSetFlags,
	))

	AssertExpectedInstruction(t, "subs xzr, x30, x29", instructions.SUB(
		registers.GPRegisterXZR,
		registers.GPRegisterX30,
		registers.GPRegisterX29,
		immediates.DoSetFlags,
	))
}
