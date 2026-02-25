package main

func majorityElement(nums []int) int {
	cnt := 0
	var majorityNum int
	for _, num := range nums {
		if cnt == 0 {
			majorityNum = num
			cnt++
		} else {
			if majorityNum != num {
				cnt--
			} else {
				cnt++
			}
		}
	}
	return majorityNum
}
