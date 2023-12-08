package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}

	return result
}
func main() {
	a := [][]int{
		{2, 2, 1},
		{4, 1},
		{1},
	}
	s := singleNumber(a[1])
	fmt.Print(s)
}
