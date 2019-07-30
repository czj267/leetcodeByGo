package main

import "fmt"

func main() {
	label := 14
	// label := 26
	fmt.Println("Output:", pathInZigZagTree(label))
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Path In Zigzag Labelled Binary Tree.
// Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Path In Zigzag Labelled Binary Tree.
func pathInZigZagTree(label int) []int {
	if label == 1 {
		return []int{1}
	}
	n := 2
	level := 1
	for n*2 <= label {
		level++
		n *= 2
	}
	level++
	var index int
	if level%2 == 0 {
		index = n*2 - 1 - label
	} else {
		index = label - n
	}
	path := make([]int, level)
	path[0] = 1
	path[level-1] = label
	for i := level - 2; i > 0; i-- {
		level--
		n /= 2
		index /= 2
		if level%2 == 0 {
			path[i] = n*2 - 1 - index
		} else {
			path[i] = n + index
		}
	}
	return path
}

// Runtime: 12 ms, faster than 8.33% of Go online submissions for Path In Zigzag Labelled Binary Tree.
// Memory Usage: 24.3 MB, less than 100.00% of Go online submissions for Path In Zigzag Labelled Binary Tree.
func pathInZigZagTreeV1(label int) []int {
	tree := make([][]int, 1)
	leaves := []int{1}
	tree[0] = leaves
	n := 2
	level := 1
	for n <= label {
		level++
		leaves = make([]int, n)
		if level%2 == 0 {
			for i := 2*n - 1; i >= n; i-- {
				leaves[2*n-1-i] = i
				if i == label {
					leaves = leaves[0 : 2*n-i]
					break
				}
			}
		} else {
			for i := n; i < 2*n; i++ {
				leaves[i-n] = i
				if i == label {
					leaves = leaves[0 : i-n+1]
					break
				}
			}
		}
		tree = append(tree, leaves)
		n *= 2
	}
	path := make([]int, len(tree))
	index := len(leaves) - 1
	level--
	path[len(path)-1] = label
	for i := len(path) - 2; i >= 0; i-- {
		level--
		path[i] = tree[level][index/2]
		index /= 2
	}
	return path
}
