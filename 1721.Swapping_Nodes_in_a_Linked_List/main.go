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
За один проход по списку мы находим левый(l) и правый(r) элемент для свопа.
Правый элемент находится путем сохранения дистанции k от указателя curr.
т.е. если между r и curr дистанция больше чем k, то двигаем r на шаг ближе к curr.
*/
func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	var l, r, curr *ListNode
	var n, step int
	n = 1
	step = 1
	curr = head
	r = curr
	for curr != nil {
		if n == k {
			l = curr
		}
		if step > k {
			r = r.Next
			step--
		}
		step++
		n++
		curr = curr.Next
	}

	l.Val, r.Val = r.Val, l.Val
	return head
}

/*
Печатает k-ый элемент с конца списка.
*/
func printNodes(head *ListNode, k int) *ListNode {
	var prev, curr *ListNode
	curr = head
	prev = curr
	step := 1
	for curr != nil {
		if step > k {
			prev = prev.Next
			step--
		}
		step++
		curr = curr.Next
	}
	if curr == nil {
		fmt.Println("right element for swapping: ", prev.Val)
	}
	return head
}

func main() {
	head := &ListNode{Val: 1, Next: nil}
	head = AddToEnd(head, 2)
	head = AddToEnd(head, 3)
	head = AddToEnd(head, 4)
	// head = AddToEnd(head, 5)
	// head = AddToEnd(head, 6)
	// head1 = AddToEnd(head1, 7)
	fmt.Printf("LIST 1: ")
	head.Print()

	k := 1
	fmt.Printf("Modified: ")
	root := swapNodes(head, k)
	root.Print()

}
