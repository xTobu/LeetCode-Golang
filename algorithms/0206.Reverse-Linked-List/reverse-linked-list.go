package problem0206

import (
	"github.com/xTobu/LeetCode-Golang/lib/utils"
)

// ListNode is tree's node
type ListNode = utils.ListNode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		// 建立一個 temp 用來暫存 head.Next，
		tmp := head.Next
		// head.Next 改為 空的 *ListNode
		head.Next = prev
		// prev 設為 head.Next 已更新過的 head
		prev = head
		// 將 head 以 temp 當依據，設成原先的 head.Next
		head = tmp
	}
	return prev
}

/**
 * 第一次迴圈
 * -------------
 * head: 1->2->3->NULL
 * tmp : 2->3->NULL
 *
 * head.Next: NULL
 * head: 1->NULL
 * prev: 1->NULL
 *
 *
 * 第二次迴圈
 * -------------
 * head: 2->3->NULL
 * tmp : 3->NULL
 *
 * head.Next: 1->NULL
 * head: 2->1->NULL
 * prev: 2->1->NULL
 *
 *
 * 第三次迴圈
 * -------------
 * head: 3->NULL
 * tmp : NULL
 *
 * head.Next: 2->1->NULL
 * head: 3->2->1->NULL
 * prev: 3->2->1->NULL
 *
 *
 * 第四次迴圈
 * -------------
 * head: tmp
 * return
 * prev: 3->2->1->NULL
 *
 *
 * ========================
 *
 * 將來源的其 Next 改成 prev，
 * 因此 prev 是每次迴圈時的 head.Val 值，
 * 加上 head.Next 為從前幾次迴圈累積下來的舊 prev
 *
 * 進而達成 ListNode 從頭反轉回去的動作，
 * 以 prev 置換 Next 的方式堆疊起來
 *
 */

// convert *ListNode to []int
func l2s(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// convert []int to *ListNode
func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}
