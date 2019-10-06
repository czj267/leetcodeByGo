package main

import "fmt"

func main() {
	nums := []int{5, 4, 6, 2}
	res := permute(nums)
	fmt.Printf("%v", res)
}

/**
回溯法
*/
//不重复数字数组全排列
func permute(nums []int) [][]int {
	var res [][]int
	//per(nums, []int{}, &res)
	per2(nums, 0, &res)
	return res
}

func per2(nums []int, index int, res *[][]int) {
	if index == len(nums) {
		//注意这里要copy，否则已经append到res的结果也会被修改
		_per1 := make([]int, len(nums))
		copy(_per1, nums)
		*res = append(*res, _per1)
		return
	}

	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		//注意和 per 对比
		per2(nums, index+1, res)
		nums[i], nums[index] = nums[index], nums[i]
	}
}

func per(nums []int, per1 []int, res *[][]int) {
	if len(nums) == 0 {
		//注意这里要copy，否则已经append到res的结果也会被修改
		_per1 := make([]int, len(per1))
		copy(_per1, per1)
		*res = append(*res, _per1)
		return
	}
	for i := 0; i < len(nums); i++ {
		per1 = append(per1, nums[i])
		//这里要注意slice的问题
		_nums := make([]int, i)
		copy(_nums, nums[0:i])
		_nums = append(_nums, nums[i+1:]...)
		per(_nums, per1, res)
		per1 = per1[0 : len(per1)-1]
	}
}

func permute1(nums []int) [][]int {
	return perm(nums, 0)
}

func perm(nums []int, k int) (res [][]int) {
	if k == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		res = append(res, tmp)
		return
	}
	for i := k; i < len(nums); i++ {
		nums[i], nums[k] = nums[k], nums[i]
		res = append(perm(nums, k+1), res...)
		nums[i], nums[k] = nums[k], nums[i]
	}
	return res
}
