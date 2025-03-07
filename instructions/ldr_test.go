package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestLdr1(t *testing.T) {
	AssertExpectedInstruction(t, "LDR X0, [X1], #8", instructions.LDR(
		registers.X0,
		registers.GPorSPRegister(registers.X1),
		immediates.Immediate9(8),
	))
}

func TestLdrWithSP(t *testing.T) {
	// Test LDR with SP as the base register
	AssertExpectedInstruction(t, "LDR X5, [SP], #16", instructions.LDR(
		registers.X5,
		registers.SP, // Using SP directly
		immediates.Immediate9(16),
	))
}

func TestLdrWithXZR(t *testing.T) {
	// Test LDR with XZR as destination register
	AssertExpectedInstruction(t, "LDR XZR, [X1], #8", instructions.LDR(
		registers.XZR,
		registers.GPorSPRegister(registers.X1),
		immediates.Immediate9(8),
	))
}
