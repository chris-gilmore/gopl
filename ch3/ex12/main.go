// Anagram ...
package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	if b := isAnagram(os.Args[1], os.Args[2]); b {
		fmt.Printf("%q is an anagram of %q\n", os.Args[1], os.Args[2])
		os.Exit(0)
	} else {
		fmt.Printf("%q is not an anagram of %q\n", os.Args[1], os.Args[2])
		os.Exit(1)
	}
}

// isAnagram ...
func isAnagram(s string, t string) bool {
	netRuneCount := make(map[rune]int)

	for _, r := range s {
		if unicode.IsLetter(r) {
			netRuneCount[r]++
		}
	}
	for _, r := range t {
		if unicode.IsLetter(r) {
			netRuneCount[r]--
		}
	}

	for _, v := range netRuneCount {
		if v != 0 {
			return false
		}
	}
	return true
}
