package problem0242

/**
 * 解題方向：
* 計算出 s 每個英文字母出現過幾次，
* 在依照 t 的出現次數去相減，
* 最後去判斷全字母是否都為 0 即可。
*/

func isAnagram(s string, t string) bool {
	// 為英文字母總數暫存而建立一個 26 長度的整數陣列，
	var alphabets [26]int

	// 將字串 s 轉成 byte 型態後遍歷，
	// 將其代表的英文字母數量加一。
	for _, v := range []byte(s) {
		alphabets[v-97]++
	}

	// 將字串 t 轉成 byte 型態後遍歷，
	// 將其代表的英文字母數量減一。
	for _, v := range []byte(t) {
		alphabets[v-97]--
	}

	// 遍歷作為英文字母總數暫存的陣列，
	// 若其值不等於 0 , 則代表有某字母出現次數無法對應，
	// 於是便可證明 s 、 t 不為易位構詞遊戲。
	for _, v := range alphabets {
		if v != 0 {
			return false
		}
	}
	return true
}
