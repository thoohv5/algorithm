package main

import "fmt"

func main() {
	fmt.Println(fib(10), fib2(10))
}

// 递归
func fib(n int) int {

	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func fib2(n int) int {

	if n < 0 {
		return 0
	}

	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
