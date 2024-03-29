// Copyright © 2020-2022 The EVEN Solutions Developers Team

package strings

import (
	"strings"
	"sync"

	"github.com/evenlab/go-kit/bytes"
)

type (
	// Rand describes string randomizer interface.
	Rand interface {
		// Rand returns random generated string with settled options.
		Rand(...RandOpt) string

		// SetDict sets dict of string randomizer.
		SetDict(RandDict) Rand

		// SetSize sets size of string randomizer.
		SetSize(RandSize) Rand

		// apply sets options of string randomizer.
		apply(...RandOpt) Rand
	}

	// rand implements Rand interface.
	rand struct {
		dict RandDict
		size RandSize
		sync sync.RWMutex
	}
)

// NewRand returns constructed string randomizer.
func NewRand(opts ...RandOpt) Rand {
	r := &rand{}

	r.apply(opts...)
	if r.dict == "" {
		r.dict = DefaultRandDict
	}

	return r
}

// Rand implements Rand interface.
func (s *rand) Rand(opts ...RandOpt) string {
	s.sync.Lock()
	defer s.sync.Unlock()

	if len(opts) > 0 {
		defer s.apply(s.dict, s.size) // backup current options
		s.apply(opts...)
	}

	size := int(s.size)
	blob, dictSize := bytes.RandBytes(size), len(s.dict)
	for i := 0; i < size; i++ {
		blob[i] = s.dict[int(blob[i])%dictSize]
	}

	builder := strings.Builder{}
	_, _ = builder.Write(blob)

	return builder.String()
}

// SetDict implements Rand interface.
func (s *rand) SetDict(dict RandDict) Rand {
	s.sync.Lock()
	s.dict = dict
	s.sync.Unlock()

	return s
}

// SetSize implements Rand interface.
func (s *rand) SetSize(size RandSize) Rand {
	s.sync.Lock()
	s.size = size
	s.sync.Unlock()

	return s
}

// apply implements Rand interface.
func (s *rand) apply(opts ...RandOpt) Rand {
	for _, opt := range opts {
		opt.apply(s)
	}

	return s
}

// SetDict implements Rand interface.
func (s *rand) setDict(dict RandDict) {
	s.dict = dict
}

// SetSize implements Rand interface.
func (s *rand) setSize(size RandSize) {
	s.size = size
}
