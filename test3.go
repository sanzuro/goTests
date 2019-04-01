package main

import (
	"fmt"
	"math"
)

var mat [1000000]int
var modulo = 1000000007

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, max int
		max = 0
		fmt.Scan(&n)
		for i := 0; i < n; i++ {
			fmt.Scan(&mat[i])
			if mat[i] > max {
				max = mat[i]
			}
		}
		total := 0
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				total += max * int(math.Abs(float64(mat[i]-mat[j])))
			}
		}
		fmt.Printf("%v", total%modulo)

	}
}
