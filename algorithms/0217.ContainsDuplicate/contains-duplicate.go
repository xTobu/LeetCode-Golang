package problem0217

import "sort"

// 解法一：建立 map 去作為 hash 表
// 依照其 key 去檢查是否該數曾出現過。
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = true
	}
	return false
}

// 解法二：先 sort ，再和下一個位置判斷是否重複
func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

/**
 * 兩個解法在 LeetCode 上 Runtime 都差不多，
 * 但因為第一個方法有建立了 map ，
 * 所以在 Memory 的使用上會高出只操作陣列的方法二。
 */
