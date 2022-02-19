package main

import "fmt"

/**
给定一个数组arr，和一个数num，请把小于等于num的数放在数 组的左边，大于num的
数放在数组的右边。要求额外空间复杂度O(1) ，时间复杂度O(N)
*/
func main() {

	arr := []int{3, 9, 10, 30, 20, 10, 5, 7, 8, 10}

	// fmt.Println(marksA(arr, 9))
	fmt.Println(marksB(arr, 9))

}

func marksA(arr []int, num int) []int {

	s := 0
	e := len(arr) - 1
	for i := 0; i < len(arr); {
		if num >= arr[i] {
			s++
			i++
			continue
		}
		if s >= e {
			break
		}
		arr[i], arr[e] = arr[e], arr[i]
		e--
	}
	return arr

}

func marksB(arr []int, num int) []int {
	s := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] <= num {
			s++
			arr[s], arr[i] = arr[i], arr[s]
		}
		// fmt.Println(i, arr, s)
	}
	return arr

}
