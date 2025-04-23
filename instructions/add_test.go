package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestAdd(t *testing.T) {
	AssertExpectedInstruction(t, "add x0, x1, x2", instructions.NewAddShiftedRegister(
		registers.GPRegisterX0,
		registers.GPRegisterX1,
		registers.GPRegisterX2,
	))

	AssertExpectedInstruction(t, "add x30, x29, x28", instructions.NewAddShiftedRegister(
		registers.GPRegisterX30,
		registers.GPRegisterX29,
		registers.GPRegisterX28,
	))

}
func TestAddWithXZR(t *testing.T) {
	AssertExpectedInstruction(t, "add xzr, x1, x2", instructions.NewAddShiftedRegister(
		registers.GPRegisterXZR,
		registers.GPRegisterX1,
		registers.GPRegisterX2,
	))

	AssertExpectedInstruction(t, "add x0, xzr, x2", instructions.NewAddShiftedRegister(
		registers.GPRegisterX0,
		registers.GPRegisterXZR,
		registers.GPRegisterX2,
	))

	AssertExpectedInstruction(t, "add x0, x1, xzr", instructions.NewAddShiftedRegister(
		registers.GPRegisterX0,
		registers.GPRegisterX1,
		registers.GPRegisterXZR,
	))

	AssertExpectedInstruction(t, "add xzr, xzr, xzr", instructions.NewAddShiftedRegister(
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
		registers.GPRegisterXZR,
	))
}

func TestAddImmediate(t *testing.T) {
	AssertExpectedInstruction(t, "add x0, x1, #42", instructions.NewAddImmediate(
		registers.GPorSPRegisterX0,
		registers.GPorSPRegisterX1,
		immediates.Immediate12(42),
	))

	AssertExpectedInstruction(t, "add sp, sp, #4095", instructions.NewAddImmediate(
		registers.GPorSPRegisterSP,
		registers.GPorSPRegisterSP,
		immediates.Immediate12(4095),
	))
}

func TestAdds(t *testing.T) {
	AssertExpectedInstruction(t, "adds x0, x1, x2", instructions.NewAddsShiftedRegister(
		registers.GPRegisterX0,
		registers.GPRegisterX1,
		registers.GPRegisterX2,
	))

	AssertExpectedInstruction(t, "adds x30, x29, x28", instructions.NewAddsShiftedRegister(
		registers.GPRegisterX30,
		registers.GPRegisterX29,
		registers.GPRegisterX28,
	))
}

func TestAddsImmediate(t *testing.T) {
	AssertExpectedInstruction(t, "adds x0, x1, #42", instructions.NewAddsImmediate(
		registers.GPRegisterX0,
		registers.GPorSPRegisterX1,
		immediates.Immediate12(42),
	))

	AssertExpectedInstruction(t, "adds xzr, sp, #4095", instructions.NewAddsImmediate(
		registers.GPRegisterXZR,
		registers.GPorSPRegisterSP,
		immediates.Immediate12(4095),
	))
}
