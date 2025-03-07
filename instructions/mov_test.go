package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestMovz1(t *testing.T) {
	AssertExpectedInstruction(t, "MOVZ X0, #1234, LSL #32", instructions.MOVZ(
		registers.X0,
		immediates.Immediate16(1234),
		instructions.MovShift32,
	))
}

func TestMovz2(t *testing.T) {
	AssertExpectedInstruction(t, "MOVZ X30, #65535", instructions.MOVZ(
		registers.X30,
		immediates.Immediate16(0xFFFF),
		instructions.MovShift0,
	))
}

func TestMovzXZR(t *testing.T) {
	// Test MOVZ with XZR as destination
	AssertExpectedInstruction(t, "MOVZ XZR, #1234", instructions.MOVZ(
		registers.XZR,
		immediates.Immediate16(1234),
		instructions.MovShift0,
	))
}

func TestMovk1(t *testing.T) {
	AssertExpectedInstruction(t, "MOVK X10, #0, LSL #48", instructions.MOVK(
		registers.X10,
		immediates.Immediate16(0),
		instructions.MovShift48,
	))
}

func TestMovk2(t *testing.T) {
	AssertExpectedInstruction(t, "MOVK X29, #65534, LSL #16", instructions.MOVK(
		registers.X29,
		immediates.Immediate16(0xFFFE),
		instructions.MovShift16,
	))
}

func TestMovkXZR(t *testing.T) {
	// Test MOVK with XZR as destination
	AssertExpectedInstruction(t, "MOVK XZR, #4321", instructions.MOVK(
		registers.XZR,
		immediates.Immediate16(4321),
		instructions.MovShift0,
	))
}

func TestMovn1(t *testing.T) {
	AssertExpectedInstruction(t, "MOVN X2, #48879, LSL #32", instructions.MOVN(
		registers.X2,
		immediates.Immediate16(0xBEEF),
		instructions.MovShift32,
	))
}

func TestMovn2(t *testing.T) {
	AssertExpectedInstruction(t, "MOVN X21, #0", instructions.MOVN(
		registers.X21,
		immediates.Immediate16(0),
		instructions.MovShift0,
	))
}

func TestMovnXZR(t *testing.T) {
	// Test MOVN with XZR as destination
	AssertExpectedInstruction(t, "MOVN XZR, #5678", instructions.MOVN(
		registers.XZR,
		immediates.Immediate16(5678),
		instructions.MovShift0,
	))
}
