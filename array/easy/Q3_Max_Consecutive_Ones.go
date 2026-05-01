package array_easy

func findMaxConsecutiveOnes(nums []int) int {
	cnt := 0
	rs := 0
	for _, num := range nums {
		if num == 1 {
			cnt++
			rs = max(rs, cnt)
		} else {
			cnt = 0
		}
	}
	return rs
}

// Kỹ thuật cắm lính canh
// Ý tưởng: Thay vì loop từng phần tử và count số lượng 1 xuất hiện thì ta sẽ biến thành
// bài toán đo khoảng cách giữa các số 0 xuất hiện trong mảng
//func findMaxConsecutiveOnes(nums []int) int {
//	  i,j:=0,0
//    nums=append(nums,0) // Cắm lính canh chốt sổ
//    res:=0
//    for idx,val:=range nums{
//        if val==0{ // mỗi lần gặp số 0
//            j=idx // gán idx vị trí số 0
//            res=max(res,j-i) res = max(res, độ dài giữa 2 số 0)
//            i=j+1
//        }
//    }
//    return res
//}
