package main

import "fmt"

func main() {
	arr := []int{3, 9, 10, 30, 20, 10, 5, 7, 8, 10, 70}

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
	e := end
	for i := start; i < e; i++ {
		if arr[i] < flag {
			s++
			arr[s], arr[i] = arr[i], arr[s]
		} else if arr[i] > flag {
			e--
			arr[e], arr[i] = arr[i], arr[e]
			i--
		}
	}

	if e < end {
		arr[e], arr[end-1] = arr[end-1], arr[e]
	}

	quicksort(arr, start, s+1)
	quicksort(arr, e+1, end)
}
