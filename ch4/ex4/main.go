package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s right by two positions.
	s = Rotate(s, -2)
	fmt.Println(s) // "[4 5 0 1 2 3]"
}

// Rotate rotates a slice of ints in either direction
func Rotate(s []int, l int) []int {
	l %= len(s)

	if l < 0 {
		l += len(s)
	}

	s = append(s, s[:l]...)

	return s[l:]
}
