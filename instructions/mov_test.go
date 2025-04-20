package instructions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestMovShiftValidateValid(t *testing.T) {
	// Valid shift values should not return an error
	validShifts := []instructions.MovShift{
		instructions.MovShift0,
		instructions.MovShift16,
		instructions.MovShift32,
		instructions.MovShift48,
	}

	for _, shift := range validShifts {
		err := shift.Validate()
		assert.NoError(t, err)
	}
}

func TestMovShiftValidateInvalid(t *testing.T) {
	// Invalid shift value (greater than 3) should return an error
	invalidShift := instructions.MovShift(4)
	err := invalidShift.Validate()
	assert.Error(t, err)
}

func TestMovz1(t *testing.T) {
	AssertExpectedInstruction(t, "movz x0, #1234, lsl #32", instructions.MOVZ(
		registers.GPRegisterX0,
		immediates.Immediate16(1234),
		instructions.MovShift32,
	))
}

func TestMovz2(t *testing.T) {
	AssertExpectedInstruction(t, "movz x30, #65535", instructions.MOVZ(
		registers.GPRegisterX30,
		immediates.Immediate16(0xFFFF),
		instructions.MovShift0,
	))
}

func TestMovzXZR(t *testing.T) {
	// Test MOVZ with XZR as destination
	AssertExpectedInstruction(t, "movz xzr, #1234", instructions.MOVZ(
		registers.GPRegisterXZR,
		immediates.Immediate16(1234),
		instructions.MovShift0,
	))
}

func TestMovk1(t *testing.T) {
	AssertExpectedInstruction(t, "movk x10, #0, lsl #48", instructions.MOVK(
		registers.GPRegisterX10,
		immediates.Immediate16(0),
		instructions.MovShift48,
	))
}

func TestMovk2(t *testing.T) {
	AssertExpectedInstruction(t, "movk x29, #65534, lsl #16", instructions.MOVK(
		registers.GPRegisterX29,
		immediates.Immediate16(0xFFFE),
		instructions.MovShift16,
	))
}

func TestMovkXZR(t *testing.T) {
	// Test MOVK with XZR as destination
	AssertExpectedInstruction(t, "movk xzr, #4321", instructions.MOVK(
		registers.GPRegisterXZR,
		immediates.Immediate16(4321),
		instructions.MovShift0,
	))
}

func TestMovn1(t *testing.T) {
	AssertExpectedInstruction(t, "movn x2, #48879, lsl #32", instructions.MOVN(
		registers.GPRegisterX2,
		immediates.Immediate16(0xBEEF),
		instructions.MovShift32,
	))
}

func TestMovn2(t *testing.T) {
	AssertExpectedInstruction(t, "movn x21, #0", instructions.MOVN(
		registers.GPRegisterX21,
		immediates.Immediate16(0),
		instructions.MovShift0,
	))
}

func TestMovnXZR(t *testing.T) {
	// Test MOVN with XZR as destination
	AssertExpectedInstruction(t, "movn xzr, #5678", instructions.MOVN(
		registers.GPRegisterXZR,
		immediates.Immediate16(5678),
		instructions.MovShift0,
	))
}
