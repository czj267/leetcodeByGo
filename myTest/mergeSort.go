package main

import "fmt"

func main() {
	nums := []int{3, 2, 3, 4, 9, 6, 7, 1, 3}
	mergeSort(nums)
	fmt.Printf("%v", nums)
}

func mergeSort(nums []int) {
	l, r := 0, len(nums)-1
	tmpNums := make([]int, len(nums)) //和nums等长的数据可以避免重复申请
	sort(nums, l, r, tmpNums)
}

func sort(nums []int, l, r int, tmpNums []int) {
	if l < r {
		m := (l + r) / 2
		sort(nums, l, m, tmpNums)
		sort(nums, m+1, r, tmpNums)
		merge(nums, l, m, r, tmpNums)
	}
}

/**
相当于是合并两个有序的数组到一个新的数组中
*/
func merge(nums []int, l, m, r int, tmpNums []int) {
	i, j, k := l, m+1, l

	for i <= m && j <= r {
		if nums[i] < nums[j] {
			tmpNums[k] = nums[i]
			i++
		} else {
			tmpNums[k] = nums[j]
			j++
		}
		k++
	}

	for i <= m {
		tmpNums[k] = nums[i]
		i++
		k++
	}

	for j <= r {
		tmpNums[k] = nums[j]
		j++
		k++
	}
	//只需要复制l-r范围内的元素
	for l <= r {
		nums[l] = tmpNums[l]
		l++
	}
}
