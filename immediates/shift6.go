package immediates

import (
	"fmt"
)

// Shift6 combines a ShiftType with a 6-bit shift amount
type Shift6 struct {
	Type   ShiftType
	Amount Immediate6
}

func NewShift6(shiftType ShiftType, amount Immediate6) Shift6 {
	return Shift6{
		Type:   shiftType,
		Amount: amount,
	}
}

func (s Shift6) Validate() error {
	return s.Amount.Validate()
}

func (s Shift6) String() string {
	return fmt.Sprintf("%s #%d", s.Type, s.Amount)
}
