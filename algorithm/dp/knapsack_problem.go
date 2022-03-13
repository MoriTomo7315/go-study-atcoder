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
	2. ナップサック問題
	をDPで解く

	N 個の品物があります。i 番目の品物はそれぞれ重さと価値が wi,　viとなっています (0≤i≤N−1)。
	これらの品物から重さの総和が W を超えないように選んだときの、価値の総和の最大値を求めてください
	
	入力は次の形式で与えられます。
	N　W
	w_0 v_0 
	w_1 v_1 
	w_{N-1} v_{N-1}w 

	N は 1 以上 100 以下の整数
	W は 1 以上 1000 以下の整数
	wi ,viは 1 以上 100 以下の整数 (0≤i≤N−1)
*/


type Item struct {
	Weight int
	Value int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	// buffer作成
	buf := make([]byte, 4096)
	// scannerに最大2*10^5Bでbuffer登録
	sc.Buffer(buf, 100000)

	// for N, W
	var N, W int
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	N, _ = strconv.Atoi(inputs[0])
	W, _ = strconv.Atoi(inputs[1])

	// for item
	items := make([]*Item, N)
	for i:=0; i < N; i++ {
		sc.Scan()
		itemInputs := strings.Split(sc.Text(), " ")
		w, _ := strconv.Atoi(itemInputs[0])
		v, _ := strconv.Atoi(itemInputs[1])
		item := &Item{
			Weight: w,
			Value: v,
		}
		items[i] = item
	}

	// solve by DP

	// init dp array
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}
	// init value
	dp[0][0] = 0

	for i, item := range items {
		for w := 0; w < W+1; w++ {
			if item.Weight <= w {
				dp[i+1][w] = max(dp[i][w-item.Weight] + item.Value, dp[i][w])
			} else {
				dp[i+1][w] = dp[i][w]
			}
		}
	}

	fmt.Println(dp[N][W])

}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}