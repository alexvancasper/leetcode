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

func CreateList(val int) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

func UnionLists(head1, head2 *ListNode) *ListNode {
	curr := head1
	for curr.Next != nil {
		curr = curr.Next
	}
	if curr.Next == nil {
		curr.Next = head2
	}
	return head1
}

/*
Works, but
O(n) time
O(n) additional memory
*/
func reverseLinkedList(head *ListNode) *ListNode {
	curr := head
	newList := CreateList(curr.Val)
	curr = curr.Next
	for curr.Next != nil {
		newList = AddToFront(newList, curr.Val)
		curr = curr.Next
	}
	newList = AddToFront(newList, curr.Val)
	return newList
}

/*
Works, but
O(n) time
O(1) additional memory
*/
func reverseLinkedListOptimized(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	head = prev
	return head
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

func UnionListNodesBack(head *ListNode, ptrs []*ListNode) *ListNode {
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	if curr.Next == nil {
		for i := len(ptrs) - 1; i >= 0; i-- {
			curr.Next = ptrs[i]
			curr = curr.Next
			ptrs[i] = nil
		}
		curr.Next = nil
	}
	return head
}

func UnionListNodesForward(head *ListNode, ptrs []*ListNode) *ListNode {
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	if curr.Next == nil {
		for i := 0; i < len(ptrs); i++ {
			if ptrs[i] != nil {
				curr.Next = ptrs[i]
				curr = curr.Next
				ptrs[i] = nil
			} else {
				break
			}

		}
		curr.Next = nil
	}
	return head
}

func UnionListNodes(head *ListNode, ptrs []*ListNode) *ListNode {
	if ptrs[len(ptrs)-1] == nil {
		return UnionListNodesForward(head, ptrs)
	} else {
		return UnionListNodesBack(head, ptrs)
	}
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
	// printNodes(head, k)
	fmt.Printf("Modified: ")
	root := swapNodes(head, k)

	// root := reverseLinkedListOptimized(head1)
	root.Print()

}
