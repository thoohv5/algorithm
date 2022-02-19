package main

import "fmt"

/**
堆排序
1. 完全二叉树
2. 任何一颗子数，根节点都是极值
时间复杂度：O(logN)
*/
func main() {
	heap := []int{6, 3, 5}
	fmt.Println(heapInsertValue(heap, 7))
	heap, val := heapifyValue(heap)
	fmt.Println(heap, val)

	heap1 := []int{8, 7, 6, 10}
	heapInsert(heap1, 3)
	fmt.Println(heap1)

	heap1[len(heap1)-1], heap1[0] = heap1[0], heap1[len(heap1)-1]
	heapify(heap1, 0, 3)
	fmt.Println(heap1)
}

// 插入
func heapInsertValue(heap []int, num int) []int {
	heap = append(heap, num)
	index := len(heap) - 1
	for heap[(index-1)/2] < heap[index] {
		heap[(index-1)/2], heap[index] = heap[index], heap[(index-1)/2]
		index = (index - 1) / 2
	}
	return heap
}

// 插入
func heapInsert(heap []int, index int) {
	for heap[(index-1)/2] < heap[index] {
		heap[(index-1)/2], heap[index] = heap[index], heap[(index-1)/2]
		index = (index - 1) / 2
	}
}

// 取极值
func heapifyValue(heap []int) ([]int, int) {
	if len(heap) == 0 {
		return nil, 0
	}
	heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]

	index := 0

	target := 0
	for index < len(heap)-2 {
		left := 2*index + 1
		if left < len(heap)-1 {
			target = left
			if right := left + 1; right < len(heap)-1 {
				if heap[right] > heap[left] {
					target = right
				}
			}
		}
		if heap[target] <= heap[index] {
			break
		}
		heap[target], heap[index] = heap[index], heap[target]
		index = target
	}

	val := heap[len(heap)-1]
	return heap[0 : len(heap)-1], val
}

// 取极值
func heapify(heap []int, index, heapSize int) {
	target := 0
	left := 2*index + 1
	for left < heapSize {
		target = left
		if right := left + 1; right < heapSize {
			if heap[right] > heap[left] {
				target = right
			}
		}
		if heap[target] <= heap[index] {
			break
		}
		heap[target], heap[index] = heap[index], heap[target]
		index = target
		left = 2*index + 1
	}
}
