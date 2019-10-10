package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 7, 9, 12, 20}
	fmt.Println(search(nums, 0))
}

/**
二分查找
*/
func search(nums []int, val int) bool {
	l, r := 0, len(nums)-1
	if l == r {
		return false
	}

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == val {
			return true
		}
		if nums[mid] < val {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
