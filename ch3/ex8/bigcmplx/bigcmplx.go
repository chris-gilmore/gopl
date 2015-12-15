// Package bigcmplx ...
package bigcmplx

import "math/big"

// NewBigFloat ...
func NewBigFloat(r, i *big.Float) *BigFloat {
	z := new(BigFloat)
	z.real.Copy(r)
	z.imag.Copy(i)
	return z
}
