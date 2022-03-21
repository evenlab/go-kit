// Copyright © 2020-2022 The EVEN Solutions Developers Team

package crypto_test

import (
	"reflect"
	"testing"

	cc "github.com/libp2p/go-libp2p-core/crypto"

	"github.com/evenlab/go-kit/crypto"
)

func Benchmark_NewPrivateKey(b *testing.B) {
	ki, _ := mockCryptoKeyPair(crypto.Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = crypto.NewPrivateKey(ki)
	}
}

func Benchmark_privateKey_Algo(b *testing.B) {
	prKey, _ := mockGenerateKeyPair(crypto.Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = prKey.Algo()
	}
}

func Benchmark_privateKey_PublicKey(b *testing.B) {
	prKey, _ := mockGenerateKeyPair(crypto.Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = prKey.PublicKey()
	}
}

func Benchmark_privateKey_Sign(b *testing.B) {
	signable, prKey := mockSignable(crypto.Ed25519, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := prKey.Sign(signable); err != nil {
			b.Fatal(err)
		}
	}
}

func Test_NewPrivateKey(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name string
			ki   cc.PrivKey
			want crypto.PrivateKey
		}
		testList []testCase
	)

	algos := crypto.GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		ki, _ := mockCryptoKeyPair(algo)
		tests = append(tests, testCase{
			name: name + "_OK",
			ki:   ki,
			want: crypto.NewPrivateKey(ki),
		})
	}

	tests = append(tests, testCase{
		name: "nil_OK",
		ki:   nil,
		want: crypto.NewPrivateKey(nil),
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := crypto.NewPrivateKey(test.ki); !reflect.DeepEqual(got, test.want) {
				t.Errorf("NewPrivateKey() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_privateKey_Algo(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name  string
			prKey crypto.PrivateKey
			want  crypto.Algo
		}
		testList []testCase
	)

	algos := crypto.GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		prKey, _ := mockGenerateKeyPair(algo)
		tests = append(tests, testCase{
			name:  name + "_OK",
			prKey: prKey,
			want:  algo,
		})
	}

	tests = append(tests, testCase{
		name:  "UNKNOWN_OK",
		prKey: crypto.NewPrivateKey(nil),
		want:  crypto.UNKNOWN,
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.prKey.Algo(); got != test.want {
				t.Errorf("Algo() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_privateKey_PublicKey(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name   string
			prKey  crypto.PrivateKey
			want   crypto.PublicKey
			wantEq bool
		}
		testList []testCase
	)

	algos := crypto.GetAlgos()
	tests := make(testList, 0, algos.Len()*2)
	for name, algo := range algos {
		notEq, _ := mockGenerateKeyPair(algo)
		prKey, pbKey := mockGenerateKeyPair(algo)
		tests = append(tests, testCase{
			name:   name + "_OK",
			prKey:  prKey,
			want:   pbKey,
			wantEq: true,
		}, testCase{
			name:  name + "_not_EQUAL",
			prKey: notEq, // there is not paired private key
			want:  pbKey,
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.prKey.PublicKey(); test.wantEq && !reflect.DeepEqual(got, test.want) {
				t.Errorf("PublicKey() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_privateKey_Sign(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name       string
			prKey      crypto.PrivateKey
			signable   crypto.Signable
			wantErr    bool
			wantVerify bool
		}
		testList []testCase
	)

	algos := crypto.GetAlgos()
	tests := make(testList, 0, algos.Len()+2)
	for name, algo := range algos {
		signable, prKey := mockSignable(algo, 1024)
		tests = append(tests, testCase{
			name:       name + "_OK",
			prKey:      prKey,
			signable:   signable,
			wantVerify: true,
		})
	}

	signable, prKey := mockSignable(crypto.Ed25519, 1024)
	tests = append(tests, testCase{
		name:     "nil_Signable_ERR",
		prKey:    prKey,
		signable: nil,
		wantErr:  true,
	}, testCase{
		name:     "nil_pointer_PrivateKey_ERR",
		prKey:    crypto.NewPrivateKey(nil),
		signable: signable,
		wantErr:  true,
	}, testCase{
		name:     "nil_pointer_Signable_ERR",
		prKey:    prKey,
		signable: crypto.NewSignable(nil),
		wantErr:  true,
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := test.prKey.Sign(test.signable)
			if (err != nil) != test.wantErr {
				t.Errorf("Sign() error: %v | want: %v", err, test.wantErr)
				return
			}

			got, err := test.prKey.PublicKey().Verify(test.signable)
			if (err != nil) != test.wantErr {
				t.Errorf("Sign() error: %v | want: %v", err, test.wantErr)
				return
			}
			if got != test.wantVerify {
				t.Errorf("Sign() got: %v | want: %v", got, test.wantVerify)
			}
		})
	}
}
