// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)      // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int   // count of lengths of UTF-8 encodings
	invalid := 0                      // count of invalid UTF-8 characters
	catCounts := make(map[string]int) // counts per Unicode category
	unknown := 0                      // count of characters of unknown category

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		switch {
		case unicode.IsLetter(r):
			catCounts["Letter"]++
		case unicode.IsDigit(r):
			catCounts["Digit"]++
		case unicode.IsSymbol(r):
			catCounts["Symbol"]++
		case unicode.IsPunct(r):
			catCounts["Punct"]++
		case unicode.IsSpace(r):
			catCounts["Space"]++
		default:
			unknown++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

	fmt.Print("\ncat\tcount\n")
	for cat, n := range catCounts {
		fmt.Printf("%s\t%d\n", cat, n)
	}
	if unknown > 0 {
		fmt.Printf("\n%d characters of unknown category\n", unknown)
	}
}
