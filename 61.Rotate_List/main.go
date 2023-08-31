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
O(n) time, тут реальная сложность O((k%n)n), где k количество rotate, n - количество элементов.
O(1) space
Решение так себе, можно написать лучше.
*/
func rotateRight(head *ListNode, k int) *ListNode {
	var first, prev, curr *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	c := head
	n := 0
	for c != nil {
		n++
		c = c.Next
	}
	rotates := k % n
	for rotates > 0 {
		prev = &ListNode{}
		first = head
		curr = head
		prev.Next = curr
		for curr != nil {
			if curr.Next == nil {
				head = curr
				head.Next = first
				prev.Next = nil
				break
			}
			prev = prev.Next
			curr = curr.Next
		}
		rotates--
	}
	return head
}

func main() {
	head := &ListNode{Val: 1, Next: nil}
	head = AddToEnd(head, 2)
	head = AddToEnd(head, 3)
	// head = AddToEnd(head, 4)
	// head = AddToEnd(head, 5)
	// head = AddToEnd(head, 6)
	// head1 = AddToEnd(head1, 7)
	fmt.Printf("LIST 1: ")
	head.Print()

	k := 2000000000
	fmt.Printf("Modified: ")
	root := rotateRight(head, k)

	root.Print()

}
