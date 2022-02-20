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

	// Start timer.
	start := time.Now()

	// Create 100,000 threads and then destroy them.
	for i := 1; i <= 100000; i++ {
		go thread(i)
		totalThreads++
	}

	// End timer
	elapsed := time.Since(start).Milliseconds() // Determine the elapsed time in milliseconds.

	// Print out report on start and elapsed times
	fmt.Println("Start time: ", start)
	fmt.Println("Elapsed time: ", elapsed)
}
