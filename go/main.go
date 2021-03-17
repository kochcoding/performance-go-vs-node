package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := 40
	var wg sync.WaitGroup
	wg.Add(jobs)

	start := time.Now()
	for i := 0; i < jobs; i++ {
		go job(i, &wg)
	}
	wg.Wait()
	fmt.Printf("Runtime %v\n", time.Since(start))
}

func job(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	_ = fibonacci(n)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
