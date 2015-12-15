package bigcmplx

import "math/big"

// BigRat ...
type BigRat struct {
	real big.Rat
	imag big.Rat
}

// Float64 ...
func (z *BigRat) Float64() (r, i float64) {
	r, _ = z.real.Float64()
	i, _ = z.imag.Float64()
	return
}

// SetFloat64 ...
func (z *BigRat) SetFloat64(r, i float64) *BigRat {
	z.real.SetFloat64(r)
	z.imag.SetFloat64(i)
	return z
}

// SetInt64 ...
func (z *BigRat) SetInt64(r, i int64) *BigRat {
	z.real.SetInt64(r)
	z.imag.SetInt64(i)
	return z
}

// Set ...
func (z *BigRat) Set(x *BigRat) *BigRat {
	z.real.Set(&x.real)
	z.imag.Set(&x.imag)
	return z
}

// Add ...
func (z *BigRat) Add(x, y *BigRat) *BigRat {
	z.real.Add(&x.real, &y.real)
	z.imag.Add(&x.imag, &y.imag)
	return z
}

// Sub ...
func (z *BigRat) Sub(x, y *BigRat) *BigRat {
	z.real.Sub(&x.real, &y.real)
	z.imag.Sub(&x.imag, &y.imag)
	return z
}

// Mul ...
// a+bi * c+di
// ac-bd + (bc+ad)i
func (z *BigRat) Mul(x, y *BigRat) *BigRat {
	z.real.Sub(new(big.Rat).Mul(&x.real, &y.real), new(big.Rat).Mul(&x.imag, &y.imag))
	z.imag.Add(new(big.Rat).Mul(&x.imag, &y.real), new(big.Rat).Mul(&x.real, &y.imag))
	return z
}

// Quo ...
// a+bi / c+di
// [a+bi * c-di] / [c+di * c-di] = [ac+bd + (bc-ad)i] / c^2+d^2
func (z *BigRat) Quo(x, y *BigRat) *BigRat {
	divisor := new(big.Rat).Add(new(big.Rat).Mul(&y.real, &y.real), new(big.Rat).Mul(&y.imag, &y.imag))
	z.real.Quo(new(big.Rat).Add(new(big.Rat).Mul(&x.real, &y.real), new(big.Rat).Mul(&x.imag, &y.imag)), divisor)
	z.imag.Quo(new(big.Rat).Sub(new(big.Rat).Mul(&x.imag, &y.real), new(big.Rat).Mul(&x.real, &y.imag)), divisor)
	return z
}
