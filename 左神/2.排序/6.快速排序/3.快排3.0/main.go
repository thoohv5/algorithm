package main

import "fmt"

func main() {
	arr := []int{3, 9, 10, 30, 20, 10, 5, 7, 8, 10}

	fmt.Println(quicksort(arr, 0, 9))
}

func quicksort(arr []int, start int, end int) []int {

	if start < 0 || end < 0 || end < start {
		return []int{}
	}

	if start == end {
		return arr[start : end+1]
	}

	// 随机参考值
	flag := arr[end]

	sx := start - 1
	sy := end + 1
	for i := start; i < end && i < sy; {
		if arr[i] < flag {
			sx++
			arr[sx], arr[i] = arr[i], arr[sx]
			i++
		} else if arr[i] == flag {
			i++
		} else {
			sy--
			arr[sy], arr[i] = arr[i], arr[sy]
		}
	}

	arr[sy], arr[end] = arr[end], arr[sy]

	fmt.Println(arr[start:end+1], arr[start:sx], arr[sx+1:sy+1], prints(arr, sy+1, end))

	return append(append(quicksort(arr, start, sx), arr[sx+1:sy+1]...), quicksort(arr, sy+1, end)...)

}

func prints(arr []int, start, end int) []int {
	if end < start {
		return []int{}
	}
	return arr[start:end]
}
