// Copyright © 2020-2022 The EVEN Solutions Developers Team

package merkle

import (
	"github.com/evenlab/go-kit/errors"
)

const (
	ErrMerkleTreeBuiltImproperlyMsg = "merkle tree built improperly"
)

var (
	errMerkleTreeBuiltImproperly = errors.New(ErrMerkleTreeBuiltImproperlyMsg)
)

func ErrMerkleTreeBuiltImproperly() error {
	return errMerkleTreeBuiltImproperly
}
