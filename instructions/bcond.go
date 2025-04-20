package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/immediates"
)

type Bcond uint32

func BCOND(cond immediates.Condition, offset immediates.Offset19Align4) Bcond {
	return Bcond(
		(0b1010100 << 24) |
			(offset.Binary() << 5) |
			(cond.Binary()),
	)
}

func (i Bcond) Offset() immediates.Offset19Align4 {
	return immediates.Offset19Align4((i >> 5) & 0x7FFFF)
}

func (i Bcond) Condition() immediates.Condition {
	return immediates.Condition(i & 0b1111)
}

func (i Bcond) String() string {
	return fmt.Sprintf("b.%s %s", i.Condition(), i.Offset())
}

func (i Bcond) Binary() uint32 {
	return uint32(i)
}
