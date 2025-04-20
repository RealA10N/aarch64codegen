package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestMul(t *testing.T) {
	AssertExpectedInstruction(t, "mul x6, x7, x8", instructions.MUL(
		registers.GPRegisterX6,
		registers.GPRegisterX7,
		registers.GPRegisterX8,
	))

	AssertExpectedInstruction(t, "mul x24, x23, x22", instructions.MUL(
		registers.GPRegisterX24,
		registers.GPRegisterX23,
		registers.GPRegisterX22,
	))
}

func TestMulWithXZR(t *testing.T) {
	// Test MUL with XZR as destination
	AssertExpectedInstruction(t, "mul xzr, x7, x8", instructions.MUL(
		registers.GPRegisterXZR,
		registers.GPRegisterX7,
		registers.GPRegisterX8,
	))

	// Test MUL with XZR as first source
	AssertExpectedInstruction(t, "mul x6, xzr, x8", instructions.MUL(
		registers.GPRegisterX6,
		registers.GPRegisterXZR,
		registers.GPRegisterX8,
	))

	// Test MUL with XZR as second source
	AssertExpectedInstruction(t, "mul x6, x7, xzr", instructions.MUL(
		registers.GPRegisterX6,
		registers.GPRegisterX7,
		registers.GPRegisterXZR,
	))
}
