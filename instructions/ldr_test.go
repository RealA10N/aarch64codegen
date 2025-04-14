package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestLdr1(t *testing.T) {
	AssertExpectedInstruction(t, "LDR X0, [X1], #8", instructions.LDR(
		registers.GPRegisterX0,
		registers.GPorSPRegisterX1,
		immediates.Immediate9(8),
	))
}

func TestLdrWithSP(t *testing.T) {
	// Test LDR with SP as the base register
	AssertExpectedInstruction(t, "LDR X5, [SP], #16", instructions.LDR(
		registers.GPRegisterX5,
		registers.GPorSPRegisterSP, // Using SP directly
		immediates.Immediate9(16),
	))
}

func TestLdrWithXZR(t *testing.T) {
	// Test LDR with XZR as destination register
	AssertExpectedInstruction(t, "LDR XZR, [X1], #8", instructions.LDR(
		registers.GPRegisterXZR,
		registers.GPorSPRegisterX1,
		immediates.Immediate9(8),
	))
}
