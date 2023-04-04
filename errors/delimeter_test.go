// Copyright © 2020-2022 The EVEN Solutions Developers Team

package errors_test

import (
	"testing"

	"github.com/evenlab/go-kit/errors"
)

var (
	defaultDelimiter = errors.GetDelimiter()
)

func Benchmark_GetDelimiter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = errors.GetDelimiter()
	}
}

func Benchmark_SetDelimiter(b *testing.B) {
	delim := errors.GetDelimiter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		errors.SetDelimiter(delim)
	}
}

func Test_GetDelimiter(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		want string
	}{
		{
			name: "Default",
			want: defaultDelimiter,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := errors.GetDelimiter(); got != test.want {
				t.Errorf("GetDelimiter() got: \"%v\" | want: \"%v\"", got, test.want)
			}
		})
	}
}

func Test_SetDelimiter(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name  string
		delim string
	}{
		{
			name:  "OK",
			delim: " | ",
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			delim := errors.GetDelimiter()  // save current delimiter
			errors.SetDelimiter(test.delim) // set test delimiter
			got := errors.GetDelimiter()    // get test delimiter
			errors.SetDelimiter(delim)      // restore previous delimiter

			if got != test.delim { // check test delimiter
				t.Errorf("SetDelimiter() got: %v | want: %v", got, test.delim)
			}
		})
	}
}
