package two_pointer_easy

/*
https://leetcode.com/problems/remove-element/description
*/

func removeElement(nums []int, val int) int {
	i := 0
	insertIndex := 0
	for i < len(nums) {
		if nums[i] == val {
			nums[insertIndex] = nums[i]
			insertIndex++
		}
		i++
	}
	return insertIndex
}
