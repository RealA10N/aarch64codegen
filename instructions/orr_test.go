package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestOrr(t *testing.T) {
	// ORR X12, X13, X14
	AssertExpectedInstruction(t, "ORR X12, X13, X14", instructions.ORR(
		registers.X12,
		registers.X13,
		registers.X14,
	))

	// ORR X18, X19, X20
	AssertExpectedInstruction(t, "ORR X18, X19, X20", instructions.ORR(
		registers.X18,
		registers.X19,
		registers.X20,
	))
}

func TestOrrWithXZR(t *testing.T) {
	// Test ORR with XZR as destination
	AssertExpectedInstruction(t, "ORR XZR, X13, X14", instructions.ORR(
		registers.XZR,
		registers.X13,
		registers.X14,
	))

	// Test ORR with XZR as first source
	AssertExpectedInstruction(t, "ORR X12, XZR, X14", instructions.ORR(
		registers.X12,
		registers.XZR,
		registers.X14,
	))

	// Test ORR with XZR as second source
	AssertExpectedInstruction(t, "ORR X12, X13, XZR", instructions.ORR(
		registers.X12,
		registers.X13,
		registers.XZR,
	))
}
