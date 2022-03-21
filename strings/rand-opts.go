// Copyright © 2020-2022 The EVEN Solutions Developers Team

package strings

type (
	// RandOpt option interface of string randomizer.
	RandOpt interface {
		apply(*rand)
	}
)
