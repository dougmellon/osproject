package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

// thread
func thread(threadName int) {
	defer wg.Done()
	fmt.Println("Creating thread ", threadName)
}

func main() {
	// Counter for total threads created.
	totalThreads := 0

	// Start timer.
	start := time.Now()

	// Create 100,000 threads and then destroy them.
	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		go thread(i)
		totalThreads++
	}

	// End timer
	elapsed := time.Since(start).Milliseconds() // Determine the elapsed time in milliseconds.

	// Wait for the goroutines to finish.
	wg.Wait()

	// Print out report on start and elapsed times.
	fmt.Println("Start time: ", start)
	fmt.Println("Elapsed time: ", elapsed)
	fmt.Println("Total threads: ", totalThreads)
}
