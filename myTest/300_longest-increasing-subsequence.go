package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//nums:=[]int{1,3,6,7,9,4,10,5,6}
	res := lengthOfLIS(nums)
	fmt.Println(res)
}

func lengthOfLIS(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	memo := make([]int, l, l)
	for i := 0; i < l; i++ {
		memo[i] = 1
	}
	memo[0] = 1
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if memo[i] < memo[j]+1 {
					memo[i] = memo[j] + 1
				}
			}
		}
	}
	res := 1
	for i := 0; i < l; i++ {
		if res < memo[i] {
			res = memo[i]
		}
	}
	return res
}
