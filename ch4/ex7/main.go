package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := []byte("Hello, 世界")
	UTF8Reverse(s)
	fmt.Println(string(s)) // "界世 ,olleH"
}

// UTF8Reverse reverses the UTF-8-encoded bytes of each character (rune)
// before reversing the entire []byte slice, in-place.
func UTF8Reverse(s []byte) {
	t := s
	for size := 0; len(t) > 0; t = t[size:] {
		_, size = utf8.DecodeRune(t)
		reverse(t[:size])
	}
	reverse(s)
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
