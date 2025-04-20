package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestLdr1(t *testing.T) {
	AssertExpectedInstruction(t, "ldr x0, [x1], #8", instructions.LDR(
		registers.GPRegisterX0,
		registers.GPorSPRegisterX1,
		immediates.Immediate9(8),
	))
}

func TestLdrWithSP(t *testing.T) {
	// Test LDR with SP as the base register
	AssertExpectedInstruction(t, "ldr x5, [sp], #16", instructions.LDR(
		registers.GPRegisterX5,
		registers.GPorSPRegisterSP, // Using SP directly
		immediates.Immediate9(16),
	))
}

func TestLdrWithXZR(t *testing.T) {
	// Test LDR with XZR as destination register
	AssertExpectedInstruction(t, "ldr xzr, [x1], #8", instructions.LDR(
		registers.GPRegisterXZR,
		registers.GPorSPRegisterX1,
		immediates.Immediate9(8),
	))
}
