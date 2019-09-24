package main

import "fmt"

func main() {
	nums := []int{3, 6, 2, 8, 5, 1, 9}
	//nums := []int{1,2,3,6,5}
	//2，6,3,8,5,1，9
	//i := 3
	//nums[i-1], nums[0] = nums[0], nums[i-1]
	heapSort(nums)
	fmt.Printf("%v", nums)
}

func heapSort(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, len(nums))
	}
	for j := len(nums) - 1; j > 0; j-- {
		nums[j], nums[0] = nums[0], nums[j]
		adjustHeap(nums, 0, j)
	}
}

func adjustHeap(nums []int, i, l int) {
	tmp := nums[i]
	for k := 2*i + 1; k < l; k = 2*k + 1 {
		if k+1 < l && nums[k] < nums[k+1] {
			k++
		}
		if nums[k] > tmp {
			nums[i] = nums[k]
			i = k
		} else {
			break
		}
	}
	nums[i] = tmp
}
