package immediates

import "fmt"

type Condition uint32

const (
	ConditionEq Condition = iota // Equal
	ConditionNe                  // Not Equal
	ConditionCs                  // Carry Set
	ConditionCc                  // Carry Clear
	ConditionMi                  // Minus/Negative
	ConditionPl                  // Plus/Positive
	ConditionVs                  // Overflow Set
	ConditionVc                  // Overflow Clear
	ConditionHi                  // Higher
	ConditionLs                  // Lower or Same
	ConditionGe                  // Greater than or Equal
	ConditionLt                  // Less Than
	ConditionGt                  // Greater Than
	ConditionLe                  // Less than or Equal
	ConditionAl                  // Always
	ConditionNv                  // Never
)

func (c Condition) Binary() uint32 {
	return uint32(c)
}

func (c Condition) Validate() error {
	if c <= ConditionNv {
		return fmt.Errorf("Condition %d out of range", c)
	}
	return nil
}

func (c Condition) String() string {
	switch c {
	case ConditionEq:
		return "eq"
	case ConditionNe:
		return "ne"
	case ConditionCs:
		return "cs"
	case ConditionCc:
		return "cc"
	case ConditionMi:
		return "mi"
	case ConditionPl:
		return "pl"
	case ConditionVs:
		return "vs"
	case ConditionVc:
		return "vc"
	case ConditionHi:
		return "hi"
	case ConditionLs:
		return "ls"
	case ConditionGe:
		return "ge"
	case ConditionLt:
		return "lt"
	case ConditionGt:
		return "gt"
	case ConditionLe:
		return "le"
	case ConditionAl:
		return "al"
	case ConditionNv:
		return "nv"
	default:
		return "??"
	}
}
