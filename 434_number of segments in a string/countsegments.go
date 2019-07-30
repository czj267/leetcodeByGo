package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, my name is John"
	fmt.Println("Output:", countSegments(s))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func countSegments(s string) int {
	return len(strings.Fields(s))
}
