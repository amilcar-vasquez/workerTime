package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(2 * time.Second) //Replace this
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	for i := 1; i <= 3; i++ {
		go worker(i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("All workers done") //Replace this
}
