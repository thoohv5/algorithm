package main

import "fmt"

func main() {
	arr := []int{3, 9, 10, 30, 20, 10, 5, 7, 8, 10}

	quicksort(arr, 0, len(arr))
	fmt.Println(arr)
}

func quicksort(arr []int, start int, end int) {

	if start < 0 || end < 0 || end <= start {
		return
	}

	if start+1 == end {
		return
	}

	flag := arr[end-1]

	s := start - 1
	for i := start; i < end-1; i++ {
		if arr[i] <= flag {
			s++
			arr[s], arr[i] = arr[i], arr[s]
		}
	}

	arr[s+1], arr[end-1] = arr[end-1], arr[s+1]

	quicksort(arr, start, s+1)
	quicksort(arr, s+2, end)
}
