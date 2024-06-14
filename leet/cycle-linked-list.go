package leet

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	visitedMap := make(map[*ListNode]bool)
	current := head
	for current != nil {
		if _, ok := visitedMap[current]; ok {
			return true
		}
		visitedMap[current] = true
		current = current.Next
	}

	return false
}
