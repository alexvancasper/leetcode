package main

import (
	"fmt"
)

/*
Heap based on int type
*/
func siftUp(heap []int, index int) {
	if index == 1 {
		return
	}
	pIndex := index / 2
	if heap[pIndex-1] < heap[index-1] {
		heap[pIndex-1], heap[index-1] = heap[index-1], heap[pIndex-1]
		siftUp(heap, pIndex)
	}
}

func headAdd(heap []int, key int) []int {
	index := len(heap) + 1
	heap = append(heap, key)
	siftUp(heap, index)
	return heap
}

func siftDown(heap []int, index int) {
	left := 2*index + 1
	right := 2*index + 2
	if len(heap) <= left {
		return
	}
	indexLargest := left
	if right < len(heap) && heap[left] < heap[right] {
		indexLargest = right
	}
	if heap[index] < heap[indexLargest] {
		heap[index], heap[indexLargest] = heap[indexLargest], heap[index]
		siftDown(heap, indexLargest)
	}
}

/*
https://stackoverflow.com/questions/39993688/are-slices-passed-by-value
*/
func popMax(heap []int) (int, []int) {
	result := heap[0]
	heap[0] = heap[len(heap)-1]
	heap = heap[:len(heap)-1] // changing size of slice.
	siftDown(heap, 0)
	return result, heap
}

func main() {
	// Adding to the Heap, MaxHeap
	heap := []int{}
	heap = headAdd(heap, 5)
	heap = headAdd(heap, 8)
	heap = headAdd(heap, 3)
	fmt.Println(heap) // Output: [8 5 3]

	// Taking the Max priority element from the heap. In our case maximum value on the top.
	myHeap := []int{8, 5, 3}
	maxValue, myHeap := popMax(myHeap)
	fmt.Println(maxValue) // Output: 8
	fmt.Println(heap)     // Output: [5 3]
}
