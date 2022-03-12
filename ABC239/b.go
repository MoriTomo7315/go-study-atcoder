package main

import (
	"fmt"
)

func main() {

	var x int64

	// 1 行目の標準出力
	fmt.Scan(&x)

	r := x % 10
	a := x / 10

	if r == 0 || x >= 0 {
		fmt.Printf("%d", a)
	} else {
		fmt.Printf("%d", a-1)
	}
}
