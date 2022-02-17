package main

import "fmt"

func main() {
	arr := []int{3, 24, 123, 12, 1, 6, 7, 0, 23, 45, 67}
	fmt.Println(sort(arr))
}

/**
插入排序
两两比较，前>后面，交换
*/
func sort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return arr

}
