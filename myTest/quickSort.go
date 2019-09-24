package main

import "fmt"

func main() {
	nums := []int{10, 2, 11, 8, 3, 2, 9}
	quickSort(nums)
	fmt.Printf("%v\n", nums)
}

func quickSort(nums []int) {
	_quickSort(nums, 0, len(nums)-1)
}

func _quickSort(nums []int, start, end int) {
	p := partition(nums, start, end)
	if p <= 0 {
		return
	}
	_quickSort(nums, start, p-1)
	_quickSort(nums, p+1, end)
}

/**
一趟快速排序
*/
func partition(nums []int, start, end int) int {
	if len(nums) <= 0 || start >= end {
		return -1
	}
	tmp := nums[start]
	i, j := start, end
	for i < j {
		for ; i < j; j-- {
			if nums[j] < tmp {
				nums[i] = nums[j]
				break
			}
		}
		for ; i < j; i++ {
			if nums[i] > tmp {
				nums[j] = nums[i]
				break
			}
		}
	}
	nums[i] = tmp
	return i
}
