package array_easy

func shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))
	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
		res[2*i+1] = nums[i+n]
	}
	return res
}
