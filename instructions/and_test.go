package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestAnd(t *testing.T) {
	AssertExpectedInstruction(t, "and x9, x10, x11", instructions.AND(
		registers.GPRegisterX9,
		registers.GPRegisterX10,
		registers.GPRegisterX11,
		immediates.NoShift6(),
	))
}

func TestAndWithXZR(t *testing.T) {
	AssertExpectedInstruction(t, "and xzr, x10, x11", instructions.AND(
		registers.GPRegisterXZR,
		registers.GPRegisterX10,
		registers.GPRegisterX11,
		immediates.NoShift6(),
	))

	AssertExpectedInstruction(t, "and x9, xzr, x11", instructions.AND(
		registers.GPRegisterX9,
		registers.GPRegisterXZR,
		registers.GPRegisterX11,
		immediates.NoShift6(),
	))

	AssertExpectedInstruction(t, "and x9, x10, xzr", instructions.AND(
		registers.GPRegisterX9,
		registers.GPRegisterX10,
		registers.GPRegisterXZR,
		immediates.NoShift6(),
	))

	AssertExpectedInstruction(t, "and xzr, xzr, xzr", instructions.AND(
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
		immediates.NoShift6(),
	))
}

func TestAndWithShift(t *testing.T) {
	// Test LSL shift
	AssertExpectedInstruction(t, "and x9, x10, x11, lsl #2", instructions.AND(
		registers.GPRegisterX9,
		registers.GPRegisterX10,
		registers.GPRegisterX11,
		immediates.NewShift6(immediates.LSL, 2),
	))

	// Test LSR shift
	AssertExpectedInstruction(t, "and x9, x10, x11, lsr #4", instructions.AND(
		registers.GPRegisterX9,
		registers.GPRegisterX10,
		registers.GPRegisterX11,
		immediates.NewShift6(immediates.LSR, 4),
	))

	// Test ASR shift
	AssertExpectedInstruction(t, "and x9, x10, x11, asr #63", instructions.AND(
		registers.GPRegisterX9,
		registers.GPRegisterX10,
		registers.GPRegisterX11,
		immediates.NewShift6(immediates.ASR, 63),
	))

	// Test ROR shift
	AssertExpectedInstruction(t, "and x9, x10, x11, ror #8", instructions.AND(
		registers.GPRegisterX9,
		registers.GPRegisterX10,
		registers.GPRegisterX11,
		immediates.NewShift6(immediates.ROR, 8),
	))
}

func TestAndWithShiftAndXZR(t *testing.T) {
	// Test with XZR as destination and shifted operand
	AssertExpectedInstruction(t, "and xzr, x10, x11, lsl #2", instructions.AND(
		registers.GPRegisterXZR,
		registers.GPRegisterX10,
		registers.GPRegisterX11,
		immediates.NewShift6(immediates.LSL, 2),
	))

	// Test with XZR as first source and shifted operand
	AssertExpectedInstruction(t, "and x9, xzr, x11, lsr #4", instructions.AND(
		registers.GPRegisterX9,
		registers.GPRegisterXZR,
		registers.GPRegisterX11,
		immediates.NewShift6(immediates.LSR, 4),
	))

	// Test with XZR as second source and shift
	AssertExpectedInstruction(t, "and x9, x10, xzr, asr #1", instructions.AND(
		registers.GPRegisterX9,
		registers.GPRegisterX10,
		registers.GPRegisterXZR,
		immediates.NewShift6(immediates.ASR, 1),
	))
}
