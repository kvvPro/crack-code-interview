package main

import "fmt"

func maxProfit(prices []int) int {

	forwardPtr := 0
	innerPtr := 1
	profit := 0

	for forwardPtr < len(prices)-1 {
		for innerPtr <= len(prices)-1 {
			if prices[innerPtr]-prices[forwardPtr] > profit {
				profit = prices[innerPtr] - prices[forwardPtr]
			}
			innerPtr++
		}
		forwardPtr++
		innerPtr = forwardPtr + 1
	}
	return profit
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	// Output: 5
	profit := maxProfit(nums)
	fmt.Printf("\nnums = %v \nk = %v", nums, profit)

	nums = []int{7, 6, 4, 3, 1}
	// Output: 0
	profit = maxProfit(nums)
	fmt.Printf("\nnums = %v \nk = %v", nums, profit)
}
