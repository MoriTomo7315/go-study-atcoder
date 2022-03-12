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

	// 1 行目の標準入力
	sc.Scan()
	var initValue = strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(initValue[0])
	m, _ := strconv.Atoi(initValue[1])
	var a []int
	var b []int

	// 2 行目の標準入力
	var value int
	sc.Scan()
	var aValue = strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(aValue[i])
		a = append(a, value)
	}

	// 3 行目の標準入力
	sc.Scan()
	var bValue = strings.Split(sc.Text(), " ")
	for i := 0; i < m; i++ {
		value, _ = strconv.Atoi(bValue[i])
		b = append(b, value)
	}

	answer := "Yes"

	for i := 0; i < m; i++ {
		for j := 0; j < len(a); j++ {
			if b[i] == a[j] {
				a = append(a[:j], a[j+1:]...)
				answer = "Yes"
				break
			}
			answer = "No"
		}
		if answer == "No" {
			break
		}
	}

	fmt.Printf(answer)
}
