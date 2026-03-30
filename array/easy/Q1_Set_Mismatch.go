package main

import "go-leetcode/common"

func findErrorNums(nums []int) []int {
	missing := 0
	dup := 0
	for i := 0; i < len(nums); i++ {
		j := common.Abs(nums[i]) - 1
		if nums[j] < 0 {
			dup = common.Abs(nums[i])
			continue
		}
		nums[j] *= -1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			missing = i + 1
		}
	}
	return []int{dup, missing}
}
