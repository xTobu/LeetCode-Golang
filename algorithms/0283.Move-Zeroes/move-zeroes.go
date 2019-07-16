package problem0283

/**
 * 比較直覺,
 * 遇到零, 抓出來往後丟
 *
 * 從後面開始遍歷,
 * 如果 nums[i] === 0
 * 便將該位置移除 ( nums.splice(i, 1) ),
 * 最後面推一個 0 回去.
 *
 * 注意：
 * 不能從頭開始找,
 * 因為從頭的話, splice 後會刪除一個元素, 之後 push 加到最後,
 * 這樣又照著遍歷跑時, 有一個元素會被忽略
 *
 * 譬如
 * i = 0 [0, 0, 1] -> [0, 1, 0]
 * i = 1 [0, 1, 0] -> [0, 1, 0]
 *
 * 連續的第二個也是零的話就會被忽略了
 *
 */
func moveZeroes(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			// Golang - remove element in a array
			// 參考：https://stackoverflow.com/questions/25025409/delete-element-in-a-slice
			// 將 nums 用 append 置換
			// 假設要移除的位置 i ， 以 i 將陣列分成頭尾兩部分
			// 再將其用 append 組合起來。
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		}
	}
}

/**
 * 移動所有的零到陣列最後面, 且不能新建多的陣列.
 *
 * 思路為遍歷每個元素,
 * 如果遇到 0 , 就忽略
 *
 * 遇到不是 0 ,
 * 就將該元素 nums[i] 的值放到 nums[position],
 * 並將 position + 1
 *
 * 以次類推,
 * nums 全部遍歷完之後,
 * 全部非零元素就會依照原本順序, 排在陣列最前面了.
 *
 * 之後再從該陣列位置的 position 為起始,
 * 將後面的值改成 0 , 即可.
 *
 * 例如：
 * position = 0, i = 0,
 * [0, 1, 0, 3, 12],
 * nums[i] = 0, 忽略
 *
 * position = 0, i = 1,
 * [0, 1, 0, 3, 12],
 * nums[i] = 1,
 * nums[position] = nums[i]
 * [1, 1, 0, 3, 12]
 * position + 1
 *
 * position = 1, i = 2,
 * [1, 1, 0, 3, 12],
 * nums[i] = 0, 忽略
 *
 * position = 1, i = 3,
 * [1, 1, 0, 3, 12],
 * nums[i] = 3,
 * nums[position] = nums[i]
 * [1, 3, 0, 3, 12]
 * position + 1
 *
 * position = 2, i = 4,
 * [1, 3, 0, 3, 12],
 * nums[i] = 12,
 * nums[position] = nums[i]
 * [1, 3, 12, 3, 12]
 * position + 1
 *
 * 最後將陣列從 position 開始後的元素都置換成 0
 *
 */
func moveZeroes2(nums []int) {

	var position, lengthNums int
	lengthNums = len(nums)

	for _, v := range nums {
		if v != 0 {
			nums[position] = v
			position++
		}
	}

	for i := position; i < lengthNums; i++ {
		nums[i] = 0
	}

}
