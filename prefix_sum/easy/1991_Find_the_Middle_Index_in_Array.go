package main

import "fmt"

/*
https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description
*/
func main() {
	nums := []int{2, 3, -1, 8, 4}
	fmt.Println(findMiddleIndex(nums))
}

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
