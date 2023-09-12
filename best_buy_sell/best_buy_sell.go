package main

import (
	"fmt"
	"math"
)

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

func maxProfit_ideal(prices []int) int {
	maxProfit := 0

	lowestValue := math.MaxInt32

	for _, value := range prices {
		if value < lowestValue {
			lowestValue = value
		}

		if value-lowestValue > maxProfit {
			maxProfit = value - lowestValue
		}
	}

	return maxProfit
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	// Output: 5
	profit := maxProfit_ideal(nums)
	fmt.Printf("\nnums = %v \nk = %v", nums, profit)

	nums = []int{7, 6, 4, 3, 1}
	// Output: 0
	profit = maxProfit_ideal(nums)
	fmt.Printf("\nnums = %v \nk = %v", nums, profit)
}
