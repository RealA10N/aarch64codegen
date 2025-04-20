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
		return "lsl"
	case LSR:
		return "lsr"
	case ASR:
		return "asr"
	case ROR:
		return "ror"
	default:
		return "???"
	}
}

func (s ShiftType) Binary() uint32 {
	return uint32(s)
}
