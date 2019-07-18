package problem0896

/**
 * 判斷是否為單調的數列（統一遞增、遞減、不變）
 * 解題方向為：
 * 從 A 第一個數字，往後遍歷一次 A ，
 * 將當前數和之前一數字相比，判斷為增加還是減少，
 * 當 increasing, decreasing 都變成 false 時，
 * 則代表為不單調數列，故即可回傳 false
 * 都為單調即回傳 true
 */
func isMonotonic(A []int) bool {
	increasing, decreasing := true, true
	for i := 1; i < len(A); i++ {
		switch {
		case A[i] > A[i-1]:
			decreasing = false
		case A[i] < A[i-1]:
			increasing = false
		}

		if decreasing == false && increasing == false {
			return false
		}
	}
	return true
}
