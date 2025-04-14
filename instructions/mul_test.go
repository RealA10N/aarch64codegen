package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestMul(t *testing.T) {
	// MUL X6, X7, X8
	AssertExpectedInstruction(t, "MUL X6, X7, X8", instructions.MUL(
		registers.GPRegisterX6,
		registers.GPRegisterX7,
		registers.GPRegisterX8,
	))

	// MUL X24, X23, X22
	AssertExpectedInstruction(t, "MUL X24, X23, X22", instructions.MUL(
		registers.GPRegisterX24,
		registers.GPRegisterX23,
		registers.GPRegisterX22,
	))
}

func TestMulWithXZR(t *testing.T) {
	// Test MUL with XZR as destination
	AssertExpectedInstruction(t, "MUL XZR, X7, X8", instructions.MUL(
		registers.GPRegisterXZR,
		registers.GPRegisterX7,
		registers.GPRegisterX8,
	))

	// Test MUL with XZR as first source
	AssertExpectedInstruction(t, "MUL X6, XZR, X8", instructions.MUL(
		registers.GPRegisterX6,
		registers.GPRegisterXZR,
		registers.GPRegisterX8,
	))

	// Test MUL with XZR as second source
	AssertExpectedInstruction(t, "MUL X6, X7, XZR", instructions.MUL(
		registers.GPRegisterX6,
		registers.GPRegisterX7,
		registers.GPRegisterXZR,
	))
}
