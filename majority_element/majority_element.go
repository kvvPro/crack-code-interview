package main

import "fmt"

func majorityElement(nums []int) int {
	elements := make(map[int]int)

	for _, el := range nums {
		elements[el]++
	}

	for key, val := range elements {
		if val > len(nums)/2 {
			return key
		}
	}

	return -1
}

func main() {
	nums := []int{3, 2, 3}
	// Output: 3
	k := majorityElement(nums)
	fmt.Printf("\nnums = %v \nk = %v", nums, k)

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	// Output: 2
	k = majorityElement(nums)
	fmt.Printf("\nnums = %v \nk = %v", nums, k)
}
