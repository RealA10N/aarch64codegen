package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/immediates"
)

type Bl uint32

func NewBl(offset immediates.Offset26Align4) Bl {
	return Bl(0b100101<<26 | offset.Binary())
}

func (i Bl) Offset() immediates.Offset26Align4 {
	return immediates.Offset26Align4(i & 0x3FFFFFF)
}

func (i Bl) String() string {
	return fmt.Sprintf("bl %s", i.Offset())
}

func (i Bl) Binary() uint32 {
	return uint32(i)
}
