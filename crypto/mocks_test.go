// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto_test

import (
	cc "github.com/libp2p/go-libp2p-core/crypto"

	"github.com/evenlab/go-kit/bytes"

	"github.com/evenlab/go-kit/crypto"
)

func mockCryptoKeyPair(algo crypto.Algo) (cc.PrivKey, cc.PubKey) {
	bits := -1
	if algo == crypto.RSA {
		bits = 2048
	}

	prKey, pbKey, err := cc.GenerateKeyPair(int(algo), bits)
	if err != nil {
		panic(err)
	}

	return prKey, pbKey
}

func mockGenerateKeyPair(algo crypto.Algo) (crypto.PrivateKey, crypto.PublicKey) {
	prKey, pbKey, err := crypto.GenerateKeyPair(algo)
	if err != nil {
		panic(err)
	}

	return prKey, pbKey
}

func mockSignable(algo crypto.Algo, size int) (*crypto.SignableStub, crypto.PrivateKey) {
	prKi, pbKi := mockCryptoKeyPair(algo)
	signable := crypto.SignableStub{Blob: bytes.RandBytes(size)}

	h256, err := signable.Hash()
	if err != nil {
		panic(err)
	}

	sign, err := prKi.Sign(h256[:])
	if err != nil {
		panic(err)
	}

	signable.Sign = crypto.NewSignature(sign)
	signable.PbKey = crypto.NewPublicKey(pbKi)

	return &signable, crypto.NewPrivateKey(prKi)
}

func mockSignature(algo crypto.Algo) (crypto.Signature, crypto.PrivateKey) {
	signable, prKey := mockSignable(algo, 1024)

	return signable.Sign, prKey
}
