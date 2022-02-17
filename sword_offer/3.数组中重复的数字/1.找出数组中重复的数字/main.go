package main

import (
	"errors"
	"fmt"
)

/**
题目一: 找出数组中重复的数字。
	在一个长度为的数组里的所有数字都在 0一n-1 的范围内。数组中某
	些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了
	几次。请找出数组中任意一个重复的数字。  例如，如果输入长度为 7 的数
	组{2, 3, 1, 0, 2, 5,3}，那么对应的输出是重复的数字 2 或者 3。
*/

func main() {
	arr := []int{2, 3, 1, 0, 2, 5, 3}

	// fmt.Println(findRepeatNumA(arr))
	fmt.Println(findRepeatNumB(arr))

}

func findRepeatNumA(arr []int) (int, error) {
	exist := make(map[int]bool)
	for _, val := range arr {
		if _, ok := exist[val]; !ok {
			exist[val] = true
		} else {
			return val, nil
		}
	}
	return 0, errors.New("not exists")
}

// 2, 3, 1, 0, 2, 5, 3
// 1, 3, 2, 0, 2, 5, 3
// 3, 1, 2, 0, 2, 5, 3
// 0, 1, 2, 3, 2, 5, 3
// 0, 1, 2, 3, 2, 5, 3
// 0, 1, 2, 3, 2, 5, 3  ---
func findRepeatNumB(arr []int) (int, error) {
	for i := 0; i < len(arr); i++ {
		if i == arr[0] {
			continue
		}
		if arr[0] < i {
			return arr[i], nil
		}
		arr[0], arr[arr[0]] = arr[arr[0]], arr[0]
	}
	return 0, errors.New("not exists")
}
