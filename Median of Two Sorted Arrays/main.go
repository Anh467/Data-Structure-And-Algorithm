package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j int = 0, 0
	var a, b int
	res := (len(nums1)+len(nums2))/2 + 1
	//int(math.Ceil((float64(lenNum1) + float64(lenNum2)) / 2))
	for k := 0; k < res; k++ {
		if nums1[i] < nums2[j] {
			if k == res-2 {
				a = nums1[i]
			} else if k == res-1 {
				b = nums1[i]
			}
			i++
		} else {
			if k == res-2 {
				a = nums2[j]
			} else if k == res-1 {
				b = nums2[j]
			}
			j++
		}
	}
	return (float64(a) + float64(b)) / 2
}

func main() {
	nums1 := [][]int{
		{1, 3},
		{1, 2},
	}
	nums2 := [][]int{
		{2},
		{3, 4},
	}
	test := 1
	s := findMedianSortedArrays(nums1[test], nums2[test])
	fmt.Print(s)
}
