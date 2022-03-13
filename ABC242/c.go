package main

import (
	"fmt"
)

/*
	atcoder
	ABC242 c 問題
	https://atcoder.jp/contests/abc242/tasks/abc242_c
*/

var (
	mod = 998244353
)

func main() {

	var n int
	fmt.Scan(&n)

	
	dp := make([][]int, 9)
	for i := 0; i < 9; i++ {
		dp[i] = make([]int, n)
	}

	// init value  (1桁目はそれぞれ1通り)
	for i := 0; i < 9; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < 9; j++ {
			if j == 0 {
				dp[j][i] = dp[0][i-1] + dp[1][i-1]
			} else if j == 8 {
				dp[j][i] = dp[8][i-1] + dp[7][i-1]
			} else {
				dp[j][i] = dp[j-1][i-1] + dp[j][i-1] + + dp[j+1][i-1]
			}
			dp[j][i] = dp[j][i] % mod
		}
	}
	var cnt = 0
	for i := 0; i < 9; i++ {
		cnt = cnt + dp[i][n-1]
	}
	ans := cnt % mod
	fmt.Println(ans)

}