package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _, exists := hashmap[complement]; exists {
			return []int{i, hashmap[complement]}
		}
		hashmap[nums[i]] = i
	}
	return nil
}

func main() {
	target := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("indices: %v", target)
}
