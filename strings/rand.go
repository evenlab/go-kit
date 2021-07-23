// Copyright © 2020-2021 The EVEN Solutions Developers Team

package strings

import (
	"github.com/evenlab/go-kit/bytes"
	"strings"
)

// RandString returns random generated string with given size.(With mutex)
func RandString_KIT30(size int) string {
	rwDictMutex.Lock()
	blob, builder, dictSize := bytes.RandBytes(size), strings.Builder{}, len(dictRandChars)
	for i := 0; i < size; i++ {
		blob[i] = dictRandChars[int(blob[i])%dictSize]
	}
	rwDictMutex.Unlock()
	_, _ = builder.Write(blob)

	return builder.String()
}

// RandString returns random generated string with given size.
func RandString(size int) string {
	s := GetDictRand()
	blob, builder, dictSize := bytes.RandBytes(size), strings.Builder{}, len(s)
	for i := 0; i < size; i++ {
		blob[i] = s[int(blob[i])%dictSize]
	}

	_, _ = builder.Write(blob)

	return builder.String()
}
