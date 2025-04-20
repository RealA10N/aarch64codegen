package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestMadd(t *testing.T) {
	// MADD X0, X1, X2, X3 (X0 = X3 + X1 * X2)
	AssertExpectedInstruction(t, "madd x0, x1, x2, x3", instructions.MADD(
		registers.GPRegisterX0,
		registers.GPRegisterX1,
		registers.GPRegisterX2,
		registers.GPRegisterX3,
	))

	// MADD X20, X21, X22, X23
	AssertExpectedInstruction(t, "madd x20, x21, x22, x23", instructions.MADD(
		registers.GPRegisterX20,
		registers.GPRegisterX21,
		registers.GPRegisterX22,
		registers.GPRegisterX23,
	))
}

func TestMaddWithXZR(t *testing.T) {
	// XZR as destination (result is discarded)
	AssertExpectedInstruction(t, "madd xzr, x1, x2, x3", instructions.MADD(
		registers.GPRegisterXZR,
		registers.GPRegisterX1,
		registers.GPRegisterX2,
		registers.GPRegisterX3,
	))

	// XZR as first multiplicand
	AssertExpectedInstruction(t, "madd x0, xzr, x2, x3", instructions.MADD(
		registers.GPRegisterX0,
		registers.GPRegisterXZR,
		registers.GPRegisterX2,
		registers.GPRegisterX3,
	))

	// XZR as second multiplicand
	AssertExpectedInstruction(t, "madd x0, x1, xzr, x3", instructions.MADD(
		registers.GPRegisterX0,
		registers.GPRegisterX1,
		registers.GPRegisterXZR,
		registers.GPRegisterX3,
	))

	// MADD with zero register (X0 = X1 * X2, i.e., MUL)
	AssertExpectedInstruction(t, "madd x0, x1, x2, xzr", instructions.MADD(
		registers.GPRegisterX0,
		registers.GPRegisterX1,
		registers.GPRegisterX2,
		registers.GPRegisterXZR,
	))

	// Multiple XZR registers
	AssertExpectedInstruction(t, "madd xzr, xzr, xzr, xzr", instructions.MADD(
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
	))
}
