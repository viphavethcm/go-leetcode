package two_pointer_easy

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-array/description
*/

func removeDuplicates(nums []int) int {
	i := 1
	insertIndex := 1
	for i < len(nums) {
		if nums[i] != nums[i-1] {
			nums[insertIndex] = nums[i]
			insertIndex++
		}
		i++
	}
	return insertIndex
}
