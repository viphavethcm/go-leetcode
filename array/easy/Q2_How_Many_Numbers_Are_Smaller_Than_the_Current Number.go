package main

import (
	"fmt"
)

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{5, 0, 10, 0, 10, 6}))
}

// DP
//func smallerNumbersThanCurrent(nums []int) []int {
//	sorted := make([]int, len(nums))
//	dp := make([]int, 101)
//	copy(sorted, nums)
//	sort.Ints(sorted)
//	for i := 1; i < len(sorted); i++ {
//		if sorted[i] == sorted[i-1] {
//			continue
//		}
//		dp[sorted[i]] = i
//	}
//	for i := 0; i < len(nums); i++ {
//		nums[i] = dp[nums[i]]
//	}
//	return nums
//}

// Prefix sum
func smallerNumbersThanCurrent(nums []int) []int {
	dp := make([]int, 100)
	// Count tần suất xuất hiện của các số trong array
	for i := 0; i < len(nums); i++ {
		dp[nums[i]]++
	}
	// Tìm tổng các số <= i = Số lần xuất hiện của i + Tổng các số <= i - 1
	for i := 1; i < 100; i++ {
		dp[i] += dp[i-1]
	}
	// Tìm số lần xuất hiện của i = Tổng các số <= i - 1 -> Tổng các số < 8 = Tổng các số <=7
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = 0
		} else {
			nums[i] = dp[nums[i]-1]
		}
	}
	return nums
}
