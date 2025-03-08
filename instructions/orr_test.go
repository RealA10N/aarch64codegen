package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestOrr(t *testing.T) {
	AssertExpectedInstruction(t, "ORR X12, X13, X14", instructions.ORR(
		registers.X12,
		registers.X13,
		registers.X14,
		immediates.NoShift6(),
	))

	AssertExpectedInstruction(t, "ORR X18, X19, X20", instructions.ORR(
		registers.X18,
		registers.X19,
		registers.X20,
		immediates.NoShift6(),
	))
}

func TestOrrWithXZR(t *testing.T) {
	// Test ORR with XZR as destination
	AssertExpectedInstruction(t, "ORR XZR, X13, X14", instructions.ORR(
		registers.XZR,
		registers.X13,
		registers.X14,
		immediates.NoShift6(),
	))

	// Test ORR with XZR as first source
	AssertExpectedInstruction(t, "ORR X12, XZR, X14", instructions.ORR(
		registers.X12,
		registers.XZR,
		registers.X14,
		immediates.NoShift6(),
	))

	// Test ORR with XZR as second source
	AssertExpectedInstruction(t, "ORR X12, X13, XZR", instructions.ORR(
		registers.X12,
		registers.X13,
		registers.XZR,
		immediates.NoShift6(),
	))
}

func TestOrrWithShift(t *testing.T) {
	// Test LSL shift
	AssertExpectedInstruction(t, "ORR X12, X13, X14, LSL #1", instructions.ORR(
		registers.X12,
		registers.X13,
		registers.X14,
		immediates.NewShift6(immediates.LSL, 1),
	))

	// Test LSR shift
	AssertExpectedInstruction(t, "ORR X12, X13, X14, LSR #32", instructions.ORR(
		registers.X12,
		registers.X13,
		registers.X14,
		immediates.NewShift6(immediates.LSR, 32),
	))

	// Test ASR shift
	AssertExpectedInstruction(t, "ORR X12, X13, X14, ASR #48", instructions.ORR(
		registers.X12,
		registers.X13,
		registers.X14,
		immediates.NewShift6(immediates.ASR, 48),
	))

	// Test ROR shift
	AssertExpectedInstruction(t, "ORR X12, X13, X14, ROR #16", instructions.ORR(
		registers.X12,
		registers.X13,
		registers.X14,
		immediates.NewShift6(immediates.ROR, 16),
	))
}

func TestOrrWithShiftAndXZR(t *testing.T) {
	// Test with XZR as destination and shifted operand
	AssertExpectedInstruction(t, "ORR XZR, X13, X14, LSL #2", instructions.ORR(
		registers.XZR,
		registers.X13,
		registers.X14,
		immediates.NewShift6(immediates.LSL, 2),
	))

	// Test with XZR as first source and shifted operand
	AssertExpectedInstruction(t, "ORR X12, XZR, X14, LSR #3", instructions.ORR(
		registers.X12,
		registers.XZR,
		registers.X14,
		immediates.NewShift6(immediates.LSR, 3),
	))

	// Test with XZR as second source and shift
	AssertExpectedInstruction(t, "ORR X12, X13, XZR, ASR #4", instructions.ORR(
		registers.X12,
		registers.X13,
		registers.XZR,
		immediates.NewShift6(immediates.ASR, 4),
	))
}
