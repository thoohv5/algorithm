package main

import "fmt"

/**
n个数，1一个数奇数次，其他偶数次，找奇数次的数
*/
func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	temp := 0
	for _, val := range arr {
		temp ^= val
	}

	fmt.Println(temp)

}
