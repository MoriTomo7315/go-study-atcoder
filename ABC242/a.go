package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	// 1行目
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	a, _ := strconv.Atoi(inputs[0])
	b, _ := strconv.Atoi(inputs[1])
	c, _ := strconv.Atoi(inputs[2])
	x, _ := strconv.Atoi(inputs[3])

	af := float64(a)
	bf := float64(b)
	cf := float64(c)
	xf := float64(x)

	ans := 1.0
	if xf <= af {
		fmt.Printf("%f", ans)
	} else if xf >= af+1 && xf <= bf {
		ans = cf / (bf - af)
		fmt.Printf("%f", ans)
	} else {
		fmt.Printf("%d", 0)
	}
}
