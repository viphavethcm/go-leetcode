package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

func singleNumber(nums []int) int {
	sort.Ints(nums)
	// Compare cặp - cặp, nếu list có 3 phần tử thì phần tử lẻ cuối cùng chính là số bị thừa
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
