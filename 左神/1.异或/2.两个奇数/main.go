package main

import "fmt"

/**
n个数，2一个数奇数次，其他偶数次，找奇数次的数
解析：2个奇数，所以比不想等（不为0），必然有一个数有一位为1，另外一个数此位为0，此问题就转换为问题1
*/
func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	result := 0
	for _, val := range arr {
		result ^= val // 8^9
	}

	// result 最右侧的1
	flag := result & (^result + 1)

	temp := 0
	for _, val := range arr {
		if flag&val == 0 {
			temp ^= val
		}
	}

	fmt.Println(temp, result^temp)

}
