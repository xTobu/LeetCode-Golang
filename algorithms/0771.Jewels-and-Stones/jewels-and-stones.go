package problem0771

func numJewelsInStones(J string, S string) int {

	// 建立一個 map 作為 Hash table
	var m = make(map[byte]bool)

	// 將寶石列表 J 的每個值，
	// 作為哈希表的 key 存進去。
	for i := range J {
		m[J[i]] = true
	}

	// 遍歷石頭列表 S ，
	// 逐格去和雜湊表做比對，並記錄總數count
	var count int
	for i := range S {
		if _, ok := m[S[i]]; ok {
			count++
		}
	}

	return count
}
