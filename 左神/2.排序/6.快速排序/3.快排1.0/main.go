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

	flag := arr[end]

	s := start - 1
	for i := start; i < end; i++ {
		if arr[i] <= flag {
			s++
			arr[s], arr[i] = arr[i], arr[s]
		}
	}

	arr[s+1], arr[end] = arr[end], arr[s+1]

	fmt.Println(arr[start:end], arr[start:s], arr[s+1], prints(arr, s+2, end))

	return append(append(quicksort(arr, start, s), arr[s+1]), quicksort(arr, s+2, end)...)

}

func prints(arr []int, start, end int) []int {
	if end < start {
		return []int{}
	}
	return arr[start:end]
}
