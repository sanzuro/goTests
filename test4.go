package main

import "fmt"

var n, m int
var t [100000]bool

func isfilled() bool {
	fmt.Printf("%v\n", t[:n])
	for i := 0; i < n; i++ {
		if t[i] == false {
			return false
		}
	}
	return true
}

func empty() {
	for i := 0; i < n; i++ {
		t[i] = false
	}
}

func main() {
	fmt.Scan(&n)
	fmt.Scan(&m)
	var e int
	for i := 0; i < m; i++ {
		fmt.Scan(&e)
		t[e-1] = true

		if isfilled() {
			fmt.Printf("1")
			empty()
		} else {
			fmt.Printf("0")
		}
	}
}
