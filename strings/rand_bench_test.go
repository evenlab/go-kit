// Copyright © 2020-2022 The EVEN Solutions Developers Team

package strings_test

import (
	"testing"

	"github.com/evenlab/go-kit/strings"
)

func Benchmark_NewRandString(b *testing.B) {
	size := strings.RandSize(1024)
	s := strings.NewRand(strings.DefaultRandDict, size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Rand()
	}
}

func Benchmark_RandString(b *testing.B) {
	size := strings.RandSize(1024)
	s := strings.NewRand()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Rand(strings.DefaultRandDict, size)
	}
}
