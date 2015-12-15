package bigcmplx

import "math/big"

// BigFloat ...
type BigFloat struct {
	real big.Float
	imag big.Float
}

// Float64 ...
func (z *BigFloat) Float64() (r, i float64) {
	r, _ = z.real.Float64()
	i, _ = z.imag.Float64()
	return
}

// Set ...
func (z *BigFloat) Set(x *BigFloat) *BigFloat {
	z.real.Set(&x.real)
	z.imag.Set(&x.imag)
	return z
}

// SetFloat64 ...
func (z *BigFloat) SetFloat64(r, i float64) *BigFloat {
	z.real.SetFloat64(r)
	z.imag.SetFloat64(i)
	return z
}

// Add ...
func (z *BigFloat) Add(x, y *BigFloat) *BigFloat {
	z.real.Add(&x.real, &y.real)
	z.imag.Add(&x.imag, &y.imag)
	return z
}

// Sub ...
func (z *BigFloat) Sub(x, y *BigFloat) *BigFloat {
	z.real.Sub(&x.real, &y.real)
	z.imag.Sub(&x.imag, &y.imag)
	return z
}

// Mul ...
// a+bi * c+di
// ac-bd + (bc+ad)i
func (z *BigFloat) Mul(x, y *BigFloat) *BigFloat {
	z.real.Sub(new(big.Float).Mul(&x.real, &y.real), new(big.Float).Mul(&x.imag, &y.imag))
	z.imag.Add(new(big.Float).Mul(&x.imag, &y.real), new(big.Float).Mul(&x.real, &y.imag))
	return z
}

// Quo ...
// a+bi / c+di
// [a+bi * c-di] / [c+di * c-di] = [ac+bd + (bc-ad)i] / c^2+d^2
func (z *BigFloat) Quo(x, y *BigFloat) *BigFloat {
	divisor := new(big.Float).Add(new(big.Float).Mul(&y.real, &y.real), new(big.Float).Mul(&y.imag, &y.imag))
	z.real.Quo(new(big.Float).Add(new(big.Float).Mul(&x.real, &y.real), new(big.Float).Mul(&x.imag, &y.imag)), divisor)
	z.imag.Quo(new(big.Float).Sub(new(big.Float).Mul(&x.imag, &y.real), new(big.Float).Mul(&x.real, &y.imag)), divisor)
	return z
}
