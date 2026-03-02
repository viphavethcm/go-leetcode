package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 5, 6}
	target := 2
	fmt.Println(searchInsert(arr, target))
}

// nums = [1,3,5,6], target = 7
// nums = [1,3,5,6], target = 2
// nums = [1,3,5,6], target = 5
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
