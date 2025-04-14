package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestMadd(t *testing.T) {
	// MADD X0, X1, X2, X3 (X0 = X3 + X1 * X2)
	AssertExpectedInstruction(t, "MADD X0, X1, X2, X3", instructions.MADD(
		registers.GPRegisterX0,
		registers.GPRegisterX1,
		registers.GPRegisterX2,
		registers.GPRegisterX3,
	))

	// MADD X20, X21, X22, X23
	AssertExpectedInstruction(t, "MADD X20, X21, X22, X23", instructions.MADD(
		registers.GPRegisterX20,
		registers.GPRegisterX21,
		registers.GPRegisterX22,
		registers.GPRegisterX23,
	))
}

func TestMaddWithXZR(t *testing.T) {
	// XZR as destination (result is discarded)
	AssertExpectedInstruction(t, "MADD XZR, X1, X2, X3", instructions.MADD(
		registers.GPRegisterXZR,
		registers.GPRegisterX1,
		registers.GPRegisterX2,
		registers.GPRegisterX3,
	))

	// XZR as first multiplicand
	AssertExpectedInstruction(t, "MADD X0, XZR, X2, X3", instructions.MADD(
		registers.GPRegisterX0,
		registers.GPRegisterXZR,
		registers.GPRegisterX2,
		registers.GPRegisterX3,
	))

	// XZR as second multiplicand
	AssertExpectedInstruction(t, "MADD X0, X1, XZR, X3", instructions.MADD(
		registers.GPRegisterX0,
		registers.GPRegisterX1,
		registers.GPRegisterXZR,
		registers.GPRegisterX3,
	))

	// MADD with zero register (X0 = X1 * X2, i.e., MUL)
	AssertExpectedInstruction(t, "MADD X0, X1, X2, XZR", instructions.MADD(
		registers.GPRegisterX0,
		registers.GPRegisterX1,
		registers.GPRegisterX2,
		registers.GPRegisterXZR,
	))

	// Multiple XZR registers
	AssertExpectedInstruction(t, "MADD XZR, XZR, XZR, XZR", instructions.MADD(
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
	))
}
