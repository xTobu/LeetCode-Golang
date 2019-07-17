package problem0657

/**
 * 假設起點為 (0, 0)
 * 定義出 x, y 來計算移動路徑
 */
func judgeCircle(moves string) bool {

	// 若總步數為奇數，
	// 則不可能回到起點，
	// 故回傳 false 。
	if len(moves)%2 == 1 {
		return false
	}

	var x, y int
	for _, v := range moves {
		switch v {
		case 'U':
			y++
		case 'D':
			y--
		case 'R':
			x++
		case 'L':
			x--
		}
	}

	// 檢查最後的座標是否為 (0, 0)
	return x == 0 && y == 0
}
