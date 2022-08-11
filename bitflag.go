package bitflag

type Flag uint

const Default Flag = 0

func New(flags ...Flag) Flag {
	flag := Default
	flag.Set(flags...)
	return flag
}

// Set is used to enable the flags
func (f *Flag) Set(flags ...Flag) {
	for _, flag := range flags {
		*f |= 1 << flag
	}
}

// Remove is used to disable the flags
func (f *Flag) Remove(flags ...Flag) {
	for _, flag := range flags {
		*f &^= 1 << flag
	}
}

// Toggle is used to enable/disable the flags
func (f *Flag) Toggle(flags ...Flag) {
	for _, flag := range flags {
		*f ^= 1 << flag
	}
}

// Has is used to check whether the flags exist or not. It will return true if all flags exist.
// It will return false if there is one flag doesn't exist
func (f *Flag) Has(flags ...Flag) bool {
	if len(flags) == 0 {
		return false
	}

	reqFlags := Default
	reqFlags.Set(flags...)
	return (*f & reqFlags) != 0
}
