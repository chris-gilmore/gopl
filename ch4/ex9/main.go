// Wordfreq computes counts of words.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}

	fmt.Printf("word\tcount\n")
	for w, n := range counts {
		fmt.Printf("%s\t%d\n", w, n)
	}
}
