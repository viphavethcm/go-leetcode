package binary_search

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left, right := 2, x/2
	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid
		if square < x {
			left = mid + 1
		} else if square > x {
			right = mid - 1
		} else {
			return mid
		}
	}
	return right
}
