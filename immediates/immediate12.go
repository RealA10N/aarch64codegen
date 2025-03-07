package immediates

import "fmt"

type Immediate12 uint16

func (imm Immediate12) Validate() error {
	if imm > 0xFFF {
		return fmt.Errorf("value %d too large for Immediate12", imm)
	}
	return nil
}

func (imm Immediate12) Binary() uint32 {
	return uint32(imm)
}

func (imm Immediate12) String() string {
	return fmt.Sprintf("#%d", imm)
}
