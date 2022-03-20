// Copyright © 2020-2021 The EVEN Solutions Developers Team

package errors_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	errKIT "github.com/evenlab/go-kit/errors"
)

const (
	testErrorMsg = "test error"
	wrapErrorMsg = "wrap error"
)

var (
	wrapFmtLibFormat = "%s" + defaultDelimiter + "%w"
)

func Benchmark_As(b *testing.B) {
	testErr := errKIT.New(testErrorMsg)
	wrapErr := errKIT.WrapErr(wrapErrorMsg, testErr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = errKIT.As(wrapErr, &testErr)
	}
}

func Benchmark_Is(b *testing.B) {
	testErr := errKIT.New(testErrorMsg)
	wrapErr := errKIT.WrapErr(wrapErrorMsg, testErr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = errKIT.Is(wrapErr, testErr)
	}
}

func Benchmark_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = errKIT.New(testErrorMsg)
	}
}

func Benchmark_Unwrap(b *testing.B) {
	wrapErr := errKIT.WrapErr(wrapErrorMsg, errKIT.New(testErrorMsg))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = errKIT.Unwrap(wrapErr)
	}
}

func Benchmark_WrapErr(b *testing.B) {
	wrapErr := errKIT.New(wrapErrorMsg)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = errKIT.WrapErr(testErrorMsg, wrapErr)
	}
}

func Benchmark_WrapStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = errKIT.WrapStr(testErrorMsg, wrapErrorMsg)
	}
}

func Benchmark_FmtErrorf_Wrap(b *testing.B) {
	wrapErr := errKIT.New(wrapErrorMsg)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Errorf(wrapFmtLibFormat, testErrorMsg, wrapErr) //nolint:forbidigo
	}
}

func Test_As(t *testing.T) {
	t.Parallel()

	testErr := errKIT.New(testErrorMsg)
	wrapErr := errKIT.WrapErr(wrapErrorMsg, testErr)

	tests := [1]struct {
		name    string
		testErr error
		wrapErr error
		want    bool
	}{
		{
			name:    "TRUE",
			testErr: testErr,
			wrapErr: wrapErr,
			want:    true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := errKIT.As(test.wrapErr, &test.testErr); got != test.want {
				t.Errorf("As() got: %v | want: %v", got, test.want)
				return
			}
			if !test.want {
				return
			}
			if !reflect.DeepEqual(test.testErr, test.wrapErr) {
				t.Errorf("As() got: %#v | want: %#v", test.testErr, test.wrapErr)
			}
		})
	}
}

func Test_Is(t *testing.T) {
	t.Parallel()

	testErr := errKIT.New(testErrorMsg)
	wrapErr := errKIT.WrapErr(wrapErrorMsg, testErr)

	tests := [2]struct {
		name    string
		testErr error
		wrapErr error
		want    bool
	}{
		{
			name:    "TRUE",
			testErr: testErr,
			wrapErr: wrapErr,
			want:    true,
		},
		{
			name:    "FALSE",
			testErr: testErr,
			wrapErr: errKIT.WrapErr(wrapErrorMsg, errKIT.New(testErrorMsg)),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := errKIT.Is(test.wrapErr, test.testErr); got != test.want {
				t.Errorf("Is() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_New(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		text string
		want string
	}{
		{
			name: "EQUAL",
			text: testErrorMsg,
			want: errors.New(testErrorMsg).Error(),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := errKIT.New(test.text).Error(); got != test.want {
				t.Errorf("New() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Unwrap(t *testing.T) {
	t.Parallel()

	testErr := errKIT.New(testErrorMsg)
	wrapErr := errKIT.WrapErr(wrapErrorMsg, testErr)

	tests := [4]struct {
		name string
		err  error
		want error
	}{
		{
			name: "<nil>",
			err:  nil,
			want: nil,
		},
		{
			name: "no_wrap_OK",
			err:  testErr,
			want: nil,
		},
		{
			name: "once_wrap_OK",
			err:  wrapErr,
			want: testErr,
		},
		{
			name: "double_wrap_OK",
			err:  errKIT.WrapErr(wrapErrorMsg, wrapErr),
			want: wrapErr,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := errKIT.Unwrap(test.err); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Unwrap() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_WrapErr(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		text string
		wrap error
		want string
	}{
		{
			name: "EQUAL",
			text: testErrorMsg,
			wrap: errors.New(wrapErrorMsg),
			want: errors.New(testErrorMsg + defaultDelimiter + wrapErrorMsg).Error(),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := errKIT.WrapErr(test.text, test.wrap).Error(); got != test.want {
				t.Errorf("WrapErr() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_WrapStr(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		text string
		wrap string
		want string
	}{
		{
			name: "EQUAL",
			text: testErrorMsg,
			wrap: wrapErrorMsg,
			want: errors.New(testErrorMsg + defaultDelimiter + wrapErrorMsg).Error(),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := errKIT.WrapStr(test.text, test.wrap).Error(); got != test.want {
				t.Errorf("WrapStr() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}
