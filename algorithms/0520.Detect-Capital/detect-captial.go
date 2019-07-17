package problem0520

import "strings"

func detectCapitalUse(word string) (answer bool) {

	// 根據題目要求，
	// 先轉成符合題意的字串形式
	// 之後用三者去做比對，與一相同則為答案。
	strLower := strings.ToLower(word)
	strUpper := strings.ToUpper(word)
	strNormal := strings.ToUpper(word[:1]) + strings.ToLower(word[1:])

	// 配合 Go 精神，使用 Switch 語句
	switch word {
	case strLower, strUpper, strNormal:
		answer = true
	default:
		answer = false
	}

	return answer
	// return strLower == word || strUpper == word || strNormal == word
}
