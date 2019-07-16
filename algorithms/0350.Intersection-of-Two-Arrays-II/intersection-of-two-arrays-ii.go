package problem0350

import "sort"

/**
 * 先排序，之後一邊移位一邊比較，
 * 沒有建立 map ，因此記憶體的使用比較少。
 **/
func intersect(nums1 []int, nums2 []int) []int {

	sort.Ints(nums1)
	sort.Ints(nums2)

	var i, j int
	res := []int{}

	// for 迴圈的標準條件為：nums1 或 nums2 其一結束即可
	for i < len(nums1) && j < len(nums2) {

		num1 := nums1[i]
		num2 := nums2[j]

		// 若數到 nums1 的某一數 num1 ，
		// 其值大於當前的 num2 時，
		// nums2 的陣列要往下一個前進。（ j 加一）
		if num1 > num2 {
			j++
		}

		// 若數到 nums2 的某一數 num2 ，
		// 其值大於當前的 num1 時，
		// nums1 的陣列要往下一個前進。（ i 加一）
		if num1 < num2 {
			i++
		}

		// 若 num1 、 num2 相等時，
		// 將當前數字 append 進 res ，
		// 並對 nums1 、 nums2 的陣列都往下個計算（ i 、 j 都加一）
		if num1 == num2 {
			res = append(res, num1)
			i++
			j++
		}

	}

	return res

}
