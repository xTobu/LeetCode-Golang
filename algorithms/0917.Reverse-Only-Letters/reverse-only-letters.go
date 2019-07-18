package problem0917

/**
 * 題目為翻轉字串，但只反轉英文字母
 * 反轉字串的方式和 344. Reverse String 一樣，
 * 但因應題目，故建立 isLetter 去判斷欲反轉的左右邊際，
 * 於是都以 left < right 並加上條件 !isLetter 的前提去跑遍歷，
 * 以尋找到都是字母的左右邊際，並將兩者反轉，
 * 全部跑完後，將 []byte 轉型成 string 回傳
 */
func reverseOnlyLetters(S string) string {
	bs := []byte(S)

	for left, right := 0, len(bs)-1; left < right; left, right = left+1, right-1 {
		for left < right && !isLetter(bs[left]) {
			left++
		}
		for left < right && !isLetter(bs[right]) {
			right--
		}
		bs[left], bs[right] = bs[right], bs[left]
	}

	return string(bs)
}

func isLetter(b byte) bool {
	return 'a' <= b && b <= 'z' ||
		'A' <= b && b <= 'Z'
}
