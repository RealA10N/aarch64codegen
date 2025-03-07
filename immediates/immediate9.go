package immediates

import "fmt"

type Immediate9 uint16

func (imm Immediate9) Validate() error {
	if imm > 0x1FF {
		return fmt.Errorf("value %d too large for Immediate9", imm)
	}
	return nil
}

func (imm Immediate9) Binary() uint32 {
	return uint32(imm)
}

func (imm Immediate9) String() string {
	return fmt.Sprintf("#%d", imm)
}
