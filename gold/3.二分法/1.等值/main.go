package main

import "fmt"

func main() {
	arr := []int{1, 4, 7, 8, 9, 14, 25, 38}
	fmt.Println(binaryA(arr, 25), binaryA(arr, 3))
	fmt.Println(binaryB(arr, 25), binaryB(arr, 3))
}

/**
在一个有序数组中，查找某个数是否存在
*/
func binaryA(arr []int, val int) bool {
	if len(arr) == 0 {
		return false
	}
	length := len(arr) / 2
	if cur := arr[length]; val == cur {
		return true
	} else if val > cur {
		return binaryA(arr[length+1:], val)
	} else {
		return binaryA(arr[:length], val)
	}

}

/**
在一个有序数组中，查找某个数是否存在
*/
func binaryB(arr []int, val int) bool {

	first := 0
	end := len(arr)
	for first < end {
		middle := (end - first) / 2
		if m := arr[first+middle]; m == val {
			return true
		} else if m < val {
			first = middle + 1
		} else {
			end = middle - 1
		}
	}
	return false
}
