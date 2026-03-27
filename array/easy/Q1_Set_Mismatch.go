package main

func findErrorNums(nums []int) []int {
	missing := 0
	dup := 0
	for i := 0; i < len(nums); i++ {
		j := abs(nums[i]) - 1
		if nums[j] < 0 {
			dup = abs(nums[i])
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

// 3, 2, 3, 4, 6, 5
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
