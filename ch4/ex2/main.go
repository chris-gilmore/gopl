package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var hash = flag.String("hash", "sha256", "sha256 | sha384 | sha512")
	flag.Parse()

	s, _ := ioutil.ReadAll(os.Stdin)

	switch *hash {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256(s))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384(s))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512(s))
	default:
		fmt.Printf("%x\n", sha256.Sum256(s))
	}
}
