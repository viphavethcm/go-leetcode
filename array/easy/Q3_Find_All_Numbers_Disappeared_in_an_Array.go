package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	rs := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		j := abs(nums[i]) - 1
		if nums[j] < 0 {
			continue
		}
		nums[j] *= -1
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			missing := i + 1
			rs = append(rs, missing)
		}
	}
	return rs
}

// 1,2,2,3,3,4,7,8
