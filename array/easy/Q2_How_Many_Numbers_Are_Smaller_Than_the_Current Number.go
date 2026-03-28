package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{6, 5, 4, 8}))
}
func smallerNumbersThanCurrent(nums []int) []int {
	sorted := make([]int, len(nums))
	dp := make([]int, 101)
	copy(sorted, nums)
	sort.Ints(sorted)
	for i := 1; i < len(sorted); i++ {
		if sorted[i] == sorted[i-1] {
			continue
		}
		dp[sorted[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = dp[nums[i]]
	}
	return nums
}
