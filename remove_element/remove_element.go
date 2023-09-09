package main

import "fmt"

func removeElement(nums []int, val int) int {
	notEqual := 0

	pointerForward := 0
	pointerBack := len(nums) - 1

	for pointerForward <= pointerBack {
		if nums[pointerForward] == val {
			// exchange
			nums[pointerForward] = nums[pointerBack]
			pointerBack--
		} else {
			notEqual++
			pointerForward++
		}
	}

	return notEqual
}

func removeElement_2(nums []int, val int) int {
	notEqual := 0

	for _, el := range nums {
		if el != val {
			nums[notEqual] = el
			notEqual++
		}
	}

	return notEqual
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	k := removeElement_2(nums, val)
	// Output: 2, nums = [2,2,_,_]
	fmt.Printf("\nnums = %v \nk = %v", nums, k)

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	k = removeElement_2(nums, val)
	// Output: 5, nums = [0,1,4,0,3,_,_,_]
	fmt.Printf("\nnums = %v \nk = %v", nums, k)
}
