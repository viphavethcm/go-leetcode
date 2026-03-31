package main

func main() {

}

// 1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1
	for p >= 0 {
		if p2 < 0 {
			break
		}
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
			p--
		} else {
			nums1[p] = nums2[p2]
			p2--
			p--
		}
	}
}
