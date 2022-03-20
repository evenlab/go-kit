// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto_test

import (
	"reflect"
	"testing"

	"github.com/evenlab/go-kit/crypto"
)

func Benchmark_GetAlgos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = crypto.GetAlgos()
	}
}

func Benchmark_Algo_Type(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = crypto.Ed25519.Type()
	}
}

func Benchmark_Algos_Copy(b *testing.B) {
	list := crypto.GetAlgos()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = list.Copy()
	}
}

func Benchmark_Algos_Len(b *testing.B) {
	list := crypto.GetAlgos()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = list.Len()
	}
}

func Test_GetAlgos(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		want crypto.Algos
	}{
		{
			name: "OK",
			want: crypto.Algos{
				"RSA":       crypto.RSA,
				"Ed25519":   crypto.Ed25519,
				"Secp256k1": crypto.Secp256k1,
				"ECDSA":     crypto.ECDSA,
			},
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := crypto.GetAlgos(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("GetAlgos() got: %#v| want: %#v", got, test.want)
			}
		})
	}
}

func Test_Algo_Type(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name string
			algo crypto.Algo
			want int
		}
		testList []testCase
	)

	algos := crypto.GetAlgos()
	tests := make(testList, 0, algos.Len())
	for name, algo := range algos {
		tests = append(tests, testCase{
			name: name + "_OK",
			algo: algo,
			want: int(algo),
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.algo.Type(); got != test.want {
				t.Errorf("Type() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Algos_Copy(t *testing.T) {
	t.Parallel()

	tests := [2]struct {
		name string
		list crypto.Algos
	}{
		{
			name: "OK",
			list: crypto.GetAlgos(),
		},
		{
			name: "empty_OK",
			list: crypto.Algos{},
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.list.Copy(); !reflect.DeepEqual(got, test.list) {
				t.Errorf("Copy() got: %v | want: %v", got, test.list)
			}
		})
	}
}

func Test_Algos_Len(t *testing.T) {
	t.Parallel()

	algos := crypto.GetAlgos()
	tests := [2]struct {
		name string
		list crypto.Algos
		want int
	}{
		{
			name: "OK",
			list: algos,
			want: len(algos),
		},
		{
			name: "empty_OK",
			list: crypto.Algos{},
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.list.Len(); got != test.want {
				t.Errorf("Len() got: %v | want: %v", got, test.want)
			}
		})
	}
}
