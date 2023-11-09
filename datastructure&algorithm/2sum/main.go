package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var a = make(map[int]int)
	for i, num := range nums {
		if index, check := a[num]; check == true {
			return []int{index, i}
		}
		a[target-num] = i
	}
	return []int{}
}
func main() {
	a := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Print(a)
}
