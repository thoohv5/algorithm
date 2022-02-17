package main

import "fmt"

/**
题目: 在一个二维数组中，每一行都按照从在到右递增的顺序排序，
每一列都按照从上到下递增的顺序排序， 请完成一个函数，输入这样的一
个二维数组和一个整数，判断数组中是否含有该整数。
*/
func main() {
	arr := [][]int{
		{3, 6, 9},
		{6, 7, 10},
		{9, 10, 12},
	}
	fmt.Println(findTwoDimensionNum(arr, 6))
}

func findTwoDimensionNum(arr [][]int, num int) bool {

	w := len(arr)
	h := len(arr[0])

	flag := 0
	for j := 0; j < w; j++ {
		if arr[0][j] == num {
			return true
		} else if arr[0][j] > num {
			flag = j - 1
		}
	}

	if flag == -1 {
		return false
	}

	for i := 1; i < h; i++ {
		if arr[i][flag] == num {
			return true
		}
	}

	return false

}
