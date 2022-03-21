// Copyright © 2020-2022 The EVEN Solutions Developers Team

package bytes

import (
	"github.com/evenlab/go-kit/errors"
)

const (
	ErrVectorsNotSameSizeMsg = "vectors must be the same size"
	ErrVectorZeroSizeMsg     = "vectors cannot be zero size"
)

var (
	errVectorsNotSameSize = errors.New(ErrVectorsNotSameSizeMsg)
	errVectorZeroSize     = errors.New(ErrVectorZeroSizeMsg)
)

func ErrVectorsNotSameSize() error {
	return errVectorsNotSameSize
}

func ErrVectorZeroSize() error {
	return errVectorZeroSize
}
