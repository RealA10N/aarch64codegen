package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestAdd(t *testing.T) {
	AssertExpectedInstruction(t, "ADD X0, X1, X2", instructions.ADD(
		registers.GPRegisterX0,
		registers.GPRegisterX1,
		registers.GPRegisterX2,
		immediates.DoNotSetFlags,
	))

	AssertExpectedInstruction(t, "ADD X30, X29, X28", instructions.ADD(
		registers.GPRegisterX30,
		registers.GPRegisterX29,
		registers.GPRegisterX28,
		immediates.DoNotSetFlags,
	))

}
func TestAddWithXZR(t *testing.T) {
	AssertExpectedInstruction(t, "ADD XZR, X1, X2", instructions.ADD(
		registers.GPRegisterXZR,
		registers.GPRegisterX1,
		registers.GPRegisterX2,
		immediates.DoNotSetFlags,
	))

	AssertExpectedInstruction(t, "ADD X0, XZR, X2", instructions.ADD(
		registers.GPRegisterX0,
		registers.GPRegisterXZR,
		registers.GPRegisterX2,
		immediates.DoNotSetFlags,
	))

	AssertExpectedInstruction(t, "ADD X0, X1, XZR", instructions.ADD(
		registers.GPRegisterX0,
		registers.GPRegisterX1,
		registers.GPRegisterXZR,
		immediates.DoNotSetFlags,
	))

	AssertExpectedInstruction(t, "ADD XZR, XZR, XZR", instructions.ADD(
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
		immediates.DoNotSetFlags,
	))
}

func TestAddImm(t *testing.T) {
	AssertExpectedInstruction(t, "ADD X0, X1, #42", instructions.ADDI(
		registers.GPorSPRegisterX0,
		registers.GPorSPRegisterX1,
		immediates.Immediate12(42),
		immediates.DoNotSetFlags,
	))

	AssertExpectedInstruction(t, "ADD SP, SP, #4095", instructions.ADDI(
		registers.GPorSPRegisterSP,
		registers.GPorSPRegisterSP,
		immediates.Immediate12(4095),
		immediates.DoNotSetFlags,
	))
}

func TestAddDoSetFlags(t *testing.T) {
	AssertExpectedInstruction(t, "ADDS X0, X1, X2", instructions.ADD(
		registers.GPRegisterX0,
		registers.GPRegisterX1,
		registers.GPRegisterX2,
		immediates.DoSetFlags,
	))

	AssertExpectedInstruction(t, "ADDS X30, X29, X28", instructions.ADD(
		registers.GPRegisterX30,
		registers.GPRegisterX29,
		registers.GPRegisterX28,
		immediates.DoSetFlags,
	))
}
