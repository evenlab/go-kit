// Copyright © 2020-2022 The EVEN Solutions Developers Team

package strings

type (
	// RandDict describes dictionary option of string randomizer.
	RandDict string
)

var (
	// Make sure RandDict implements RandOpt interface.
	_ RandOpt = (*RandDict)(nil)
)

// apply implements RandOpt interface.
func (s RandDict) apply(r *rand) {
	r.SetDict(s)
}
