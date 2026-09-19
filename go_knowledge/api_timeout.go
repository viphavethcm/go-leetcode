package go_knowledge

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				fmt.Println("Working...")
			case <-ctx.Done():
				fmt.Println("Timed out")
			}
		}
	}(ctx)
	wg.Wait()
}
