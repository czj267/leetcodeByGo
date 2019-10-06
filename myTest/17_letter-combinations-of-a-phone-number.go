package main

func main() {
	letterCombinations("23")
}

var letterMap = [...]string{
	" ",
	"",
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}
var res []string

func letterCombinations(digits string) []string {
	res = []string{}
	if digits == "" {
		return res
	}
	s := ""
	findCombinations(digits, 0, s)
	return res
}

func findCombinations(digits string, index int, s string) {
	if index == len(digits) {
		res = append(res, s)
		return
	}
	num := digits[index]
	letters := letterMap[num-'0']
	for i := 0; i < len(letters); i++ {
		findCombinations(digits, index+1, s+string(letters[i]))
	}
	return
}
