package problem0189

func rotate(nums []int, k int) {
	// 取出 nums 的總長
	lenNums := len(nums)
	// 將 k 除與 lenNums 並取其 餘數 設為 k (steps)
	k %= lenNums
	if k == 0 {
		return
	}

	// nums 已 k 為範圍分前後兩部分，並將其相反 append 進 temp
	// 關於後面三個點：
	// func append(slice []Type, elems ...Type) []Type
	// append 第一個可以擺 slice，但往後則是擺和第一個 slice 一樣 Type 的 數個 value
	// 故後面必須加上 ... 將 slice 打散後加入 append 。
	// 參照：
	// https://github.com/zhangyachen/zhangyachen.github.io/issues/137
	// https://coderwall.com/p/tlnhmg/appending-two-slices-in-go-what-s-with-that
	tmp := append(nums[lenNums-k:], nums[0:lenNums-k]...)

	// 將 nums 設置回去
	for index, value := range tmp {
		nums[index] = value
	}

}
