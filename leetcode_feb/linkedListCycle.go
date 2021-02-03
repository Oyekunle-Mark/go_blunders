package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	nodeStore := map[*ListNode]bool{}

	for head != nil {
		if nodeStore[head] {
			return true
		}

		nodeStore[head] = true
		head = head.Next
	}

	return false
}
