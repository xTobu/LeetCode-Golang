package utils

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// ListNode2Ints convert *ListNode to []int
func ListNode2Ints(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Ints2ListNode convert []int to *ListNode
func Ints2ListNode(nums []int) *ListNode {
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
