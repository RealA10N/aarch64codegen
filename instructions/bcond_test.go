package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"github.com/stretchr/testify/assert"
)

func AssertExpectedBcondInstruction(
	t *testing.T,
	expected string,
	cond immediates.Condition,
	offset int32,
) {
	off19, err := immediates.NewOffset19Align4(offset)
	assert.NoError(t, err)
	instruction := instructions.BCOND(cond, off19)
	AssertExpectedInstruction(t, expected, instruction)
}

func TestBcond0(t *testing.T) {
	AssertExpectedBcondInstruction(t, "b.eq #0", immediates.ConditionEq, 0)
}

func TestBcond1(t *testing.T) {
	AssertExpectedBcondInstruction(t, "b.eq #4", immediates.ConditionEq, 4)
}

func TestBcondMinus1(t *testing.T) {
	AssertExpectedBcondInstruction(t, "b.eq #-4", immediates.ConditionEq, -4)
}

func TestBcondMax(t *testing.T) {
	AssertExpectedBcondInstruction(t, "b.eq #1048572", immediates.ConditionEq, 0xFFFFC)
}

func TestBcondMin(t *testing.T) {
	AssertExpectedBcondInstruction(t, "b.eq #-1048576", immediates.ConditionEq, -0x100000)
}

func TestBcondAllConditions(t *testing.T) {
	for _, cond := range []immediates.Condition{
		immediates.ConditionEq,
		immediates.ConditionNe,
		immediates.ConditionCs,
		immediates.ConditionCc,
		immediates.ConditionMi,
		immediates.ConditionPl,
		immediates.ConditionVs,
		immediates.ConditionVc,
		immediates.ConditionHi,
		immediates.ConditionLs,
		immediates.ConditionGe,
		immediates.ConditionLt,
		immediates.ConditionGt,
		immediates.ConditionLe,
	} {
		AssertExpectedBcondInstruction(t, "b."+cond.String()+" #0", cond, 0)
	}
}
