package main

import "fmt"

func removeDuplicates(nums []int) int {
	unique := 0

	for pos := 1; pos <= len(nums)-1; pos++ {
		if nums[pos] != nums[unique] {
			nums[unique+1] = nums[pos]
			unique++
		}
	}
	unique++
	return unique
}

func main() {
	nums := []int{1, 1, 2}
	// Output: 2, nums = [1,2,_]
	k := removeDuplicates(nums)
	fmt.Printf("\nnums = %v \nk = %v", nums, k)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
	k = removeDuplicates(nums)
	fmt.Printf("\nnums = %v \nk = %v", nums, k)

}
