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
	var a [10]int
	// 1 行目の標準出力
	sc.Scan()
	var initValue = strings.Split(sc.Text(), " ")
	for i := 0; i < 10; i++ {
		a[i], _ = strconv.Atoi(initValue[i])
	}

	num := a[0]
	num = a[num]
	num = a[num]

	fmt.Printf("%d", num)
}
