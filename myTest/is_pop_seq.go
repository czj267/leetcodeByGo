package main

import (
	"fmt"
	stack2 "github.com/golang-collections/collections/stack"
)

/**
判断一个序列是否是一个stack的弹出序列
*/
func main() {
	s1 := []int{1, 2, 3, 4, 5}
	//s2 := []int{4, 5, 3, 2, 1}
	s2 := []int{4, 3, 5, 1, 2}
	s1 = []int{1}
	s2 = []int{1}

	res := isPopStackSeq(s1, s2)
	fmt.Println(res)

}

/*
 * Complete the simpleArraySum function below.
 */
func simpleArraySum(ar []int32) int32 {
	/*
	 * Write your code here.
	 */
	l := len(ar)
	sum := int32(0)
	for i := 0; i < l; i++ {
		sum += ar[i]
	}
	return sum

}
func isPopStackSeq(s1, s2 []int) bool {
	l := len(s1)
	if l != len(s2) {
		return false
	}

	stack := stack2.New()

	i1 := 0
	i2 := 0
	ps := false
	for i2 < l {
		if stack.Len() == 0 && i1 < l {
			stack.Push(s1[i1])
			i1++
		}
		v1 := stack.Peek()
		if v1 == s2[i2] {
			stack.Pop()
			i2++
		} else {
			if i1 < l {
				stack.Push(s1[i1])
				i1++
			} else {
				ps = false
				break
			}
		}
		if stack.Len() == 0 && i1 == l && i2 == l {
			ps = true
		}
	}
	return ps

}
