package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println("Output:", findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	m1 := l1 / 2
	m2 := l2 / 2
	if nums1[m1] == nums2[m2] {
		return float64(nums1[m1])
	}

	if nums1[m1] > nums2[m2] {

	} else {

	}
	return 0.1
}

//执行用时 :16 ms, 在所有Go提交中击败了97.95%的用户
//内存消耗 :6 MB, 在所有Go提交中击败了27.06%的用户
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1)%2 == 0 {
		return float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2]) / 2
	}
	return float64(nums1[len(nums1)/2])
}
