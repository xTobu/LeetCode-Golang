package problem0647

func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		// 統計以 i 為中心的回文總數
		count += countPalindrome(s, i, i)
		// 統計以當 i, i+1 為左右邊際時的中心，其回文總數
		count += countPalindrome(s, i, i+1)
	}
	return count
}

// 設定左右邊界，以中點為中心，往兩側擴展，
// 判斷是否為相同，相同即代表回文
// 回傳該中心往外的回文的總數
func countPalindrome(s string, left int, right int) int {
	var count int
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}
