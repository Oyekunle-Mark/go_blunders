package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// first pass

	//nodeStore := make(map[*ListNode]bool) // hold list node from lists
	//
	//current := headA // head of list headA
	//
	//// for all node, store
	//for current != nil {
	//	nodeStore[current] = true
	//	current = current.Next
	//}
	//
	//current = headB // head of list headB
	//
	//// for all node in headB, the first node found in store is intersection point
	//for current != nil {
	//	if _, ok := nodeStore[current]; ok {
	//		return current
	//	}
	//
	//	current = current.Next
	//}
	//
	//return nil // not found

	// second pass

	currentA := headA // point currentA to head of headA list
	currentB := headB // point currentB to head of headB list

	// loop while both list node are not equivalent
	for currentA != currentB {
		// at end of currentA, start at the beginning of headB list
		if currentA == nil {
			currentA = headB
		} else { // otherwise, advance pointer
			currentA = currentA.Next
		}

		// at end of currentB, start at the beginning of headA list
		if currentB == nil {
			currentB = headA
		} else { // otherwise advance pointer
			currentB = currentB.Next
		}
	}

	return currentA // returning meeting point
}
