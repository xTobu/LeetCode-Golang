package problem0001

func twoSum(nums []int, target int) []int {
	// 使用 make 建立一個空 map , 並設置 key 和 value 的型態
	temp := make(map[int]int)
	for index, num := range nums {
		// 取 temp[target-num] 的值（某 num 的 index），
		// 如果有值則代表在 for 迴圈的前某次，有儲存過了 target 和當下 num 的 差，
		// 意味在 temp 找得到 num 的另一半，
		// 於是可以回傳兩者的位置，則為答案。
		if value, ok := temp[target-num]; ok {
			return []int{value, index}
		}
		// 如果沒有值，
		// 則將當前的 num 作為 key; 其 index 作為 value ,
		// 存進 temp 裡，
		// 等待 range 的下一個 num 來比對。
		temp[num] = index
	}
	// func 設定了欲回傳型態，
	// 於是沒有結果的話，還是得回傳 nil
	return nil
}
