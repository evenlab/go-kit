// Copyright © 2020-2022 The EVEN Solutions Developers Team

package strings_test

import (
	"reflect"
	"testing"

	"github.com/evenlab/go-kit/strings"
)

func Test_rand_Rand(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		opts []strings.RandOpt
		want string
	}{
		{
			name: "OK",
			opts: []strings.RandOpt{
				strings.RandSize(10),
				strings.RandDict("0"),
			},
			want: "0000000000",
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rand := strings.NewRand()
			if got := rand.Rand(test.opts...); got != test.want {
				t.Errorf("Rand() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_rand_SetDict(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		dict strings.RandDict
		want strings.Rand
	}{
		{
			name: "OK",
			dict: "0",
			want: strings.NewRand(strings.RandDict("0")),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rand := strings.NewRand()
			if got := rand.SetDict(test.dict); !reflect.DeepEqual(got, test.want) {
				t.Errorf("SetDict() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_rand_SetSize(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		size strings.RandSize
		want strings.Rand
	}{
		{
			name: "OK",
			size: 10,
			want: strings.NewRand(strings.RandSize(10)),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rand := strings.NewRand()
			if got := rand.SetSize(test.size); !reflect.DeepEqual(got, test.want) {
				t.Errorf("SetSize() got: %v | want: %v", got, test.want)
			}
		})
	}
}
