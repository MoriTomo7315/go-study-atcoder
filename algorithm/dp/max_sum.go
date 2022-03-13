package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	https://qiita.com/drken/items/a5e6fe22863b7992efdb
	1. 最大和問題
	をDPで解く

	n個の整数 a[0],a[1],…,a[n−1]a[0],a[1],…,a[n−1] が与えられる。  
	これらの整数から何個かの整数を選んで総和をとったときの、総和の最大値を求めよ。  
	また、何も選ばない場合の総和は 0 であるものとする。

	入力例
	N
	a0 a1 ... an-1 an

	N は 1 以上 10000 以下の整数
	Ai は −1000 以上 1000 以下の整数 (0≤i≤N−1)
*/


func main() {
	sc := bufio.NewScanner(os.Stdin)
	// buffer作成
	buf := make([]byte, 4096)
	// scannerに最大2*10^5Bでbuffer登録
	sc.Buffer(buf, 100000)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	// for a0 - an
	a := make([]int, n)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	for i,v := range inputs {
		v_value, _ := strconv.Atoi(v)
		a[i] = v_value
	}

	// solve by DP
	dp := make([]int, n+1)
	// init value
	dp[0] = 0
	for i := 0; i < n; i++ {
		dp[i+1] = max(dp[i], dp[i]+a[i])
	}

	fmt.Println(dp[n])

}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}