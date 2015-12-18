package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	var hash = flag.String("hash", "sha256", "sha256 | sha384 | sha512")
	flag.Parse()

	switch *hash {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256(fileBytes(os.Stdin)))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384(fileBytes(os.Stdin)))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512(fileBytes(os.Stdin)))
	default:
		fmt.Printf("%x\n", sha256.Sum256(fileBytes(os.Stdin)))
	}
}

func fileBytes(f *os.File) []byte {
	var buf bytes.Buffer
	input := bufio.NewScanner(f)

	input.Split(bufio.ScanBytes)
	for input.Scan() {
		buf.Write(input.Bytes())
	}

	return buf.Bytes()
}
