package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (p *ListNode) Print() {
	for p.Next != nil {
		fmt.Printf("%d -> ", p.Val)
		p = p.Next
	}
	if p.Next == nil {
		fmt.Printf("%d -> nil\n", p.Val)
	}
}

func AddToEnd(head *ListNode, val int) *ListNode {
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	if curr.Next == nil {
		curr.Next = &ListNode{Val: val, Next: nil}
	}
	return head
}

func AddToFront(head *ListNode, val int) *ListNode {
	return &ListNode{Val: val, Next: head}

}

func AddNodeFront(head, node *ListNode) *ListNode {
	node.Next = head
	return node
}

func AddNodeBack(head, node *ListNode) *ListNode {
	curr := head
	for curr.Next != nil {
		curr = head.Next
	}
	if curr.Next == nil {
		curr.Next = node
	}
	return head
}

/*
Solution: Accepted
O(n) time
O(1) space
Хорошее решение, побило все 100% по памяти. и 60% по runtime
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	var pp, prev, curr *ListNode
	curr = head
	pp = curr
	prev = curr
	step := 1
	stepPP := 1
	for curr != nil {
		if step > n {
			prev = prev.Next
			step--
		}
		if stepPP > n+1 {
			pp = pp.Next
			stepPP--
		}
		step++
		stepPP++
		curr = curr.Next
	}
	if curr == nil {
		if pp == prev {
			head = prev.Next
		} else {
			pp.Next = prev.Next
		}

	}
	return head
}

func main() {
	head := &ListNode{Val: 1, Next: nil}
	head = AddToEnd(head, 2)
	// head = AddToEnd(head, 3)
	// head = AddToEnd(head, 4)
	// head = AddToEnd(head, 5)
	head = AddToEnd(head, 6)
	// head1 = AddToEnd(head1, 7)
	fmt.Printf("LIST 1: ")
	head.Print()

	k := 3
	// fmt.Printf("Modified: ")
	root := removeNthFromEnd(head, k)
	root.Print()

}
