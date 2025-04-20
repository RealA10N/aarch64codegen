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
		return "EQ"
	case ConditionNe:
		return "NE"
	case ConditionCs:
		return "CS"
	case ConditionCc:
		return "CC"
	case ConditionMi:
		return "MI"
	case ConditionPl:
		return "PL"
	case ConditionVs:
		return "VS"
	case ConditionVc:
		return "VC"
	case ConditionHi:
		return "HI"
	case ConditionLs:
		return "LS"
	case ConditionGe:
		return "GE"
	case ConditionLt:
		return "LT"
	case ConditionGt:
		return "GT"
	case ConditionLe:
		return "LE"
	case ConditionAl:
		return "AL"
	case ConditionNv:
		return "NV"
	default:
		return "??"
	}
}
