package main

import (
	"fmt"
	"sort"
)

func main() {
	g, s := []int{1, 3, 2}, []int{1, 1}
	res := findContentChildren(g, s)
	fmt.Println(res)
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gi, si, res := len(g)-1, len(s)-1, 0
	for gi >= 0 && si >= 0 {
		if s[si] >= g[gi] {
			si--
			gi--
			res++
		} else {
			gi--
		}
	}
	return res
}
