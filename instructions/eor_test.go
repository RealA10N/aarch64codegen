package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestEor(t *testing.T) {
	AssertExpectedInstruction(t, "EOR X15, X16, X17", instructions.EOR(
		registers.X15,
		registers.X16,
		registers.X17,
		immediates.NoShift6(),
	))

	AssertExpectedInstruction(t, "EOR X27, X28, X29", instructions.EOR(
		registers.X27,
		registers.X28,
		registers.X29,
		immediates.NoShift6(),
	))
}

func TestEorWithXZR(t *testing.T) {
	// Test EOR with XZR as destination
	AssertExpectedInstruction(t, "EOR XZR, X16, X17", instructions.EOR(
		registers.XZR,
		registers.X16,
		registers.X17,
		immediates.NoShift6(),
	))

	// Test EOR with XZR as first source
	AssertExpectedInstruction(t, "EOR X15, XZR, X17", instructions.EOR(
		registers.X15,
		registers.XZR,
		registers.X17,
		immediates.NoShift6(),
	))

	// Test EOR with XZR as second source
	AssertExpectedInstruction(t, "EOR X15, X16, XZR", instructions.EOR(
		registers.X15,
		registers.X16,
		registers.XZR,
		immediates.NoShift6(),
	))
}

func TestEorWithShift(t *testing.T) {
	// Test LSL shift
	AssertExpectedInstruction(t, "EOR X15, X16, X17, LSL #5", instructions.EOR(
		registers.X15,
		registers.X16,
		registers.X17,
		immediates.NewShift6(immediates.LSL, 5),
	))

	// Test LSR shift
	AssertExpectedInstruction(t, "EOR X15, X16, X17, LSR #24", instructions.EOR(
		registers.X15,
		registers.X16,
		registers.X17,
		immediates.NewShift6(immediates.LSR, 24),
	))

	// Test ASR shift
	AssertExpectedInstruction(t, "EOR X15, X16, X17, ASR #31", instructions.EOR(
		registers.X15,
		registers.X16,
		registers.X17,
		immediates.NewShift6(immediates.ASR, 31),
	))

	// Test ROR shift
	AssertExpectedInstruction(t, "EOR X15, X16, X17, ROR #42", instructions.EOR(
		registers.X15,
		registers.X16,
		registers.X17,
		immediates.NewShift6(immediates.ROR, 42),
	))
}

func TestEorWithShiftAndXZR(t *testing.T) {
	// Test with XZR as destination and shifted operand
	AssertExpectedInstruction(t, "EOR XZR, X16, X17, LSL #3", instructions.EOR(
		registers.XZR,
		registers.X16,
		registers.X17,
		immediates.NewShift6(immediates.LSL, 3),
	))

	// Test with XZR as first source and shifted operand
	AssertExpectedInstruction(t, "EOR X15, XZR, X17, LSR #7", instructions.EOR(
		registers.X15,
		registers.XZR,
		registers.X17,
		immediates.NewShift6(immediates.LSR, 7),
	))

	// Test with XZR as second source and shift
	AssertExpectedInstruction(t, "EOR X15, X16, XZR, ASR #9", instructions.EOR(
		registers.X15,
		registers.X16,
		registers.XZR,
		immediates.NewShift6(immediates.ASR, 9),
	))
}
