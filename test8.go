package main

import "fmt"

var size = int('z' - 'a')
var a = make([]byte, 0)
var b = make([]int, size)

func function() {
	flag := 0
	for i := 0; i < size; i++ {
		if b[i] >= 2 {
			fmt.Printf("Yes\n")
			flag = 1
		}
	}
	if flag == 0 {
		fmt.Printf("No\n")
	}
}

func main() {
	var i int
	fmt.Scan(&i)
	for t := 0; t < i; t++ {
		c := make([]byte, 0)
		for j := 0; j < size; j++ {
			b[j] = 0
		}
		fmt.Scan(&c)
		copy(a, c)
		for j := 0; j < len(a); j++ {
			b[a[j]-'a']++
		}
		function()
	}
}
