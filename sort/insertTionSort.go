package main

import "fmt"

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if j <= 0 {
				break
			}
			if arr[j-1] > arr[j] {
				arr[j] = arr[j] ^ arr[j-1]
				arr[j-1] = arr[j-1] ^ arr[j]
				arr[j] = arr[j] ^ arr[j-1]
			} else {
				break
			}
			j--
		}
	}
}
func main() {
	a := []int{12, 11, 13, 5, 6}
	fmt.Print(a)
	InsertionSort(a)
	fmt.Print(a)
}
