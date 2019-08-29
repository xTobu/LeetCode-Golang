package problem0021

import (
	"github.com/xTobu/LeetCode-Golang/lib/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode ...
type ListNode = utils.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	switch {
	case l1 == nil:
		return l2
	case l2 == nil:
		return l1
	case l1.Val < l2.Val:
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	default:
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

/**
 * 1:
 * --------
 * l1 : 1->3->4
 * l2 : 1->2->4
 *
 * default:
 * l2.Next = mergeTwoLists(l1, l2.Next)
 * l2.Val = 1
 * return l2
 *
 *
 * 2: from l2.Next = mergeTwoLists(l1, l2.Next)
 * --------
 * l1 : 1->3->4
 * l2 : 2->4
 *
 * case l1.Val < l2.Val:
 * l1.Next = mergeTwoLists(l1.Next, l2)
 * l1.Val = 1
 * return l1
 *
 *
 * 3: from l1.Next = mergeTwoLists(l1.Next, l2)
 * --------
 * l1 : 3->4
 * l2 : 2->4
 *
 * default:
 * l2.Next = mergeTwoLists(l1, l2.Next)
 * l2.Val = 2
 * return l2
 *
 *
 * 4: from l2.Next = mergeTwoLists(l1, l2.Next)
 * --------
 * l1 : 3->4
 * l2 : 4
 *
 * case l1.Val < l2.Val:
 * l1.Next = mergeTwoLists(l1.Next, l2)
 * l1.Val = 3
 * return l1
 *
 *
 * 5: from l1.Next = mergeTwoLists(l1.Next, l2)
 * --------
 * l1 : 4
 * l2 : 4
 *
 * default:
 * l2.Next = mergeTwoLists(l1, l2.Next)
 * l2.Val = 4
 * return l2
 *
 *
 * 6: from l2.Next = mergeTwoLists(l1, l2.Next)
 * --------
 * l1 : 4
 * l2 : nil
 *
 * case l2 == nil:
 * return l1
 *
 *
 * ===
 * 結束遞迴，開始所有 return
 *
 * From: Recursive 6
 * return l1 (4)
 * ↓
 * From: Recursive 5
 * l2.Next = mergeTwoLists(l1, l2.Next) = 4
 * l2.Val = 4
 * return l2
 * ↓
 * From: Recursive 4
 * l1.Next = mergeTwoLists(l1.Next, l2) = 4->4
 * l1.Val = 3
 * return l1
 * ↓
 * From: Recursive 3
 * l2.Next = mergeTwoLists(l1, l2.Next) = 3->4->4
 * l2.Val = 2
 * return l2
 * ↓
 * From: Recursive 2
 * l1.Next = mergeTwoLists(l1.Next, l2) = 2->3->4->4
 * l1.Val = 1
 * return l1
 * ↓
 * From: Recursive 1
 * l2.Next = mergeTwoLists(l1, l2.Next) = 1->2->3->4->4
 * l2.Val = 1
 * return l2
 * ↓
 * Final
 * return l2 = 1->1->2->3->4->4
**/
