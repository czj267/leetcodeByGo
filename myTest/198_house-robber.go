package main

import (
	"fmt"
	"math"
)

func main() {
	//nums := []int{2, 7, 9, 3, 1}
	nums := []int{82, 217, 170, 215, 153, 55, 185, 55, 185, 232, 69, 131, 130, 102}
	res := rob(nums)
	fmt.Println(res)
}

var memoRob = make(map[int]int)

func rob(nums []int) int {
	memoRob = make(map[int]int)
	return tryRob(nums, 0)
}
func tryRob(nums []int, index int) int {
	if len(nums)-1 == index {
		return nums[index]
	}
	val, ok := memoRob[index]
	if ok {
		return val
	}
	res := 0
	for i := index; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i]+tryRob(nums, i+2))))
	}
	memoRob[index] = res
	return res
}
