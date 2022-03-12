package main

import (
	"fmt"
)

func main() {

	var a, b int

	// 1 行目の標準出力
	fmt.Scan(&a, &b)

	var isYes bool = false
	if a == 1 {
		if b == 2 || b == 10 {
			isYes = true
		}
	} else if a == 10 {
		if b == 1 || b == 9 {
			isYes = true
		}
	} else {
		if a-b == 1 || a-b == -1 {
			isYes = true
		}
	}

	if isYes {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}
}
