// Copyright © 2020-2021 The EVEN Solutions Developers Team

package equal_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/evenlab/go-kit/bytes"

	"github.com/evenlab/go-kit/equal"
)

func Benchmark_BasicEqual(b *testing.B) {
	const size = 1024
	blob := bytes.RandBytes(size)
	equaler := equal.NewEqualer(blob)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = equal.BasicEqual(equaler, equaler)
	}
}

func Benchmark_NewEqualer(b *testing.B) {
	const size = 1024
	blob := bytes.RandBytes(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = equal.NewEqualer(blob)
	}
}

func Benchmark_equaler_Raw(b *testing.B) {
	const size = 1024
	blob := bytes.RandBytes(size)
	equaler := equal.NewEqualer(blob)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := equaler.Raw(); err != nil {
			log.Fatal(err)
		}
	}
}

func Test_BasicEqual(t *testing.T) {
	t.Parallel()

	const size = 1024
	blob := bytes.RandBytes(size)

	tests := [7]struct {
		name string
		equ1 equal.Equaler
		equ2 equal.Equaler
		want bool
	}{
		{
			name: "TRUE",
			equ1: equal.NewEqualer(blob),
			equ2: equal.NewEqualer(blob),
			want: true,
		},
		{
			name: "nil_Equalers_TRUE",
			equ1: nil,
			equ2: nil,
			want: true,
		},
		{
			name: "FALSE",
			equ1: equal.NewEqualer(blob),
			equ2: equal.NewEqualer(bytes.RandBytes(size)),
		},
		{
			name: "nil_Equ1_FALSE",
			equ1: nil,
			equ2: equal.NewEqualer(blob),
		},
		{
			name: "nil_Equ2_FALSE",
			equ1: equal.NewEqualer(blob),
			equ2: nil,
		},
		{
			name: "zero_size_Equ1_FALSE",
			equ1: equal.NewEqualer(make([]byte, 0)),
			equ2: equal.NewEqualer(blob),
		},
		{
			name: "zero_size_Equ2_FALSE",
			equ1: equal.NewEqualer(blob),
			equ2: equal.NewEqualer(make([]byte, 0)),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := equal.BasicEqual(test.equ1, test.equ2); got != test.want {
				t.Errorf("BasicEqual() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_NewEqualer(t *testing.T) {
	t.Parallel()

	const size = 1024
	blob := bytes.RandBytes(size)

	tests := [1]struct {
		name string
		blob []byte
		want []byte
	}{
		{
			name: "OK",
			blob: blob,
			want: blob,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			equaler := equal.NewEqualer(test.blob)
			if got, _ := equaler.Raw(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("NewEqualer() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_equaler_Raw(t *testing.T) {
	t.Parallel()

	const size = 1024
	blob := bytes.RandBytes(size)

	tests := [2]struct {
		name    string
		equaler equal.Equaler
		want    []byte
		wantErr bool
	}{
		{
			name:    "OK",
			equaler: equal.NewEqualer(blob),
			want:    blob,
		},
		{
			name:    "ERR",
			equaler: equal.NewEqualer(nil),
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.equaler.Raw()
			if (err != nil) != test.wantErr {
				t.Errorf("Raw() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Raw() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}
