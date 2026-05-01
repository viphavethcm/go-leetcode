package array_easy

// 1,2,3 -> 1,2,4 | 4,3,2,1 -> 4,3,9,9 | 9 -> 1,0 , 9999
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	// append = 0 vì lý do số sẽ tăng 1 hàng đơn vị, VD: 99 = 100
	digits = append(digits, 0)
	digits[0] = 1
	return digits
}
