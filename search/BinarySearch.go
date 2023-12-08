package main

import "fmt"

/*
	STATUS: Accepted
	Runtime: 30ms Beats 72.58%of users with Go
	Memory: 6.75MB Beats 80.28%of users with Go
*/

func search(nums []int, target int) int {
	var l, r int = 0, len(nums) - 1
	for {
		if l >= r {
			if nums[r] == target {
				return r
			}
			break
		}
		temp := (l + r + 1) / 2
		if nums[temp] == target {
			return temp
		}
		if nums[temp] > target {
			r = temp - 1

			continue
		}
		if nums[temp] < target {
			l = temp + 1
			continue
		}
	}
	return -1
}

func main() {
	arr := []int{-1, 0, 3, 5, 9, 12}
	s := search(arr, 9)
	fmt.Print(s)
}
