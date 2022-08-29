package main

import "fmt"

func main() {
	arr := []int{5, 7, 7, 8, 8, 10}
	// fmt.Println(binaryB(arr, 1))
	fmt.Println(searchRange(arr, 8))
}

/**
在一个有序数组中，查找>=某个数最左的位置
*/
func binaryB(nums []int, target int) int {

	first := -1
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)>>1
		if middle := nums[mid]; middle > target {
			end = mid - 1
		} else if middle < target {
			start = mid + 1
		} else {
			first = mid
			end = mid - 1
		}
	}
	return first
}

func searchRange(nums []int, target int) []int {

	first := -1
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)>>1
		if middle := nums[mid]; middle > target {
			end = mid - 1
		} else if middle < target {
			start = mid + 1
		} else {
			first = mid
			end = mid - 1
		}
	}

	second := -1
	start = 0
	end = len(nums) - 1
	for start <= end {
		mid := end - (end-start)>>1
		if middle := nums[mid]; middle > target {
			end = mid - 1
		} else if middle < target {
			start = mid + 1
		} else {
			second = mid
			start = mid + 1
		}
	}

	return []int{first, second}
}
