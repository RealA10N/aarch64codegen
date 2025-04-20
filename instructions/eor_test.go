package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestEor(t *testing.T) {
	AssertExpectedInstruction(t, "eor x15, x16, x17", instructions.EOR(
		registers.GPRegisterX15,
		registers.GPRegisterX16,
		registers.GPRegisterX17,
		immediates.NoShift6(),
	))

	AssertExpectedInstruction(t, "eor x27, x28, x29", instructions.EOR(
		registers.GPRegisterX27,
		registers.GPRegisterX28,
		registers.GPRegisterX29,
		immediates.NoShift6(),
	))
}

func TestEorWithXZR(t *testing.T) {
	// Test EOR with XZR as destination
	AssertExpectedInstruction(t, "eor xzr, x16, x17", instructions.EOR(
		registers.GPRegisterXZR,
		registers.GPRegisterX16,
		registers.GPRegisterX17,
		immediates.NoShift6(),
	))

	// Test EOR with XZR as first source
	AssertExpectedInstruction(t, "eor x15, xzr, x17", instructions.EOR(
		registers.GPRegisterX15,
		registers.GPRegisterXZR,
		registers.GPRegisterX17,
		immediates.NoShift6(),
	))

	// Test EOR with XZR as second source
	AssertExpectedInstruction(t, "eor x15, x16, xzr", instructions.EOR(
		registers.GPRegisterX15,
		registers.GPRegisterX16,
		registers.GPRegisterXZR,
		immediates.NoShift6(),
	))
}

func TestEorWithShift(t *testing.T) {
	// Test LSL shift
	AssertExpectedInstruction(t, "eor x15, x16, x17, lsl #5", instructions.EOR(
		registers.GPRegisterX15,
		registers.GPRegisterX16,
		registers.GPRegisterX17,
		immediates.NewShift6(immediates.LSL, 5),
	))

	// Test LSR shift
	AssertExpectedInstruction(t, "eor x15, x16, x17, lsr #24", instructions.EOR(
		registers.GPRegisterX15,
		registers.GPRegisterX16,
		registers.GPRegisterX17,
		immediates.NewShift6(immediates.LSR, 24),
	))

	// Test ASR shift
	AssertExpectedInstruction(t, "eor x15, x16, x17, asr #31", instructions.EOR(
		registers.GPRegisterX15,
		registers.GPRegisterX16,
		registers.GPRegisterX17,
		immediates.NewShift6(immediates.ASR, 31),
	))

	// Test ROR shift
	AssertExpectedInstruction(t, "eor x15, x16, x17, ror #42", instructions.EOR(
		registers.GPRegisterX15,
		registers.GPRegisterX16,
		registers.GPRegisterX17,
		immediates.NewShift6(immediates.ROR, 42),
	))
}

func TestEorWithShiftAndXZR(t *testing.T) {
	// Test with XZR as destination and shifted operand
	AssertExpectedInstruction(t, "eor xzr, x16, x17, lsl #3", instructions.EOR(
		registers.GPRegisterXZR,
		registers.GPRegisterX16,
		registers.GPRegisterX17,
		immediates.NewShift6(immediates.LSL, 3),
	))

	// Test with XZR as first source and shifted operand
	AssertExpectedInstruction(t, "eor x15, xzr, x17, lsr #7", instructions.EOR(
		registers.GPRegisterX15,
		registers.GPRegisterXZR,
		registers.GPRegisterX17,
		immediates.NewShift6(immediates.LSR, 7),
	))

	// Test with XZR as second source and shift
	AssertExpectedInstruction(t, "eor x15, x16, xzr, asr #9", instructions.EOR(
		registers.GPRegisterX15,
		registers.GPRegisterX16,
		registers.GPRegisterXZR,
		immediates.NewShift6(immediates.ASR, 9),
	))
}
