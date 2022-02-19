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

	fmt.Println(mergeSort(arr, 0, 9))
}

func mergeSort(arr []int, start int, end int) []int {

	if start < 0 || end < 0 || start > end {
		return []int{}
	}
	if start == end {
		return arr[start : end+1]
	}

	mid := start + (end-start)>>1

	m1 := mergeSort(arr, start, mid)
	m2 := mergeSort(arr, mid+1, end)

	sum := len(m1) + len(m2)

	sm := make([]int, 0, sum)

	x := 0
	y := 0
	for x < len(m1) && y < len(m2) {

		if m1[x] <= m2[y] {
			sm = append(sm, m1[x])
			x++
		} else {
			sm = append(sm, m2[y])
			y++
		}

	}

	for x < len(m1) {
		sm = append(sm, m1[x])
		x++
	}
	for y < len(m2) {
		sm = append(sm, m2[y])
		y++
	}

	// fmt.Println(arr[start:end+1], m1, m2, sm)

	return sm

}
