package main

import (
	"fmt"
)

/**
题目二: 不修改数组找出重复的数字

在一个长度为 a+l 的数组里的所有数字都在 1一疡的范围内，所以数组
中至少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能
修改输入的数组。例如，如果输入长度为 8 的数组{2, 3, 5,4, 3, 2, 6,7}，
那么对应的输出是  复 的数字 2 政者 3。

*/

func main() {
	arr := []int{2, 3, 5, 4, 3, 2, 6, 7}

	fmt.Println(findRepeatNumA(arr, 0))

}

// 8->4
// {2,3,4,3,2} {5,6,7}

// {1,1}
//
func findRepeatNumA(arr []int, start int) (int, error) {
	return 0, nil
}
