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
	// buffer作成
	buf := make([]byte, 4096)
	// scannerに最大2*10^5Bでbuffer登録
	sc.Buffer(buf, 10000)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	// for A
	sc.Scan()
	a := strings.Split(sc.Text(), " ")

	// for B
	sc.Scan()
	b := strings.Split(sc.Text(), " ")

	// for 1.
	cnt1 := 0
	for i := 0; i < n; i++ {
		if a[i] == b[i] {
			cnt1 = cnt1 + 1
		}
	}

	// for 2.
	cnt2 := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && a[i] == b[j] {
				cnt2 = cnt2 + 1
				break
			}
		}
	}

	fmt.Println(cnt1)
	fmt.Println(cnt2)
}
