package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestAnd(t *testing.T) {
	AssertExpectedInstruction(t, "AND X9, X10, X11", instructions.AND(
		registers.X9,
		registers.X10,
		registers.X11,
		immediates.NoShift6(),
	))
}

func TestAndWithXZR(t *testing.T) {
	AssertExpectedInstruction(t, "AND XZR, X10, X11", instructions.AND(
		registers.XZR,
		registers.X10,
		registers.X11,
		immediates.NoShift6(),
	))

	AssertExpectedInstruction(t, "AND X9, XZR, X11", instructions.AND(
		registers.X9,
		registers.XZR,
		registers.X11,
		immediates.NoShift6(),
	))

	AssertExpectedInstruction(t, "AND X9, X10, XZR", instructions.AND(
		registers.X9,
		registers.X10,
		registers.XZR,
		immediates.NoShift6(),
	))

	AssertExpectedInstruction(t, "AND XZR, XZR, XZR", instructions.AND(
		registers.XZR,
		registers.XZR,
		registers.XZR,
		immediates.NoShift6(),
	))
}

func TestAndWithShift(t *testing.T) {
	// Test LSL shift
	AssertExpectedInstruction(t, "AND X9, X10, X11, LSL #2", instructions.AND(
		registers.X9,
		registers.X10,
		registers.X11,
		immediates.NewShift6(immediates.LSL, 2),
	))

	// Test LSR shift
	AssertExpectedInstruction(t, "AND X9, X10, X11, LSR #4", instructions.AND(
		registers.X9,
		registers.X10,
		registers.X11,
		immediates.NewShift6(immediates.LSR, 4),
	))

	// Test ASR shift
	AssertExpectedInstruction(t, "AND X9, X10, X11, ASR #63", instructions.AND(
		registers.X9,
		registers.X10,
		registers.X11,
		immediates.NewShift6(immediates.ASR, 63),
	))

	// Test ROR shift
	AssertExpectedInstruction(t, "AND X9, X10, X11, ROR #8", instructions.AND(
		registers.X9,
		registers.X10,
		registers.X11,
		immediates.NewShift6(immediates.ROR, 8),
	))
}

func TestAndWithShiftAndXZR(t *testing.T) {
	// Test with XZR as destination and shifted operand
	AssertExpectedInstruction(t, "AND XZR, X10, X11, LSL #2", instructions.AND(
		registers.XZR,
		registers.X10,
		registers.X11,
		immediates.NewShift6(immediates.LSL, 2),
	))

	// Test with XZR as first source and shifted operand
	AssertExpectedInstruction(t, "AND X9, XZR, X11, LSR #4", instructions.AND(
		registers.X9,
		registers.XZR,
		registers.X11,
		immediates.NewShift6(immediates.LSR, 4),
	))

	// Test with XZR as second source and shift
	AssertExpectedInstruction(t, "AND X9, X10, XZR, ASR #1", instructions.AND(
		registers.X9,
		registers.X10,
		registers.XZR,
		immediates.NewShift6(immediates.ASR, 1),
	))
}
