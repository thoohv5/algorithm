package main

import "fmt"

/**
归并排序

1) 整体就是一个简单递归，左边排好序、右边排好序、让其整体有序
2) 让其整体有序的过程里用了排外序方法
3) 利用master公式来求解时间复杂度
4) 归并排序的实质

时间复杂度0 (N*logN) ，额外空间复杂度0 (N)

*/
func main() {
	arr := []int{3, 9, 10, 30, 20, 10, 5, 7, 8, 10}
	mergeSort(arr, 0, len(arr))
	fmt.Println(arr)
}

func mergeSort(arr []int, start int, end int) {
	if start < 0 || end < 0 || start >= end {
		return
	}
	if start+1 == end {
		return
	}

	mid := start + (end-start)>>1

	mergeSort(arr, start, mid)
	mergeSort(arr, mid, end)

	x := start
	y := mid
	help := make([]int, end-start)
	i := 0
	for x < mid && y < end {
		if arr[x] < arr[y] {
			help[i] = arr[x]
			x++
		} else {
			help[i] = arr[y]
			y++
		}
		i++
	}

	for x < mid {
		help[i] = arr[x]
		x++
		i++
	}
	for y < end {
		help[i] = arr[y]
		y++
		i++
	}

	copy(arr[start:], help)
}
