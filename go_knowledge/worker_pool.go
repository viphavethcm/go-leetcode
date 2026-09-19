package go_knowledge

import (
	"fmt"
	"sync"
)

func main() {
	jobs := produce()
	results := consume(jobs)
	for num := range results {
		fmt.Println(num)
	}
}

func produce() <-chan int {
	jobs := make(chan int)
	go func() {
		defer close(jobs)
		for i := 0; i < 20; i++ {
			jobs <- i
		}
	}()
	return jobs
}

func consume(jobs <-chan int) <-chan int {
	results := make(chan int)
	wg := sync.WaitGroup{}
	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				results <- job
			}
		}()
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}
