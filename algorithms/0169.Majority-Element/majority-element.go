package problem0169

// 建立雜湊表
// -
// 以 num 作為雜湊表的 key,
// 遇到 num 存在時 +1 ,
// 依據題意去判斷數量是否有大於長度一半
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = m[v] + 1
		if m[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}

/**
 * 假設第一個元素是眾數，記錄它的數量，遍歷
 * 如果和眾數相等，就遞增，否則就遞減
 * 如果數量等於0，說明不是眾數，重新開始
 *
 * 因為題目有定義：
 * 眾數其數量大於總數的一半；
 * 不會有空陣列且陣列裡一定有眾數。
 *
 * 故可以假設第一個元素是眾數並記錄它的數量，
 * 而後的操作才可以成立
 *
 * 方法超聰明
**/
func majorityElement2(nums []int) int {
	c := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == c {
			count++
		} else {
			count--
		}
		if count < 1 {
			count = 1
			c = nums[i]
		}
	}
	return c
}
