package main

import "fmt"

/*
STATUS: Accepted
Runtime: 52ms (Beats 44.80%of users with Go)
Memory: 7.18MB (Beats 99.12%of users with Go)
*/
/*
	Input: nums = [-1,0,1,2,-1,-4]
	Output: [[-1,-1,2],[-1,0,1]]
	Explanation:
	nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
	nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
	nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
	The distinct triplets are [-1,0,1] and [-1,-1,2].
	Notice that the order of the output and the order of the triplets does not matter.
*/
func threeSum(nums []int) [][]int {
	var s [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		var j, k = i + 1, len(nums) - 1
		for {
			if i >= j || j >= k {
				break
			}
			temp := nums[i] + nums[j] + nums[k]
			if temp == 0 {
				if j <= i+1 || nums[j] != nums[j-1] {

				} else {
					s = append(s, []int{nums[i], nums[j], nums[k]})
				}
			}
			if temp <= 0 {
				j += 1
				continue
			}
			if temp > 0 {
				k -= 1
				continue
			}
		}
	}
	return s
}
func main() {
	//num := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	//num := []int{0, 0, 0, 0}
	num := []int{-1, 0, 1, 2, -1, -4}
	s := threeSum(num)
	fmt.Print(s)
}
