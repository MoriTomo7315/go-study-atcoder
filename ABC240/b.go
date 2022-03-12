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
	n, _ := strconv.Atoi(sc.Text())

	numList := []string{}
	judge := make(map[string]bool)

	// ２行目入力

	sc.Scan()
	// 区切り文字 "", " "に注意
	var inputs = strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		if !judge[inputs[i]] {
			judge[inputs[i]] = true
			numList = append(numList, inputs[i])
		}
	}

	fmt.Printf("%d", len(numList))
}
