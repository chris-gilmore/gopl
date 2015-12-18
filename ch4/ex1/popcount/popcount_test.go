package popcount

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

var c1 = sha256.Sum256([]byte("x"))
var c2 = sha256.Sum256([]byte("X"))

// -- Tests --

func TestHammingDistance(t *testing.T) {
	if HammingDistance(c1, c2) != 125 {
		t.Error(`HammingDistance(c1, c2) != 125`)
	}
}

// -- Benchmarks --

func BenchmarkHammingDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HammingDistance(c1, c2)
	}
}

// -- Examples --

func ExampleHammingDistance() {
	fmt.Println(HammingDistance(c1, c2))
	// Output:
	// 125
}

// $ go test -bench=.
// BenchmarkHammingDistance-4   20000000        68.0 ns/op
