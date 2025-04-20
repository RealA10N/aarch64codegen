package immediates

type SetFlags bool

const (
	DoNotSetFlags SetFlags = false
	DoSetFlags    SetFlags = true
)

func (s SetFlags) Binary() uint32 {
	if s {
		return 1 << 29
	} else {
		return 0
	}
}

func (s SetFlags) String() string {
	if s {
		return "s"
	} else {
		return ""
	}
}

func SetFlagsFromBinary(binary uint32) SetFlags {
	if (binary & (1 << 29)) != 0 {
		return DoSetFlags
	} else {
		return DoNotSetFlags
	}
}
