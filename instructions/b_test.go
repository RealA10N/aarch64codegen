package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"github.com/stretchr/testify/assert"
)

func AssertExpectedBranchInstruction(t *testing.T, expected string, offset int32) {
	off26, err := immediates.NewOffset26Align4(offset)
	assert.NoError(t, err)
	instruction := instructions.B(off26)
	AssertExpectedInstruction(t, expected, instruction)
}

func TestBranch0(t *testing.T) {
	AssertExpectedBranchInstruction(t, "B #0", 0)
}

func TestBranch1(t *testing.T) {
	AssertExpectedBranchInstruction(t, "B #4", 4)
}

func TestBranchMinus1(t *testing.T) {
	AssertExpectedBranchInstruction(t, "B #-4", -4)
}

func TestBranchMax(t *testing.T) {
	AssertExpectedBranchInstruction(t, "B #134217724", 0x7FFFFFC)
}

func TestBranchMin(t *testing.T) {
	AssertExpectedBranchInstruction(t, "B #-134217728", -0x8000000)
}
