package main

import "fmt"

func main() {
	res := climbStairs(10)
	fmt.Println(res)
}

var memo = make(map[int]int)

func climbStairs(n int) int {
	if n < 0 {
		return 0
	}
	//LeetCode环境下需要清空
	memo = make(map[int]int)
	return climS(n)
}

func climS(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	_, ok := memo[n]
	if !ok {
		memo[n] = climS(n-1) + climS(n-2)
	}
	return memo[n]
}
