package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	k := 2
	vals := []int{3, 1, 4, -1, 2}
	// k := 3
	// vals := []int{5, 3, 6, 2, 4, -1, -1, 1}
	root := common.MakeTree(vals)
	fmt.Println("Output:", kthSmallest(root, k))
	fmt.Println("Output:", myKthSmallest(root, k))
}

func myKthSmallest(root *TreeNode, k int) int {
	kk = k
	midOrder(root)
	return val
}

var val, kk int

func midOrder(root *TreeNode) {
	if root == nil {
		return
	}
	midOrder(root.Left)
	kk--
	if kk == 0 {
		val = root.Val
		return
	}
	midOrder(root.Right)
}

func kthSmallest1(root *TreeNode, k int) int {
	val := -1
	midOrder1(root, &k, &val)
	return val
}

func midOrder1(root *TreeNode, k *int, val *int) {
	if root == nil || *k == 0 {
		return
	}
	midOrder1(root.Left, k, val)
	*k--
	if *k == 0 {
		*val = root.Val
		//val = &root.Val 这种写法是错误的
		return
	}
	midOrder1(root.Right, k, val)
}

//执行用时 :12 ms, 在所有 Go 提交中击败了99.30%的用户
//内存消耗 :6 MB, 在所有 Go 提交中击败了56.10%的用户
func kthSmallest(root *TreeNode, k int) int {
	kv := 0
	var findKth func(r *TreeNode)
	findKth = func(r *TreeNode) {
		if r == nil || k == 0 {
			return
		}
		findKth(r.Left)
		k--
		if k == 0 {
			kv = r.Val
			return
		}
		findKth(r.Right)
	}
	findKth(root)
	return kv
}
