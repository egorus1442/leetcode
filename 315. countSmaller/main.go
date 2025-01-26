package main

import "sort"

func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	counts := make([]int, len(nums))

	sorted := []int{}

	for i := len(nums) - 1; i >= 0; i-- {
		pos := sort.Search(len(sorted), func(j int) bool {
			return sorted[j] >= nums[i]
		})

		counts[i] = pos

		sorted = append(sorted, 0)
		copy(sorted[pos+1:], sorted[pos:])
		sorted[pos] = nums[i]
	}

	return counts
}
