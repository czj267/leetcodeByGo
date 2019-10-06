package main

import "fmt"

func main() {
	nums := []int{1, 5, 11, 5, 6}
	res := canPartition(nums)
	fmt.Println(res)
}

type PartitionEqualSubset struct {
	memo map[int]map[int]int
}

func canPartition(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 != 0 {
		return false
	}
	p := PartitionEqualSubset{memo: make(map[int]map[int]int)}
	res := p.tryPartition(nums, len(nums)-1, sum/2)
	return res
}

func (p PartitionEqualSubset) tryPartition(nums []int, index, sum int) bool {
	if sum == 0 {
		return true
	}
	if sum < 0 || index < 0 {
		return false
	}
	val, ok := p.memo[index][sum]
	if ok && val != -1 {
		return val == 1
	}
	res := p.tryPartition(nums, index-1, sum) || p.tryPartition(nums, index-1, sum-nums[index])
	var _res int
	if res {
		_res = 1
	} else {
		_res = 0
	}
	if p.memo[index] == nil {
		//make多维数组，只能初始化外城的，内层的数组都是nil
		p.memo[index] = make(map[int]int)
	}
	p.memo[index][sum] = _res
	return res
}
