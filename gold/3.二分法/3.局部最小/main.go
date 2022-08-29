package main

import "fmt"

func main() {
	arr := []int{10, 30, 80, 8, 9, 9, 9, 14, 25, 38}
	fmt.Println(binaryB(arr))
}

/**
局部最小值
*/
func binaryB(arr []int) int {

	first := 0
	end := len(arr)
	if end == 0 {
		return -1
	}
	for first < end {
		if end-first == 1 {
			return arr[0]
		}
		middle := first + (end-first)/2
		if arr[middle] > arr[middle+1] {
			first = middle
		} else {
			end = middle - 1
		}
	}
	return -1
}
