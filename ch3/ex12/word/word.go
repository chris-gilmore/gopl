// Package word provides utilities for word games.
package word

import "unicode"

// IsAnagram reports whether two strings are anagrams of each other.
func IsAnagram(s string, t string) bool {
	netRuneCount := make(map[rune]int)

	for _, r := range s {
		if unicode.IsLetter(r) {
			netRuneCount[unicode.ToLower(r)]++
		}
	}
	for _, r := range t {
		if unicode.IsLetter(r) {
			netRuneCount[unicode.ToLower(r)]--
		}
	}

	for _, v := range netRuneCount {
		if v != 0 {
			return false
		}
	}
	return true
}
