package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Bài 6 — Atomic flag + Channel để dừng vòng lặp
// Viết một goroutine "producer" liên tục gửi số tăng dần vào channel.
// Dùng biến atomic.Bool làm cờ stopped.
// Khi giá trị cờ được set true từ goroutine khác (giả lập nút "stop" bên ngoài),
// producer phải kiểm tra cờ này mỗi vòng lặp và dừng gửi, đóng channel an toàn.
func main() {
	flag := atomic.Bool{}
	go func() {
		time.Sleep(1 * time.Second)
		flag.Store(true)

	}()
	wg := sync.WaitGroup{}
	ch := produce(&flag)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range ch {
			fmt.Println("data:", data)
		}
	}()
	wg.Wait()
}

func produce(flag *atomic.Bool) chan int {
	ch := make(chan int, 1)
	go func() {
		defer close(ch)
		i := 1
		for {
			if flag.Load() {
				fmt.Println("Stop")
				return
			}
			ch <- i
			i++
		}
	}()
	return ch
}
