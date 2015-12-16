// Comma prints its argument numbers with a comma at each power of 1000.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a floating point string with an optional sign.
func comma(s string) string {
	var buf bytes.Buffer

	if s[0] == '-' || s[0] == '+' {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	p := strings.Index(s, ".")
	if p < 0 {
		p = len(s)
	}

	if p > 0 {
		rem := p % 3
		if rem == 0 {
			rem = 3
		}
		buf.WriteString(s[0:rem])

		for i := rem; i < p; i = i + 3 {
			buf.WriteByte(',')
			buf.WriteString(s[i : i+3])
		}
	}

	buf.WriteString(s[p:])

	return buf.String()
}
