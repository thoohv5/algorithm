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
	s := -1
	ss := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] <= num {
			s++
			if i > s {
				arr[s], arr[i] = arr[i], arr[s]
			}
			if arr[s] < num {
				ss++
				if s > ss {
					arr[ss], arr[s] = arr[s], arr[ss]
				}
			}
		}
	}
	return arr
}

func marksB(arr []int, num int) []int {
	x := -1
	y := len(arr)
	for i := 0; i < len(arr) && y > i; {
		if arr[i] < num {
			x++
			arr[i], arr[x] = arr[x], arr[i]
			i++
		} else if arr[i] == num {
			i++
		} else {
			y--
			arr[i], arr[y] = arr[y], arr[i]
		}
		// fmt.Println(i, arr, x, y)
	}
	return arr
}
