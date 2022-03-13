package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	atcoder
	Educational DP Contest A-Flog 2
	https://atcoder.jp/contests/dp/tasks/dp_b
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

	k := nextInt()

	// // for h_0 - h_n-1
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = nextInt()
	}

	// solve by DP
	dp := make([]int, n)
	// init value 
	dp[0] = 0
	dp[1] = abs(h[1]-h[0])
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1]+abs(h[i]-h[i-1])
		for j := 2; j < k+1; j++ {
			if (i-j > -1) {
				dp[i] = min(dp[i],dp[i-j]+abs(h[i]-h[i-j]))
			}
		}
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