package go_knowledge

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Request struct {
	ID int
}

type RateLimiter struct {
	requestChan chan Request
	// Dùng atomic để tracking metric hiệu năng cao
	successCount  int64
	rejectedCount int64
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		requestChan: make(chan Request, 100),
	}
}
func (rl *RateLimiter) Start(ctx context.Context) {
	// Cứ mỗi 100ms sẽ có một "vạch thời gian" mới để reset quota
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	allowedInWindow := 0
	const maxPerWindow = 3

	for {
		select {
		case <-ctx.Done():
			// --- 🛑 CHẶNG GRACEFUL SHUTDOWN: VÉT SẠCH CHANNEL ---
			fmt.Println("⏳ Đang vét nốt các request còn sót lại trong queue...")

			for {
				select {
				case _, ok := <-rl.requestChan:
					if !ok {
						return // Kênh đã đóng và hết hàng -> Thoát hoàn toàn
					}
					// Vì hệ thống đã shutdown, các request còn sót lại này
					// có thể xử lý tiếp hoặc đánh dấu là bị Từ chối do hệ thống tắt.
					// Ở đây tui chọn đếm vào rejectedCount cho đúng thực tế.
					atomic.AddInt64(&rl.rejectedCount, 1)

				default:
					// Kỹ thuật then chốt: Khi rl.requestChan không còn hàng để đọc ngay lập tức,
					// nhánh default sẽ được kích hoạt để bẻ gãy vòng lặp, tránh bị nghẽn (block).
					return
				}
			}

		case <-ticker.C:
			// Chỉ có ông ticker này mới có quyền reset cửa ngõ
			allowedInWindow = 0

		case _, ok := <-rl.requestChan:
			if !ok {
				return
			}

			// Logic giới hạn tần suất chuẩn
			if allowedInWindow < maxPerWindow {
				atomic.AddInt64(&rl.successCount, 1)
				allowedInWindow++
			} else {
				// Vượt quá quota thì reject, giữ nguyên allowedInWindow = maxPerWindow
				atomic.AddInt64(&rl.rejectedCount, 1)
			}
		}
	}
}

func main() {
	// Khởi tạo Context có thể cancel bằng tay
	ctx, cancel := context.WithCancel(context.Background())

	limiter := NewRateLimiter()
	var wg sync.WaitGroup

	// Chạy Rate Limiter ngầm
	wg.Add(1)
	go func() {
		defer wg.Done()
		limiter.Start(ctx)
	}()

	// Giả lập luồng đẩy request vào liên tục (Producer)
	go func() {
		for i := 1; i <= 50; i++ {
			limiter.requestChan <- Request{ID: i}
			time.Sleep(15 * time.Millisecond) // Đẩy request siêu nhanh
		}
	}()

	// Cho hệ thống chạy thử 500ms
	time.Sleep(200 * time.Millisecond)

	// --- KÍCH HOẠT GRACEFUL SHUTDOWN ---
	fmt.Println("\n🛑 Đang kích hoạt Graceful Shutdown...")
	cancel() // Phát tín hiệu dừng qua Context

	fmt.Println("\n--- 📊 KẾT QUẢ THỐNG KÊ ---")
	fmt.Printf("✅ Requests được xử lý (Success): %d\n", atomic.LoadInt64(&limiter.successCount))
	fmt.Printf("❌ Requests bị từ chối (Rejected): %d\n", atomic.LoadInt64(&limiter.rejectedCount))
}

//Nhận request liên tục từ client qua một channel.
//
//Cứ mỗi 100ms, hệ thống chỉ cho phép tối đa 3 requests được đi qua (xử lý thành công), các request thừa trong khoảng thời gian đó phải bị từ chối (Rate Limited).
//
//Sử dụng atomic để đếm tổng số request đã xử lý thành công trên toàn hệ thống mà không dùng Mutex để tránh contention cao.
//
//Hệ thống phải hỗ trợ Graceful Shutdown: Khi nhấn lệnh dừng (thông qua context), hệ thống phải dừng nhận request mới, xử lý nốt các request đang nghẽn trong channel trong vòng tối đa 500ms, sau đó tắt hoàn toàn.
