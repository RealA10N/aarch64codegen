package instructions

import (
	"fmt"
)

// SVC (Supervisor Call) instruction
type Svc uint32

// SVC generates a supervisor call with the given immediate value
func SVC(imm16 uint16) Svc {
	return Svc(0xD4000001 | (uint32(imm16) << 5))
}

func (i Svc) Binary() uint32 {
	return uint32(i)
}

func (i Svc) String() string {
	imm16 := (i >> 5) & 0xFFFF
	return fmt.Sprintf("SVC #%d", imm16)
}
