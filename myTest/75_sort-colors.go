package main

func main() {

}
func sortColors(nums []int) {
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
