package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var n, total int
	fmt.Scanln(&total)

	data := new(MinHeap)
	heap.Init(data)
	for i := 0; i < total; i++ {
		fmt.Scanln(&n)
		if n == 0 { // Pop
			if data.Len() == 0 {
				fmt.Println(0)
			} else {
				fmt.Println(heap.Pop(data).(int))
			}
		} else {
			heap.Push(data, n)
		}
	}
}
