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
Runtime: 5ms
Memory: 3.7MB
not so good solution, but I created it. :)
O(n) time
O(k) additional memory- по честному это O(n)
Смысл в том, что мы собираем K нод в массив, а затем добавляем его в новый лист.
Если массив забит не полностью, значит читаем массив сначала и добавляем в конец нового листа
Если массив забит полностью, то читаем массив с конца и добавляем ноды в конец листа.
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	ptrs := make([]*ListNode, k)
	newList := &ListNode{}
	curr := head
	i := 0
	for curr != nil {
		if i < k {
			ptrs[i] = curr
			i++
		} else {
			newList = UnionListNodes(newList, ptrs) // UnionListNodesBack
			ptrs[0] = curr
			i = 1
		}

		curr = curr.Next
	}

	newList = UnionListNodes(newList, ptrs) // UnionListNodesForward
	return newList.Next
}

func main() {
	head1 := &ListNode{Val: 1, Next: nil}
	head1 = AddToEnd(head1, 2)
	head1 = AddToEnd(head1, 3)
	// head1 = AddToEnd(head1, 4)
	// head1 = AddToEnd(head1, 5)
	// head1 = AddToEnd(head1, 6)
	// head1 = AddToEnd(head1, 7)
	fmt.Printf("LIST 1: ")
	head1.Print()

	fmt.Printf("Modified: ")
	root := reverseKGroup(head1, 3)
	root.Print()
}
