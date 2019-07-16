package problem0485

/**
 * 建立兩個 int (res, count)，
 * 遍歷 nums ，
 * 若遇 1 ，則 count + 1 ，
 * 並和 res 比對，且將 res 設為兩者中的較大者
 * 若不是 1，則歸零 count
 * 最後回傳結果
 */
func findMaxConsecutiveOnes(nums []int) int {
	var res, count int
	for _, v := range nums {
		switch v {
		case 1:
			count++
			res = Max(res, count)
			// Bigger(&res, &count)
		default:
			count = 0
		}
	}
	return res
}

// Max : find the max one out
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/**
 * 傳址不一定比較快，
 * 在資料量很大或複製成本很高的時候，
 * 傳址才有明顯優勢
 */
// // Bigger : change res to the bigger one
// func Bigger(res, count *int) {
// 	if *res < *count {
// 		*res = *count
// 	}
// }
