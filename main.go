package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker started")
	fmt.Println("Worker done")
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	fmt.Println("All workers done")
}
