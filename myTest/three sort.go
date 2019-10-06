package main

import "fmt"

func main() {
	var Ns [][]int
	Ns = append(Ns, []int{1, 0, 2, 0, 1, 2})
	Ns = append(Ns, []int{1, 1, 1, 1, 1, 1})
	Ns = append(Ns, []int{2, 2, 2, 2, 2, 2})
	Ns = append(Ns, []int{0, 0, 0, 0, 0, 0, 0})
	Ns = append(Ns, []int{0, 2, 1, 1, 0, 0})
	Ns = append(Ns, []int{1, 1, 2, 2, 2, 0, 1, 1, 1, 1})
	Ns = append(Ns, []int{0, 0, 0, 2, 2, 0, 1, 1, 0, 0})
	Ns = append(Ns, []int{})
	for _, nums := range Ns {
		threeSort1(nums)
		fmt.Printf("%v\n", nums)
	}
}

/**
以中间的指针为主，移动两边的指针更好理解
下面一种方法是先移动两边的指针再以中间的指针会又更多的边界条件
*/
func threeSort1(nums []int) {
	l := len(nums)
	p1, p3 := -1, l
	for i := 0; i < p3; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 0 {
			p1++
			nums[p1], nums[i] = nums[i], nums[p1]
			i++
		} else if nums[i] == 2 {
			p3--
			nums[p3], nums[i] = nums[i], nums[p3]
		} else {
			panic("error")
		}
	}
}

/**
三路排序
*/
func threeSort(nums []int) {
	l := len(nums) - 1
	p1, p2, p3 := 0, 0, l

	for p1 < p3 && nums[p1] == 0 {
		p1++
	}
	for p3 > p1 && nums[p3] == 2 {
		p3--
	}
	for p2 <= p3 && p1 <= l && p2 <= l && p3 >= 0 {
		if nums[p2] == 0 && p2 > p1 {
			if nums[p1] == 0 {
				p1++
				continue
			}
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p1++
			p2++
		} else if nums[p2] == 2 && p2 < p3 {
			if nums[p3] == 2 {
				p3--
				continue
			}
			nums[p3], nums[p2] = nums[p2], nums[p3]
			p3--
			p2++
		} else {
			p2++
		}
	}
}
