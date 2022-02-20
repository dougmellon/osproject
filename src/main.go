package main

import (
	"fmt"
	"time"
)

func thread(threadName int) {
	fmt.Println("Creating thread ", threadName)
}

func main() {

	totalThreads := 0

	// Start timer
	start := time.Now()

	// Create 100,000 threads and then destroy them
	for i := 0; i < 10; i++ {
		go thread(i)
		time.Sleep(time.Millisecond * 250) // Delay by a quarter of a second
		totalThreads++
	}

	// End timer
	elapsed := time.Since(start)
	fmt.Println("Start time: ", start)
	fmt.Println("Elapsed time: ", elapsed)
}
