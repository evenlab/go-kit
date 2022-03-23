// Copyright © 2020-2022 The EVEN Solutions Developers Team

package strings

type (
	// RandSize describes size option of string randomizer.
	RandSize uint
)

var (
	// Make sure RandSize implements RandOpt interface.
	_ RandOpt = (*RandSize)(nil)
)

// apply implements RandOpt interface.
func (s RandSize) apply(r *rand) {
	r.setSize(s)
}
