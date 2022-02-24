package main

import "fmt"

func main() {
	arr := []int{3, 24, 123, 12, 1, 6, 7, 0, 23, 45, 67}
	fmt.Println(sort(arr))
}

/**
选择排序
选择最小的和第一个交换
*/
func sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		chooseIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[chooseIndex] > arr[j] {
				chooseIndex = j
			}
		}
		arr[i], arr[chooseIndex] = arr[chooseIndex], arr[i]
	}

	return arr
}
