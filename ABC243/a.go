package main

import (
	"fmt"
)

func main() {

	// 1行目
	var v, a, b, c int
	fmt.Scan(&v, &a, &b, &c)
	for v >= 0 {
		if v < a {
			fmt.Println("F")
			break
		}
		v = v - a

		if v < b {
			fmt.Println("M")
			break
		}

		v = v - b
		if v < c {
			fmt.Println("T")
			break
		}
		v = v - c
	}
}
