package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	res := permuteUnique(nums)
	fmt.Printf("%v", res)
}

/**
回溯法
*/

/**
包含重复元素的全排列
*/
func permuteUnique(nums []int) [][]int {
	var res [][]int
	perUi(nums, 0, &res)
	return res
}

func perUi(nums []int, index int, res *[][]int) {
	if index == len(nums) {
		_per1 := make([]int, len(nums))
		copy(_per1, nums)
		*res = append(*res, _per1)
		return
	}
	//注意这个哈希是放在函数内部的
	var used = make(map[int]bool)
	for i := index; i < len(nums); i++ {
		if _, ok := used[nums[i]]; ok {
			continue
		}
		nums[i], nums[index] = nums[index], nums[i]
		perUi(nums, index+1, res)
		nums[i], nums[index] = nums[index], nums[i]
		used[nums[i]] = true
	}
}
