// Copyright © 2020-2022 The EVEN Solutions Developers Team

package zero_test

import (
	"testing"

	"github.com/evenlab/go-kit/bytes"

	"github.com/evenlab/go-kit/zero"
)

func Benchmark_Bytea28(b *testing.B) {
	b224 := [28]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zero.Bytea28(&b224)
	}
}

func Benchmark_Bytea32(b *testing.B) {
	b256 := [32]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zero.Bytea32(&b256)
	}
}

func Benchmark_Bytea64(b *testing.B) {
	b512 := [64]byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zero.Bytea64(&b512)
	}
}

func Benchmark_Bytes(b *testing.B) {
	blob := bytes.RandBytes(1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zero.Bytes(blob)
	}
}

func Test_Bytea28(t *testing.T) {
	t.Parallel()

	b224, b := [28]byte{}, bytes.RandBytes(28)
	copy(b224[:], b)

	tests := [1]struct {
		name string
		b224 [28]byte
	}{
		{
			name: "OK",
			b224: b224,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			zero.Bytea28(&test.b224)
			for i := range test.b224 {
				if test.b224[i] != 0 {
					t.Errorf("Bytea28() got: %v | want: [28]byte zeros filled", test.b224)
					return
				}
			}
		})
	}
}

func Test_Bytea32(t *testing.T) {
	t.Parallel()

	b256, b := [32]byte{}, bytes.RandBytes(32)
	copy(b256[:], b)

	tests := [1]struct {
		name string
		b256 [32]byte
	}{
		{
			name: "OK",
			b256: b256,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			zero.Bytea32(&test.b256)
			for i := range test.b256 {
				if test.b256[i] != 0 {
					t.Errorf("Bytea32() got: %v | want: [32]byte zeros filled", test.b256)
					return
				}
			}
		})
	}
}

func Test_Bytea64(t *testing.T) {
	t.Parallel()

	b512, b := [64]byte{}, bytes.RandBytes(64)
	copy(b512[:], b)

	tests := [1]struct {
		name string
		b512 [64]byte
	}{
		{
			name: "OK",
			b512: b512,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			zero.Bytea64(&test.b512)
			for i := range test.b512 {
				if test.b512[i] != 0 {
					t.Errorf("Bytea64() got: %v | want: [64]byte zeros filled", test.b512)
					return
				}
			}
		})
	}
}

func Test_Bytes(t *testing.T) {
	t.Parallel()

	blob := bytes.RandBytes(1024)
	tests := [1]struct {
		name string
		blob []byte
	}{
		{
			name: "OK",
			blob: blob,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			zero.Bytes(test.blob)
			for i := range test.blob {
				if test.blob[i] != 0 {
					t.Errorf("Bytes() got: %v | want: []byte zeros filled", test.blob)
					return
				}
			}
		})
	}
}
