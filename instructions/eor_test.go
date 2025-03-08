package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestEor(t *testing.T) {
	// EOR X15, X16, X17
	AssertExpectedInstruction(t, "EOR X15, X16, X17", instructions.EOR(
		registers.X15,
		registers.X16,
		registers.X17,
	))

	// EOR X27, X28, X29
	AssertExpectedInstruction(t, "EOR X27, X28, X29", instructions.EOR(
		registers.X27,
		registers.X28,
		registers.X29,
	))
}

func TestEorWithXZR(t *testing.T) {
	// Test EOR with XZR as destination
	AssertExpectedInstruction(t, "EOR XZR, X16, X17", instructions.EOR(
		registers.XZR,
		registers.X16,
		registers.X17,
	))

	// Test EOR with XZR as first source
	AssertExpectedInstruction(t, "EOR X15, XZR, X17", instructions.EOR(
		registers.X15,
		registers.XZR,
		registers.X17,
	))

	// Test EOR with XZR as second source
	AssertExpectedInstruction(t, "EOR X15, X16, XZR", instructions.EOR(
		registers.X15,
		registers.X16,
		registers.XZR,
	))
}
