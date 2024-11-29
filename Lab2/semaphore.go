package main

import (
	"fmt"
	"sync"
	"time"
)

// Semaphore struct
type Semaphore struct {
	tokens chan struct{}
}

// NewSemaphore creates a new semaphore with the given initial count
func NewSemaphore(count int) *Semaphore {
	return &Semaphore{
		tokens: make(chan struct{}, count),
	}
}

// Acquire a token from the semaphore
func (s *Semaphore) Acquire() {
	s.tokens <- struct{}{}
}

// Release a token back to the semaphore
func (s *Semaphore) Release() {
	<-s.tokens
}

func main() {
	// Example usage of the semaphore
	sem := NewSemaphore(2) // Semaphore with 2 tokens
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem.Acquire()
			fmt.Printf("Goroutine %d entering critical section\n", id)
			time.Sleep(1 * time.Second) // Simulate work
			fmt.Printf("Goroutine %d leaving critical section\n", id)
			sem.Release()
		}(i)
	}

	wg.Wait()
}

//so at the start 2 go routines can acquire 2 tokens (line 33, can add more) and enter the critical section
//then one go routine finishes, and releases the token (line 38)
//then another go routine can acquire the token and enter the critical section
//the semaphore prevents both routines from releasing or acquiring at the same run time (a flip flop)
//when all go routines finish the remaining 2 are released at the end
