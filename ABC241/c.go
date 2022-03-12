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
	// boardの初期化
	board := make([][]string, n)
	for i := range board {
		board[i] = make([]string, n)
	}

	// boardに入力
	for i := 0; i < n; i++ {
		sc.Scan()
		var inputs = strings.Split(sc.Text(), "")
		for j := 0; j < n; j++ {
			board[i][j] = inputs[j]
		}
	}

	answer := "No"
	cnt := 0
	// boardの検査
	// fmt.Printf(board[0][0])
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// fmt.Printf("始点board[%d][%d]: ", i, j)
			// 横の検査
			cnt = 0
			if j+5 < n {
				for k := 0; k < 6; k++ {
					if board[i][j+k] == "#" {
						cnt = cnt + 1
					}
					if cnt > 3 {
						answer = "Yes"
						break
					}
				}
			}

			// 縦の検査
			cnt = 0
			if i+5 < n {
				for k := 0; k < 6; k++ {
					if board[i+k][j] == "#" {
						cnt = cnt + 1
					}
				}
				if cnt > 3 {
					answer = "Yes"
					break
				}
			}
			// fmt.Printf("%d, ", cnt)
			// 右下斜めの検査
			cnt = 0
			if i+5 < n && j+5 < n {
				for k := 0; k < 6; k++ {
					if board[i+k][j+k] == "#" {
						cnt = cnt + 1
					}
				}
				if cnt > 3 {
					answer = "Yes"
					break
				}
			}
			// 左下斜めの検査
			cnt = 0
			if i-5 >= 0 && j+5 < n {
				for k := 0; k < 6; k++ {
					if board[i-k][j+k] == "#" {
						cnt = cnt + 1
					}
					if cnt > 3 {
						answer = "Yes"
						break
					}
				}
			}
			// fmt.Printf("%d\n", cnt)
		}
		if answer == "Yes" {
			break
		}
	}
	fmt.Printf(answer)
}
