package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	X int
	Y int
	// Direction string
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	// buffer作成
	buf := make([]byte, 4096)
	// scannerに最大2*10^5Bでbuffer登録
	sc.Buffer(buf, 600000)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	// // checkのためのmap
	// same_cnt := make(map[int]int, n)

	// N 人の座標を取得
	pos_list := make([]*pos, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		x, _ := strconv.Atoi(inputs[0])
		y, _ := strconv.Atoi(inputs[1])
		p := &pos{
			X: x,
			Y: y,
			// Direction: "",
		}
		pos_list[i] = p
	}

	// 各人の方向を取得
	sc.Scan()
	S := strings.Split(sc.Text(), "")

	// R向いてるxminの座標のマップ(key: y座標, value: xが最小であるindex)
	right_xmin_map := map[int]int{}
	// L向いてるxmaxの座標のマップ(key: y座標, value: xが最大であるindex)
	left_xmax_map := map[int]int{}

	isConflict := false
	for i := 0; i < n; i++ {

		if S[i] == "R" {
			// Rを向いている場合、Y座標が同じ、かつLを向いていてX座標が大きいものが存在するのであれば衝突
			_, ok := left_xmax_map[pos_list[i].Y]
			if ok && pos_list[i].X < left_xmax_map[pos_list[i].Y] {
				isConflict = true
				break
			}
		}

		if S[i] == "L" {
			// Lを向いている場合、Y座標が同じ、かつRを向いていてX座標が小さいものが存在するのであれば衝突
			_, ok := right_xmin_map[pos_list[i].Y]
			if ok && pos_list[i].X > right_xmin_map[pos_list[i].Y] {
				isConflict = true
				break
			}
		}

		// R向いてるxminの座標のマップの更新
		if S[i] == "R" {
			_, ok := right_xmin_map[pos_list[i].Y]
			if ok {
				right_xmin_map[pos_list[i].Y] = min(right_xmin_map[pos_list[i].Y], pos_list[i].X)
			} else {
				right_xmin_map[pos_list[i].Y] = pos_list[i].X
			}
		} else { // L向いてるxmaxの座標のマップの更新
			_, ok := left_xmax_map[pos_list[i].Y]
			if ok {
				left_xmax_map[pos_list[i].Y] = max(left_xmax_map[pos_list[i].Y], pos_list[i].X)
			} else {
				left_xmax_map[pos_list[i].Y] = pos_list[i].X
			}
		}
	}

	if isConflict {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
