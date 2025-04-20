package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestOrr(t *testing.T) {
	AssertExpectedInstruction(t, "orr x12, x13, x14", instructions.ORR(
		registers.GPRegisterX12,
		registers.GPRegisterX13,
		registers.GPRegisterX14,
		immediates.NoShift6(),
	))

	AssertExpectedInstruction(t, "orr x18, x19, x20", instructions.ORR(
		registers.GPRegisterX18,
		registers.GPRegisterX19,
		registers.GPRegisterX20,
		immediates.NoShift6(),
	))
}

func TestOrrWithXZR(t *testing.T) {
	// Test ORR with XZR as destination
	AssertExpectedInstruction(t, "orr xzr, x13, x14", instructions.ORR(
		registers.GPRegisterXZR,
		registers.GPRegisterX13,
		registers.GPRegisterX14,
		immediates.NoShift6(),
	))

	// Test ORR with XZR as first source
	AssertExpectedInstruction(t, "orr x12, xzr, x14", instructions.ORR(
		registers.GPRegisterX12,
		registers.GPRegisterXZR,
		registers.GPRegisterX14,
		immediates.NoShift6(),
	))

	// Test ORR with XZR as second source
	AssertExpectedInstruction(t, "orr x12, x13, xzr", instructions.ORR(
		registers.GPRegisterX12,
		registers.GPRegisterX13,
		registers.GPRegisterXZR,
		immediates.NoShift6(),
	))
}

func TestOrrWithShift(t *testing.T) {
	// Test LSL shift
	AssertExpectedInstruction(t, "orr x12, x13, x14, lsl #1", instructions.ORR(
		registers.GPRegisterX12,
		registers.GPRegisterX13,
		registers.GPRegisterX14,
		immediates.NewShift6(immediates.LSL, 1),
	))

	// Test LSR shift
	AssertExpectedInstruction(t, "orr x12, x13, x14, lsr #32", instructions.ORR(
		registers.GPRegisterX12,
		registers.GPRegisterX13,
		registers.GPRegisterX14,
		immediates.NewShift6(immediates.LSR, 32),
	))

	// Test ASR shift
	AssertExpectedInstruction(t, "orr x12, x13, x14, asr #48", instructions.ORR(
		registers.GPRegisterX12,
		registers.GPRegisterX13,
		registers.GPRegisterX14,
		immediates.NewShift6(immediates.ASR, 48),
	))

	// Test ROR shift
	AssertExpectedInstruction(t, "orr x12, x13, x14, ror #16", instructions.ORR(
		registers.GPRegisterX12,
		registers.GPRegisterX13,
		registers.GPRegisterX14,
		immediates.NewShift6(immediates.ROR, 16),
	))
}

func TestOrrWithShiftAndXZR(t *testing.T) {
	// Test with XZR as destination and shifted operand
	AssertExpectedInstruction(t, "orr xzr, x13, x14, lsl #2", instructions.ORR(
		registers.GPRegisterXZR,
		registers.GPRegisterX13,
		registers.GPRegisterX14,
		immediates.NewShift6(immediates.LSL, 2),
	))

	// Test with XZR as first source and shifted operand
	AssertExpectedInstruction(t, "orr x12, xzr, x14, lsr #3", instructions.ORR(
		registers.GPRegisterX12,
		registers.GPRegisterXZR,
		registers.GPRegisterX14,
		immediates.NewShift6(immediates.LSR, 3),
	))

	// Test with XZR as second source and shift
	AssertExpectedInstruction(t, "orr x12, x13, xzr, asr #4", instructions.ORR(
		registers.GPRegisterX12,
		registers.GPRegisterX13,
		registers.GPRegisterXZR,
		immediates.NewShift6(immediates.ASR, 4),
	))
}
