package main

import "fmt"

func main() {
	data := []string{"one", "two", "two", "two", "three"}
	fmt.Printf("%q\n", dedup(data)) // `["one" "two" "three"]`
}

func dedup(strings []string) []string {
	out := strings[:1] // in-place slice algorithm
	t := strings[0]
	for _, s := range strings[1:] {
		if s != t {
			out = append(out, s)
			t = s
		}
	}
	return out
}
