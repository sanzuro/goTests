package main

import (
	"fmt"
	"sort"
)

func main() {
	var t, n, q int
	fmt.Scan(&t)
	for z := 0; z < t; z++ {
		fmt.Scan(&n)
		fmt.Scan(&q)
		trip := make([]int, 0)
		var temp int
		for t := 0; t < n; t++ {
			fmt.Scan(&temp)
			trip = append(trip, temp)
		}
		sort.Ints(trip)
		ans := 0
		for t := 0; t < n-1; t++ {
			ans += q * -(trip[t] - trip[t+1])
		}
		fmt.Printf("%v", ans)
	}
}
