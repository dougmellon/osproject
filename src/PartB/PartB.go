package main

import (
	"fmt"
	"sync"
	"time"
)

// Container - Mutex counter for concurrent access
type Container struct {
	mu      sync.Mutex
	counter int
}

// Lock the container.
func (c *Container) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

// Initiate the waitgroup.
var wg sync.WaitGroup

// Total thread counter.
var totalThreads = 0
var totalChildThreads = 0

// thread.
func thread(threadName int) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		go childThread(threadName, i)
		totalChildThreads++
	}

}

// childThread.
func childThread(threadName int, childThread int) {
	fmt.Printf("Creating thread %d - %d \n", threadName, childThread)
}

func main() {

	c := Container{counter: 0}

	// Start timer.
	start := time.Now()

	// Create 100 parent threads.
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go thread(i)
		c.inc()
	}

	// End timer
	elapsed := time.Since(start).Milliseconds() // Determine the elapsed time in milliseconds.

	// Wait for the goroutines to finish.
	wg.Wait()

	// Print out report on start and elapsed times.
	fmt.Println("Start time: ", start)
	fmt.Println("Elapsed time: ", elapsed)
	fmt.Println("Total threads: ", c.counter)
}
