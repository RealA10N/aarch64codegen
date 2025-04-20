package immediates

import "fmt"

type Offset19Align4 int32

func NewOffset19Align4(offset int32) (Offset19Align4, error) {
	if offset%4 != 0 {
		return 0, fmt.Errorf("offset must be a multiple of 4")
	}

	offset /= 4
	if offset < -0x40000 || offset > 0x3FFFF {
		return 0, fmt.Errorf("offset overflows 19 bits")
	}

	return Offset19Align4(offset), nil
}

func (o Offset19Align4) Binary() uint32 {
	return uint32(o) & 0x7FFFF
}

func (o Offset19Align4) Signed() int32 {
	return (int32(o.Binary()) << 13) >> 13
}

func (o Offset19Align4) String() string {
	return fmt.Sprintf("#%d", o.Signed()*4)
}
