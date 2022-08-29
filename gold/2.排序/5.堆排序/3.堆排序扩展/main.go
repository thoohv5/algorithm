package main

import (
	"container/heap"
	"fmt"
)

/**
topK算法
浮动窗口

已知一个几乎有序的数组，几乎有序是指，如果把数组排好顺序的话，每个元
素移动的距离可以不超过k，并且k相对于数组来说比较小。请选择一个合适的
排序算法针对这个数据进行排序。
*/
func main() {

	arr := []int{1, 2, 8, 7} // K = 2

	h := &IntHeap{}
	heap.Init(h)

	k := 2
	for i := 0; i < k; i++ {
		heap.Push(h, arr[i])
	}

	for i := k; i < len(arr); i++ {
		fmt.Println(heap.Pop(h).(int))
		heap.Push(h, arr[i])
	}

	for h.Len() > 0 {
		fmt.Println(heap.Pop(h).(int))
	}

}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
