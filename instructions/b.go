package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/immediates"
)

type Branch uint32

func B(offset immediates.Offset26Align4) Branch {
	return Branch(0b101<<26 | offset.Binary())
}

func (i Branch) Offset() immediates.Offset26Align4 {
	return immediates.Offset26Align4(i & 0x3FFFFFF)
}

func (i Branch) String() string {
	return fmt.Sprintf("B %s", i.Offset())
}

func (i Branch) Binary() uint32 {
	return uint32(i)
}
