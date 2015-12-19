package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := []byte("x \n X\r \tA")
	s = Squash(s)
	fmt.Println(s)         // "[120 32 88 32 65]"
	fmt.Println(string(s)) // "x X A"
}

// Squash squashes runs of adjacent white space into a single space
func Squash(s []byte) []byte {
	out := s[:0]

	for spaceRun, size := false, 0; len(s) > 0; s = s[size:] {
		var r rune
		r, size = utf8.DecodeRune(s)
		if unicode.IsSpace(r) {
			if !spaceRun {
				spaceRun = true
				out = append(out, ' ')
			}
		} else {
			spaceRun = false
			out = append(out, s[:size]...)
		}
	}

	return out
}
