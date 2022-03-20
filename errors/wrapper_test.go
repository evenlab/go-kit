// Copyright © 2020-2021 The EVEN Solutions Developers Team

package errors_test

import (
	"errors"
	"log"
	"testing"

	kit "github.com/evenlab/go-kit/errors"
)

func Benchmark_wrapper_Error(b *testing.B) {
	err := kit.WrapErr(wrapErrorMsg, kit.New(testErrorMsg))
	if err == nil {
		log.Fatal("want error interface but got nil value")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = err.Error()
	}
}

func Benchmark_wrapper_Unwrap(b *testing.B) {
	wrapErr := kit.WrapErr(wrapErrorMsg, kit.New(testErrorMsg))
	err, ok := wrapErr.(kit.Wrapper) // nolint: errorlint
	if !ok {
		log.Fatal("got not wrapper interface")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = err.Unwrap()
	}
}

func Test_wrapper_Error(t *testing.T) {
	t.Parallel()

	tests := [2]struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name: "OK",
			want: errors.New(testErrorMsg + defaultDelimiter + wrapErrorMsg).Error(),
		},
		{
			name:    "ERR",
			want:    errors.New(wrapErrorMsg + defaultDelimiter + testErrorMsg).Error(),
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := kit.WrapErr(testErrorMsg, kit.New(wrapErrorMsg)).Error()
			if (got != test.want) != test.wantErr {
				t.Errorf("Error() got: %v | want: %v", got, test.wantErr)
			}
		})
	}
}

func Test_wrapper_Unwrap(t *testing.T) {
	t.Parallel()

	testErr := kit.New(testErrorMsg)
	wrapErr := kit.WrapErr(wrapErrorMsg, testErr)

	tests := [2]struct {
		name    string
		testErr error
		wrapErr error
		want    bool
	}{
		{
			name:    "TRUE",
			testErr: testErr,
			wrapErr: kit.WrapErr(wrapErrorMsg, wrapErr),
			want:    true,
		},
		{
			name:    "FALSE",
			testErr: testErr,
			wrapErr: kit.WrapErr(wrapErrorMsg, kit.New(testErrorMsg)),
			want:    false,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := errors.Is(test.wrapErr, test.testErr); got != test.want {
				t.Errorf("Unwrap() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}
