package immediates

// ShiftType represents the AArch64 shift operations
type ShiftType uint8

const (
	LSL ShiftType = iota // Logical Shift Left
	LSR                  // Logical Shift Right
	ASR                  // Arithmetic Shift Right
	ROR                  // Rotate Right
)

func (s ShiftType) String() string {
	switch s {
	case LSL:
		return "LSL"
	case LSR:
		return "LSR"
	case ASR:
		return "ASR"
	case ROR:
		return "ROR"
	default:
		return "???"
	}
}

func (s ShiftType) Binary() uint32 {
	return uint32(s)
}
