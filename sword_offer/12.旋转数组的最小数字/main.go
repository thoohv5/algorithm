package main

import "math"

/**
题目: 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为
数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小
元素。例如，数组{3, 4, 5, 1, 2}为{1, 2, 3, 4, 5}的一个旋转，该数组的最小
值为1。
*/
func main() {

}

func findMin(nums []int) int {
	min := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if nums[i] <= min {
			min = nums[i]
		}
	}
	return min
}
