package prefix_sum_easy

/*
https://leetcode.com/problems/find-the-middle-index-in-array/description
*/

func findMiddleIndex(nums []int) int {
	sum, leftSum := 0, 0
	for _, num := range nums {
		sum += num
	}
	for idx, num := range nums {
		rightSum := sum - leftSum - num
		if leftSum == rightSum {
			return idx
		}
		leftSum += num
	}
	return -1
}
