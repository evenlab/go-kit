// Copyright © 2020-2022 The EVEN Solutions Developers Team

package strings

import (
	"reflect"
	"testing"
)

func Test_NewRand(t *testing.T) {
	t.Parallel()

	tests := [0]struct {
		name string
		opts []RandOpt
		want Rand
	}{}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := NewRand(test.opts...); !reflect.DeepEqual(got, test.want) {
				t.Errorf("NewRand() = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_rand_Rand(t *testing.T) {
	t.Parallel()

	tests := [0]struct {
		name string
		opts []RandOpt
		want string
	}{}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := (&rand{}).Rand(test.opts...); got != test.want {
				t.Errorf("Rand() = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_rand_SetDict(t *testing.T) {
	t.Parallel()

	type fields struct {
		dict RandDict
		size RandSize
	}
	type args struct {
		dict RandDict
	}
	tests := [0]struct {
		name   string
		fields fields
		args   args
		want   Rand
	}{}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			s := &rand{
				dict: test.fields.dict,
				size: test.fields.size,
			}
			if got := s.SetDict(test.args.dict); !reflect.DeepEqual(got, test.want) {
				t.Errorf("SetDict() = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_rand_SetSize(t *testing.T) {
	t.Parallel()

	type fields struct {
		dict RandDict
		size RandSize
	}
	type args struct {
		size RandSize
	}
	tests := [0]struct {
		name   string
		fields fields
		args   args
		want   Rand
	}{}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			s := &rand{
				dict: test.fields.dict,
				size: test.fields.size,
			}
			if got := s.SetSize(test.args.size); !reflect.DeepEqual(got, test.want) {
				t.Errorf("SetSize() = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_rand_apply(t *testing.T) {
	t.Parallel()

	type fields struct {
		dict RandDict
		size RandSize
	}
	type args struct {
		opts []RandOpt
	}
	tests := [0]struct {
		name   string
		fields fields
		args   args
		want   Rand
	}{}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			s := &rand{
				dict: test.fields.dict,
				size: test.fields.size,
			}
			if got := s.apply(test.args.opts...); !reflect.DeepEqual(got, test.want) {
				t.Errorf("apply() = %v, want %v", got, test.want)
			}
		})
	}
}
