package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestAdd(t *testing.T) {
	AssertExpectedInstruction(t, "ADD X0, X1, X2", instructions.ADD(
		registers.X0,
		registers.X1,
		registers.X2,
	))

	AssertExpectedInstruction(t, "ADD X30, X29, X28", instructions.ADD(
		registers.X30,
		registers.X29,
		registers.X28,
	))

	AssertExpectedInstruction(t, "ADD XZR, X1, X2", instructions.ADD(
		registers.XZR,
		registers.X1,
		registers.X2,
	))

	AssertExpectedInstruction(t, "ADD X0, XZR, X2", instructions.ADD(
		registers.X0,
		registers.XZR,
		registers.X2,
	))

	AssertExpectedInstruction(t, "ADD X0, X1, XZR", instructions.ADD(
		registers.X0,
		registers.X1,
		registers.XZR,
	))
}

func TestAddImm(t *testing.T) {
	AssertExpectedInstruction(t, "ADD X0, X1, #42", instructions.ADDI(
		registers.GPorSPRegister(registers.X0),
		registers.GPorSPRegister(registers.X1),
		immediates.Immediate12(42),
	))

	AssertExpectedInstruction(t, "ADD SP, SP, #4095", instructions.ADDI(
		registers.SP,
		registers.SP,
		immediates.Immediate12(4095),
	))
}
