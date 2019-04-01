package main

import "fmt"

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
	fmt.Printf("%v", fab(2)/(fab(0)*fab(2)))
}
