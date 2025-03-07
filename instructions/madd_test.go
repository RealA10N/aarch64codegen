package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestMadd(t *testing.T) {
	// MADD X0, X1, X2, X3 (X0 = X3 + X1 * X2)
	AssertExpectedInstruction(t, "MADD X0, X1, X2, X3", instructions.MADD(
		registers.X0,
		registers.X1,
		registers.X2,
		registers.X3,
	))

	// MADD X20, X21, X22, X23
	AssertExpectedInstruction(t, "MADD X20, X21, X22, X23", instructions.MADD(
		registers.X20,
		registers.X21,
		registers.X22,
		registers.X23,
	))
}

func TestMaddWithXZR(t *testing.T) {
	// XZR as destination (result is discarded)
	AssertExpectedInstruction(t, "MADD XZR, X1, X2, X3", instructions.MADD(
		registers.XZR,
		registers.X1,
		registers.X2,
		registers.X3,
	))

	// XZR as first multiplicand
	AssertExpectedInstruction(t, "MADD X0, XZR, X2, X3", instructions.MADD(
		registers.X0,
		registers.XZR,
		registers.X2,
		registers.X3,
	))

	// XZR as second multiplicand
	AssertExpectedInstruction(t, "MADD X0, X1, XZR, X3", instructions.MADD(
		registers.X0,
		registers.X1,
		registers.XZR,
		registers.X3,
	))

	// MADD with zero register (X0 = X1 * X2, i.e., MUL)
	AssertExpectedInstruction(t, "MADD X0, X1, X2, XZR", instructions.MADD(
		registers.X0,
		registers.X1,
		registers.X2,
		registers.XZR,
	))

	// Multiple XZR registers
	AssertExpectedInstruction(t, "MADD XZR, XZR, XZR, XZR", instructions.MADD(
		registers.XZR,
		registers.XZR,
		registers.XZR,
		registers.XZR,
	))
}
