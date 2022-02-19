package main

import "fmt"

/**
堆排序
时间复杂度：O(N*logN)
空间复杂度：O(1)
*/
func main() {
	arr := []int{3, 9, 10, 30, 20, 10, 5, 7, 8, 10}

	// fmt.Println(heapSort(arr))
	fmt.Println(heapSort2(arr))
}

func heapSort(arr []int) []int {

	l := len(arr)

	for j := 1; j < l; j++ {
		index := j
		for arr[(index-1)/2] < arr[index] {
			arr[(index-1)/2], arr[index] = arr[index], arr[(index-1)/2]
			index = (index - 1) / 2
		}
	}

	// fmt.Println(arr)

	for 0 < l {
		arr[0], arr[l-1] = arr[l-1], arr[0]
		l = l - 1

		// fmt.Println(arr, l)

		index := 0
		target := 0
		for index < l {
			left := 2*index + 1
			if left < l {
				target = left
				if right := left + 1; right < l {
					if arr[right] > arr[left] {
						target = right
					}
				}
			}
			if arr[target] <= arr[index] {
				break
			}
			arr[target], arr[index] = arr[index], arr[target]
			index = target
		}
		// fmt.Println(arr)
	}

	return arr

}

func heapSort2(arr []int) []int {

	l := len(arr)

	// for j := 1; j < l; j++ {
	// 	heapInsert(arr, j)
	// }
	for j := l; j >= 0; j-- {
		heapify(arr, j, l)
	}

	for l > 0 {
		arr[l-1], arr[0] = arr[0], arr[l-1]
		heapify(arr, 0, l-1)
		l--
	}

	return arr

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
