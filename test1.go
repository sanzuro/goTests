package main

import "fmt"

var t int
var sign string
var build = make([]byte, 0)

func main() {
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scan(&n)
		fmt.Scan(&m)
		k := m * n
		for z := 0; z < k; z++ {
			fmt.Scan(&sign)
			g := []byte(sign)
			build = append(build, g[0])
		}
		fmt.Printf("%v\n", build)
		ans := 0
		k = m * (n - 1)
		fmt.Printf("%v %v \n", m, k)
		for z := 0; z < k; z++ {
			if z%(m-1) != 0 || z == 0 {
				fmt.Printf("%v   \n", build[z])
				if build[z] == byte('/') && build[z+1] == byte('\\') && build[z+m] == byte('\\') && build[z+m+1] == byte('/') {
					ans++
				}
			}
		}

		fmt.Printf("%v\n", ans)
	}
}
