package main

import "fmt"

/**
堆排序
1. 完全二叉树
2. 任何一颗子数，根节点都是极值
时间复杂度：O(logN)
*/
func main() {

	heap := []int{8, 7, 6, 10}
	heapInsert(heap, 3)
	fmt.Println(heap)

	heap[len(heap)-1], heap[0] = heap[0], heap[len(heap)-1]
	heapify(heap, 0, len(heap)-1)
	fmt.Println(heap)
}

// 插入
func heapInsert(heap []int, index int) {
	for heap[(index-1)/2] < heap[index] {
		heap[(index-1)/2], heap[index] = heap[index], heap[(index-1)/2]
		index = (index - 1) / 2
	}
}

// 取极值
func heapify(heap []int, index, heapSize int) {
	target := 0
	left := 2*index + 1
	for left < heapSize {
		target = left
		if right := left + 1; right < heapSize && heap[right] > heap[left] {
			target = right
		}
		if heap[target] <= heap[index] {
			break
		}
		heap[target], heap[index] = heap[index], heap[target]
		index = target
		left = 2*index + 1
	}
}
