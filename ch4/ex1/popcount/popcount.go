// Package popcount counts set bits.
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// HammingDistance returns the population count (number of set bits) of (x XOR y).
func HammingDistance(x [32]byte, y [32]byte) int {
	c := 0
	for i := 0; i < 32; i++ {
		c += int(pc[x[i]^y[i]])
	}
	return c
}
