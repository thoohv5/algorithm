package main

import "fmt"

func main() {
	arr := []int{1, 4, 7, 8, 9, 9, 9, 14, 25, 38}
	fmt.Println(binaryB(arr, 8))
}

/**
在一个有序数组中，查找>=某个数最左的位置
*/
func binaryB(arr []int, val int) int {

	first := 0
	end := len(arr)
	for first < end {
		middle := (end - first) / 2
		if middle == 0 {
			return end
		}
		if m := arr[first+middle]; m < val {
			first = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1
}
