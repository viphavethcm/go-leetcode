package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(2))
}

// 1,2,3,4,5,6,7,8,9,10
func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func isBadVersion(version int) bool {
	return true
}
