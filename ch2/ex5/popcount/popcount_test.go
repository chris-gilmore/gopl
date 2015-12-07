package popcount

import (
	"testing"
)

// -- Tests --

func TestPopCount(t *testing.T) {
	if PopCount(0x1234567890ABCDEF) != 32 {
		t.Error(`PopCount(0x1234567890ABCDEF) != 32`)
	}
}

func TestByShifting(t *testing.T) {
	if ByShifting(0x1234567890ABCDEF) != 32 {
		t.Error(`popCount2(0x1234567890ABCDEF) != 32`)
	}
}

func TestByClearing(t *testing.T) {
	if ByClearing(0x1234567890ABCDEF) != 32 {
		t.Error(`popCount2(0x1234567890ABCDEF) != 32`)
	}
}

func TestBitCount(t *testing.T) {
	if BitCount(0x1234567890ABCDEF) != 32 {
		t.Error(`popCount2(0x1234567890ABCDEF) != 32`)
	}
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByShifting(0x1234567890ABCDEF)
	}
}

func BenchmarkByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0x1234567890ABCDEF)
	}
}

// $ go test -bench=.
// BenchmarkPopCount-4          200000000       8.14 ns/op
// BenchmarkByShifting-4        20000000        79.7 ns/op
// BenchmarkByClearing-4        50000000        28.5 ns/op
// BenchmarkBitCount-4          500000000       3.20 ns/op
