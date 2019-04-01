package main

import (
	"fmt"
	"math"
)

var modulo = 1000000007

func fab(a int) int {
	if a == 0 {
		return 1
	}
	if a == 1 {
		return 1
	}
	return a * fab(a-1)
}

func main() {
	var t int

	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Printf("%v\n", int(math.Pow(2, float64(i)))%modulo)
	}
}
