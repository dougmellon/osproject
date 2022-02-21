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

var c = Container{counter: 0}

// Lock the container.
func (c *Container) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

// Initiate the waitgroup.
var wg sync.WaitGroup

// thread.
func thread(threadName int) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		go childThread(threadName, i)
		c.inc()
	}

}

// childThread.
func childThread(threadName int, childThread int) {
	fmt.Printf("Creating thread %d - %d \n", threadName, childThread)
}

func main() {

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
