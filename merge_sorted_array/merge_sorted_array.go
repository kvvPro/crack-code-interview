package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1[:m+n], nums2[:n])
		return
	}
	start := 0
	actualLength := m
	for i := 0; i < n; i++ {
		for j := start; j < actualLength; j++ {
			if nums1[j] > nums2[i] {
				// put nums2[i] after nums1[j]
				nums1 = append(nums1[:j], append([]int{nums2[i]}, nums1[j:actualLength]...)...)
				start = j + 1
				actualLength++
				break
			} else if j == actualLength-1 {
				nums1 = append(nums1[:actualLength], nums2[i])
				start = j + 1
				actualLength++
				break
			}
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Printf("\nnums1: %v", nums1)

	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0
	merge(nums1, m, nums2, n)
	fmt.Printf("\nnums1: %v", nums1)

	nums1 = []int{1, 0}
	m = 1
	nums2 = []int{2}
	n = 1
	merge(nums1, m, nums2, n)
	fmt.Printf("\nnums1: %v", nums1)

	nums1 = []int{4, 5, 6, 0, 0, 0}
	m = 3
	nums2 = []int{1, 2, 3}
	n = 3
	merge(nums1, m, nums2, n)
	fmt.Printf("\nnums1: %v", nums1)

	nums1 = make([]int, 0)
	m = 0
	nums2 = []int{1}
	n = 1
	merge(nums1, m, nums2, n)
	fmt.Printf("\nnums1: %v", nums1)
}
