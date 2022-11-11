package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	reminder := false
	// the first node in the result list will be skipped
	result := &ListNode{0, nil}
	lastNode := result

	for {
		if l1 == nil && l2 == nil && !reminder {
			break
		}

		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if reminder {
			sum += 1
		}

		reminder = false

		if sum > 9 {
			reminder = true
			sum -= 10
		}

		// save the result
		lastNode.Next = &ListNode{sum % 10, nil}
		lastNode = lastNode.Next
	}

	return result.Next

}
