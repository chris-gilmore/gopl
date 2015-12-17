package main

import "fmt"

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	fmt.Printf("KB = %.0e\n", float64(KB)) // KB = 1e+03
	fmt.Printf("MB = %.0e\n", float64(MB)) // MB = 1e+06
	fmt.Printf("GB = %.0e\n", float64(GB)) // GB = 1e+09
	fmt.Printf("TB = %.0e\n", float64(TB)) // TB = 1e+12
	fmt.Printf("PB = %.0e\n", float64(PB)) // PB = 1e+15
	fmt.Printf("EB = %.0e\n", float64(EB)) // EB = 1e+18
	fmt.Printf("ZB = %.0e\n", float64(ZB)) // ZB = 1e+21
	fmt.Printf("YB = %.0e\n", float64(YB)) // YB = 1e+24
}
