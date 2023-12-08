package main

import "fmt"

func BinaryInsertionSort(arr []int) {
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
	fmt.Print("hi")
}
