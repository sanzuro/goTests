package main

import "fmt"

func main() {
	var pr, pmb, pab float64
	fmt.Scan(&pmb)
	fmt.Scan(&pab)
	fmt.Scan(&pr)
	fmt.Printf("%9.8f", pr*(pmb*(1-pab)+pab*(1-pmb)))
}
