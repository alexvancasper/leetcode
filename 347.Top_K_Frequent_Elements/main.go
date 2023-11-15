package main

import (
	"fmt"
)

type Elem struct {
	elem int
	freq int
}

/*
Heap based on Elem type
*/

func mySiftUp(heap []Elem, index int) {
	if index == 1 {
		return
	}
	pIndex := index / 2
	if heap[pIndex-1].freq < heap[index-1].freq {
		heap[pIndex-1], heap[index-1] = heap[index-1], heap[pIndex-1]
		mySiftUp(heap, pIndex)
	}
}

func myHeadAdd(heap []Elem, key Elem) []Elem {
	index := len(heap) + 1
	heap = append(heap, key)
	mySiftUp(heap, index)
	return heap
}

func mySiftDown(heap []Elem, index int) {
	left := 2*index + 1
	right := 2*index + 2
	if len(heap) <= left {
		return
	}
	indexLargest := left
	if right < len(heap) && heap[left].freq < heap[right].freq {
		indexLargest = right
	}
	if heap[index].freq < heap[indexLargest].freq {
		heap[index], heap[indexLargest] = heap[indexLargest], heap[index]
		mySiftDown(heap, indexLargest)
	}
}

func myPopMax(heap []Elem) (Elem, []Elem) {
	result := heap[0]
	heap[0] = heap[len(heap)-1]
	heap = heap[:len(heap)-1]
	mySiftDown(heap, 0)
	return result, heap
}

/*
topKFrequentMap -

	pMap: {
		"element":frequency
	}

Then add these elements into the Max Heap based on frequency value
Take from the heap K max elements
*/
func topKFrequent(nums []int, k int) []int {
	pMap := make(map[int]int)
	for _, num := range nums {
		if _, ok := pMap[num]; ok {
			pMap[num]++
		} else {
			pMap[num] = 1
		}
	}
	heap := []Elem{}
	for key, value := range pMap {
		elem := Elem{elem: key, freq: value}
		// delete(pMap, key)
		heap = myHeadAdd(heap, elem)

	}

	answer := []int{}
	var value Elem
	for k > 0 {
		value, heap = myPopMax(heap)
		answer = append(answer, value.elem)
		k--
	}
	return answer
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3} //[]int{1,2}
	k := 2
	res := topKFrequent(nums, k)
	fmt.Printf("%v\n", res)

}
