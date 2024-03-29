// Copyright © 2020-2022 The EVEN Solutions Developers Team

package bytes_test

import (
	"testing"

	"github.com/evenlab/go-kit/bytes"
)

func Benchmark_RandBytes(b *testing.B) {
	const size = 1024
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bytes.RandBytes(size)
	}
}

func Test_RandBytes(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		size int
	}{
		{
			name: "OK",
			size: 1024,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := bytes.RandBytes(test.size); len(got) != test.size {
				t.Errorf("RandBytes() size: %v | want: %v", got, test.size)
			}
		})
	}
}
