// Copyright © 2020-2021 The EVEN Solutions Developers Team

package errors_test

import (
	"errors"
	"log"
	"testing"

	. "github.com/platsko/go-kit/errors"
)

func Benchmark_wrapper_Error(tb *testing.B) {
	err := WrapErr(wrapErrorMsg, New(testErrorMsg))
	if err == nil {
		log.Fatal("want error interface but got nil value")
	}
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = err.Error()
	}
}

func Benchmark_wrapper_Unwrap(tb *testing.B) {
	wrapErr := WrapErr(wrapErrorMsg, New(testErrorMsg))
	err, ok := wrapErr.(Wrapper)
	if !ok {
		log.Fatal("got not wrapper interface")
	}
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
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
			want: errors.New(testErrorMsg + errDelimiter + wrapErrorMsg).Error(),
		},
		{
			name:    "ERR",
			want:    errors.New(wrapErrorMsg + errDelimiter + testErrorMsg).Error(),
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := WrapErr(testErrorMsg, New(wrapErrorMsg)).Error()
			if (got != test.want) != test.wantErr {
				t.Errorf("Error() got: %v | want: %v", got, test.wantErr)
			}
		})
	}
}

func Test_wrapper_Unwrap(t *testing.T) {
	t.Parallel()

	testErr := New(testErrorMsg)
	wrapErr := WrapErr(wrapErrorMsg, testErr)

	tests := [2]struct {
		name    string
		testErr error
		wrapErr error
		want    bool
	}{
		{
			name:    "TRUE",
			testErr: testErr,
			wrapErr: WrapErr(wrapErrorMsg, wrapErr),
			want:    true,
		},
		{
			name:    "FALSE",
			testErr: testErr,
			wrapErr: WrapErr(wrapErrorMsg, New(testErrorMsg)),
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
