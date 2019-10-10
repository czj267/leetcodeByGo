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

//最大堆
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
		//i节点下是否有右节点，如果有是否比左节点大，如果是则取右节点
		if k+1 < l && nums[k] < nums[k+1] {
			k++
		}
		//子节点是否比根节点大，如果是则把子节点的值赋给父节点，否则接着调整子节点
		//因为这里是向下调整，所以这里把i节点下的最大节点给i，即k节点
		if nums[k] > tmp { //注意这里是和根节点对比
			nums[i] = nums[k]
			i = k //
		} else {
			//如果当前的两个子节点没有比父节点大，就可以结束，因为连个子节点的堆已经调整过了
			break
		}
	}
	//最后把值赋给当前节点
	nums[i] = tmp
}
