package main

import (
	"fmt"
	"strings"
)

func main() {
	ips := restoreIpAddresses1("25525511135")
	fmt.Printf("%v", ips)
}

/**
回溯法
*/
func restoreIpAddresses1(s string) []string {
	var res []string
	findIp(s, []string{}, &res)
	return res
}

func findIp(s string, subs []string, res *[]string) {
	if len(subs) == 4 {
		if len(s) == 0 {
			*res = append(*res, subs[0]+"."+subs[1]+"."+subs[2]+"."+subs[3])
		} else {
			return
		}
	}
	for i := 1; i <= 3; i++ {
		if len(s) < i {
			return
		}
		sub := s[0:i]
		if len(sub) == 3 && strings.Compare(sub, "255") > 0 {
			return
		}
		if len(sub) > 1 && sub[0] == '0' {
			return
		}
		subs = append(subs, sub)
		findIp(s[i:], subs, res)
		subs = subs[0 : len(subs)-1]
	}
}

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	backtrack(make([]string, 0), s, &result)
	return result
}

func backtrack(putStr []string, s string, result *[]string) {
	resultLen := len(putStr)
	if resultLen == 4 {
		if len(s) == 0 {
			*result = append(*result, putStr[0]+"."+putStr[1]+"."+putStr[2]+"."+putStr[3])
		}
		return
	}
	//
	for i := 1; i <= 3; i++ {
		if len(s) < i {
			return
		}
		str := s[0:i]
		strLen := len(str)
		if strLen == 3 && strings.Compare(str, "255") > 0 {
			return
		}
		if strLen > 1 && s[0] == '0' {
			return
		}
		putStr = append(putStr, str)
		backtrack(putStr, s[i:], result)
		putStr = putStr[:len(putStr)-1]
	}
}
