package bitflag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	FlagOne Flag = iota
	FlagTwo
	FlagThree
	FlagFour
)

func TestAll(t *testing.T) {
	f := Default
	shouldHave(t, f)
	shouldNotHave(t, f, FlagOne, FlagTwo, FlagThree, FlagFour)

	f.Set(FlagOne)
	shouldHave(t, f, FlagOne)
	shouldNotHave(t, f, FlagTwo, FlagThree, FlagFour)

	f.Remove(FlagOne)
	shouldHave(t, f)
	shouldNotHave(t, f, FlagOne, FlagTwo, FlagThree, FlagFour)

	f.Set(FlagOne, FlagTwo)
	shouldHave(t, f, FlagOne, FlagTwo)
	shouldNotHave(t, f, FlagThree, FlagFour)

	f.Remove(FlagOne)
	shouldHave(t, f, FlagTwo)
	shouldNotHave(t, f, FlagOne, FlagThree, FlagFour)

	f.Toggle(FlagThree)
	shouldHave(t, f, FlagTwo, FlagThree)
	shouldNotHave(t, f, FlagOne, FlagFour)

	f.Toggle(FlagThree)
	shouldHave(t, f, FlagTwo)
	shouldNotHave(t, f, FlagOne, FlagThree, FlagFour)

	f.Toggle(FlagTwo, FlagThree)
	shouldHave(t, f, FlagThree)
	shouldNotHave(t, f, FlagOne, FlagTwo, FlagFour)

	f.Toggle(FlagTwo, FlagThree)
	shouldHave(t, f, FlagTwo)
	shouldNotHave(t, f, FlagOne, FlagThree, FlagFour)

	f.Toggle(FlagTwo)
	shouldHave(t, f)
	shouldNotHave(t, f, FlagOne, FlagTwo, FlagThree, FlagFour)

	f.Remove(FlagOne, FlagTwo)
	shouldHave(t, f)
	shouldNotHave(t, f, FlagOne, FlagTwo, FlagThree, FlagFour)

	f.Set(FlagOne, FlagTwo)
	shouldHave(t, f, FlagOne, FlagTwo)
	shouldNotHave(t, f, FlagThree, FlagFour)

	f.Remove(FlagOne, FlagTwo)
	shouldHave(t, f)
	shouldNotHave(t, f, FlagOne, FlagTwo, FlagThree, FlagFour)

	f2 := New(FlagOne, FlagTwo)
	shouldHave(t, f2, FlagOne, FlagTwo)
	shouldNotHave(t, f2, FlagThree, FlagFour)
}

func shouldHave(t *testing.T, src Flag, flags ...Flag) {
	if len(flags) == 0 {
		assert.False(t, src.Has(flags...))
		return
	}

	assert.True(t, src.Has(flags...))
	for _, flag := range flags {
		assert.True(t, src.Has(flag))
	}
}

func shouldNotHave(t *testing.T, src Flag, flags ...Flag) {
	if len(flags) == 0 {
		assert.False(t, src.Has(flags...))
		return
	}

	for _, flag := range flags {
		assert.False(t, src.Has(flag))
	}
	assert.False(t, src.Has(flags...))
}
