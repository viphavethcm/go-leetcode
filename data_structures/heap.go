package data_structures

import (
	"container/heap"
)

// ---------------------------------------------------------
// BƯỚC 1: KHAI BÁO TYPE VÀ IMPLEMENT 5 HÀM BẮT BUỘC
// ---------------------------------------------------------

// Định nghĩa một type mới dựa trên slice của int
type IntHeap []int

// 1. Len: Trả về kích thước của heap
func (h IntHeap) Len() int { return len(h) }

// 2. Less: Quyết định đây là Min-Heap hay Max-Heap
// Dùng < cho Min-Heap (thằng nhỏ nhất nổi lên đỉnh)
// Dùng > cho Max-Heap
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

// 3. Swap: Đổi chỗ 2 phần tử
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// 4. Push: Thêm phần tử vào cuối slice (chưa cần sắp xếp, thư viện sẽ tự lo)
// Lưu ý pointer receiver (*IntHeap) vì ta đang làm thay đổi độ dài của slice
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// 5. Pop: Lấy phần tử ở cuối slice ra (thư viện đã dồn thằng Min xuống cuối cho mình)
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]     // Lấy phần tử cuối
	*h = old[0 : n-1] // Cắt bỏ phần tử cuối
	return x
}

// ---------------------------------------------------------
// BƯỚC 2: HÀM LOGIC CHÍNH XỬ LÝ BÀI TOÁN
// ---------------------------------------------------------

func findKthLargest(nums []int, k int) int {
	// Khởi tạo heap rỗng
	h := &IntHeap{}
	heap.Init(h)

	for _, num := range nums {
		// Nhét phần tử vào Heap
		heap.Push(h, num)

		// Nếu kích thước Heap vượt quá K, ta "đá" thằng nhỏ nhất (ở đỉnh) ra ngoài.
		// Nhờ vậy, Heap luôn chỉ giữ lại đúng K thằng BỰ NHẤT mà nó từng gặp.
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Sau khi duyệt hết mảng, thằng nằm trên đỉnh Min-Heap lúc này
	// chính là thằng nhỏ nhất trong "Hội K thằng bự nhất"
	// => Nó chính là phần tử lớn thứ K của toàn bộ mảng!
	// h[0] KHÔNG hoạt động trực tiếp trên kiểu pointer, nên ta phải ép kiểu về []int hoặc dùng Pop
	return heap.Pop(h).(int)
}
