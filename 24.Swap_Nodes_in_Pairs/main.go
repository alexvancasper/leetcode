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
Runtime: 1ms
Memory: 2MB
I created this solution. I like it!
O(n) Time
O(1) Space
Смысл в том, что мы идем по списку с помощью трех указателей p (prev),c (current), n (next).
Изначально head меняется на head.next, а затем меняются ссылки
c->p, p->n.next
перемещаем указатели на один шаг вперед и снова
c->p, p->n.next
Если где-то встречаем nil, то просто выходим с занулением либо с присвоением последнего элемента. Зависит от четности элементов в списке
*/
func swapPairs(head *ListNode) *ListNode {
	var p, c, n *ListNode
	if head == nil {
		return head
	}
	p = head
	if p.Next != nil {
		c = p.Next
		n = c.Next
		head = head.Next
	}

	for c != nil {

		c.Next = p
		if n == nil {
			// четный
			p.Next = nil
			break
		}

		if n.Next == nil {
			// нечетный
			p.Next = n
			break
		} else {
			p.Next = n.Next
		}

		p = n
		c = p.Next
		n = c.Next
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

	fmt.Printf("Modified: ")
	root := swapPairs(head)
	root.Print()
}
