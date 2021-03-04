package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeStore := make(map[*ListNode]bool) // hold list node from lists

	current := headA // head of list headA

	// for all node, store
	for current != nil {
		nodeStore[current] = true
		current = current.Next
	}

	current = headB // head of list headB

	// for all node in headB, the first node found in store is intersection point
	for current != nil {
		if _, ok := nodeStore[current]; ok {
			return current
		}

		current = current.Next
	}

	return nil // not found
}
