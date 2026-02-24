package main

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-array/description
*/

func main() {

}

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
