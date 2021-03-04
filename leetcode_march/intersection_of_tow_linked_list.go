package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeStore := make(map[*ListNode]bool)

	current := headA

	for current != nil {
		nodeStore[current] = true
		current = current.Next
	}

	current = headB

	for current != nil {
		if _, ok := nodeStore[current]; ok {
			return current
		}

		current = current.Next
	}

	return nil
}
