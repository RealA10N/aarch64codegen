package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestAnd(t *testing.T) {
	AssertExpectedInstruction(t, "AND X9, X10, X11", instructions.AND(
		registers.X9,
		registers.X10,
		registers.X11,
	))
}

func TestAndWithXZR(t *testing.T) {
	AssertExpectedInstruction(t, "AND XZR, X10, X11", instructions.AND(
		registers.XZR,
		registers.X10,
		registers.X11,
	))

	AssertExpectedInstruction(t, "AND X9, XZR, X11", instructions.AND(
		registers.X9,
		registers.XZR,
		registers.X11,
	))

	AssertExpectedInstruction(t, "AND X9, X10, XZR", instructions.AND(
		registers.X9,
		registers.X10,
		registers.XZR,
	))

	AssertExpectedInstruction(t, "AND XZR, XZR, XZR", instructions.AND(
		registers.XZR,
		registers.XZR,
		registers.XZR,
	))
}
