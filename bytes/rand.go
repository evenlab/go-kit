// Copyright © 2020-2022 The EVEN Solutions Developers Team

package bytes

import (
	"crypto/rand"
)

// RandBytes returns random generated bytes with given size.
func RandBytes(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b) //nolint:errcheck

	return b
}
