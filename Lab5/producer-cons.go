package main

import (
	"fmt"
	"sync"
	"time"
)

// Event represents an event to be processed
type Event struct {
	id int
}

// createEvent creates a new event with a given id
func createEvent(id int) Event {
	return Event{id: id}
}

// process processes the event
func (e Event) process() {
	fmt.Printf("Processing event %d\n", e.id)
}

// producer creates events and adds them to the buffer
func producer(buffer chan<- Event, id int, numLoops int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numLoops; i++ {
		event := createEvent(i)
		buffer <- event //will automatically ensure no other producer adds to the buffer
		fmt.Printf("Producer %d produced event %d\n", id, i)
		time.Sleep(time.Millisecond * 100) // Simulate work
	}
}

// consumer takes events from the buffer and processes them
func consumer(buffer <-chan Event, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for event := range buffer { //will automatically ensure no other consumer retrieves from the buffer
		fmt.Printf("Consumer %d consumed event %d\n", id, event.id)
		event.process()
		time.Sleep(time.Millisecond * 150) // Simulate work
	}
}

func main() {
	const numProducers = 3
	const numConsumers = 2
	const bufferSize = 10
	const numLoops = 5

	buffer := make(chan Event, bufferSize)
	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	// Start producers
	for i := 0; i < numProducers; i++ {
		producerWg.Add(1)
		go producer(buffer, i, numLoops, &producerWg)
	}

	// Start consumers
	for i := 0; i < numConsumers; i++ {
		consumerWg.Add(1)
		go consumer(buffer, i, &consumerWg)
	}

	// Wait for all producers to finish
	producerWg.Wait()
	close(buffer)

	// Wait for all consumers to finish
	consumerWg.Wait()
}
