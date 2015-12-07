// Package popcount counts set bits.
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// ByShifting counts bits by shifting its argument through 64 bit positions,
// testing the rightmost bit each time.
func ByShifting(x uint64) (count int) {
	for i := 0; i < 64; i++ {
		count += int(byte(x & 1))
		x >>= 1
	}
	return
}

// ByClearing clears rightmost non-zero bit.
func ByClearing(x uint64) (count int) {
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return
}

// BitCount ... Hacker's Delight, Figure 5-2.
func BitCount(x uint64) int {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}
