package array_easy

// Cyclic Sort
func findDisappearedNumbers(nums []int) []int {
	i := 0
	n := len(nums)
	rs := make([]int, 0, len(nums))
	for i < n {
		rightIdx := nums[i] - 1
		if nums[i] > 0 && nums[i] <= n && nums[rightIdx] != nums[i] {
			nums[i], nums[rightIdx] = nums[rightIdx], nums[i]
		} else {
			i++
		}
	}
	for idx, num := range nums {
		if num != idx+1 {
			rs = append(rs, idx+1)
		}
	}
	return rs
}

// In-Place Index Mapping
//func findDisappearedNumbers(nums []int) []int {
//	rs := make([]int, 0, len(nums))
//	for i := 0; i < len(nums); i++ {
//		j := common.Abs(nums[i]) - 1
//		if nums[j] < 0 {
//			continue
//		}
//		nums[j] *= -1
//	}
//
//	for i := 0; i < len(nums); i++ {
//		if nums[i] > 0 {
//			missing := i + 1
//			rs = append(rs, missing)
//		}
//	}
//	return rs
//}
