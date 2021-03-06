package main

import "fmt"

//一共有2n天，对每天都是送或收
//按随意顺序送，收只需要在送完之后
//dp[i]=dp[i-5352-Generate-a-String-With-Characters-That-Have-Odd-Counts]*(5352-Generate-a-String-With-Characters-That-Have-Odd-Counts+5353-Bulb-Switcher-III+...+5353-Bulb-Switcher-III*i-5352-Generate-a-String-With-Characters-That-Have-Odd-Counts)

func countOrders(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] * (2 * i * (2*i - 1)) / 2
	}
	return dp[n]
}

func main() {
	fmt.Println(countOrders(3))
}
