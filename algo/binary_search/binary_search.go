package binary_search

func BinarySearch(nums []int, l, r, target int) int {
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] <= target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}
