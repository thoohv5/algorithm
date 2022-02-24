package main

/**
一只青蛙一次可以跳上 1 级台阶，也可以跳上 2 级台阶. 求青蛙跳上
上一个 n 级的台阶总共有多少种跳法。
*/
func main() {

}

func numWays(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 || n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		dp[i] %= 1000000007
	}

	return dp[n]
}
