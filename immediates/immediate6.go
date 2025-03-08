package immediates

import "fmt"

type Immediate6 uint8

func (imm Immediate6) Validate() error {
	if imm > 0x3F {
		return fmt.Errorf("value %d too large for Immediate6", imm)
	}
	return nil
}

func (imm Immediate6) Binary() uint32 {
	return uint32(imm)
}

func (imm Immediate6) String() string {
	return fmt.Sprintf("#%d", imm)
}
