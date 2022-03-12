package main

import (
	"fmt"
	"math"
)

func main() {

	var h float64

	// 1 行目の標準出力
	fmt.Scan(&h)

	fmt.Printf("%f", math.Sqrt(h*(12800000+h)))
}
