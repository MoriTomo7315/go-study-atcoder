package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	atcoder
	Educational DP Contest A-Flog 1
	https://atcoder.jp/contests/dp/tasks/dp_a
*/

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}


func main() {

	sc.Split(bufio.ScanWords)
	n := nextInt()

	// // for h_0 - h_n-1
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = nextInt()
	}

	// solve by DP
	dp := make([]int, n)
	// init value
	dp[0] = 0
	dp[1] = abs(h[0]-h[1])
	for i := 2 ; i < n; i++ {
		dp[i] = min(dp[i-1] + abs(h[i-1]-h[i]), dp[i-2] + abs(h[i-2]-h[i]))
	}

	fmt.Println(dp[n-1])

}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}