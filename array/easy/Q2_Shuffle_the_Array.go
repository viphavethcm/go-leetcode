package main

import "fmt"

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	for _, num := range shuffle(nums, 3) {
		fmt.Println(num)
	}
}

func shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))
	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
		res[2*i+1] = nums[i+n]
	}
	return res
}
