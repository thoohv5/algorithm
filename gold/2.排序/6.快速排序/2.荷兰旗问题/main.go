package main

import "fmt"

/**
给定一个数组arr，和一个数num，请把小于等于num的数放在数 组的左边，大于num的
数放在数组的右边。要求额外空间复杂度O(1) ，时间复杂度O(N)
*/
func main() {

	arr := []int{3, 9, 10, 30, 20, 10, 5, 7, 8, 10}

	fmt.Println(marks(arr, 9))

}

func marks(arr []int, num int) []int {
	x := -1
	y := len(arr)
	for i := 0; i < y; i++ {
		if arr[i] < num {
			x++
			arr[i], arr[x] = arr[x], arr[i]
		} else if arr[i] > num {
			y--
			arr[i], arr[y] = arr[y], arr[i]
			i--
		}
	}
	return arr
}
