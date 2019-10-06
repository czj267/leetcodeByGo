package main

import "fmt"

func main() {
	res := combine(4, 2)
	fmt.Println(res)
}

/**
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
*/
func combine(n int, k int) [][]int {
	var res [][]int
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	combination(n, k, 1, []int{}, &res)
	return res
}

func combination(n, k, index int, com []int, res *[][]int) {
	if len(com) == k {
		_com := make([]int, k)
		copy(_com, com)
		*res = append(*res, _com)
		return
	}
	//注意这里进行了剪枝
	for i := index; i <= n-(k-len(com))+1; i++ {
		com = append(com, i)
		combination(n, k, i+1, com, res)
		com = com[:len(com)-1]
	}
}
