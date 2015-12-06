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

func TestPopCount2(t *testing.T) {
	if popCount2(0x1234567890ABCDEF) != 32 {
		t.Error(`popCount2(0x1234567890ABCDEF) != 32`)
	}
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount2(0x1234567890ABCDEF)
	}
}

// $ go test -bench=.
// BenchmarkPopCount-4          200000000       7.92 ns/op
// BenchmarkPopCount2-4         100000000       10.9 ns/op
