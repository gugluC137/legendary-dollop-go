package code

import . "github.com/gugluC137/legendary-dollop-go/models"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	l1, l2 := list1, list2
	ans := new(ListNode)
	cur := ans

	for l1 != nil && l2 != nil {
		val1, val2 := l1.Val, l2.Val
		if val1 <= val2 {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 == nil {
		cur.Next = l2
	}

	if l2 == nil {
		cur.Next = l1
	}

	return ans.Next
}
