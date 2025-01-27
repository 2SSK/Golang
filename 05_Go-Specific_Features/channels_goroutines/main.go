package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	bufferSize   = 5  // Size of the buffer (queue)
	numProducers = 2  // Number of producers
	numConsumers = 3  // Number of consumers
	numItems     = 10 // Total items each producer generates
)

// Produce adds items to the buffer
func Produce(id int, buffer chan int, producerWg *sync.WaitGroup) {
	defer producerWg.Done()

	for i := 0; i < numItems; i++ {
		item := rand.Intn(100) // Generate a random item
		buffer <- item         // Add the item to the buffer
		fmt.Printf("Producer %d produced: %d\n", id, item)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

// Consume removes items from the buffer and processes them
func Consume(id int, buffer chan int, consumerWg *sync.WaitGroup) {
	defer consumerWg.Done()

	for item := range buffer { // Read items from the buffer until it's closed
		fmt.Printf("\tConsumer %d consumed: %d\n", id, item)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	buffer := make(chan int, bufferSize) // Create a buffered channel for the shared buffer

	var producerWg sync.WaitGroup // WaitGroup for producers
	var consumerWg sync.WaitGroup // WaitGroup for consumers

	// Start producers
	for i := 1; i <= numProducers; i++ {
		producerWg.Add(1)
		go Produce(i, buffer, &producerWg)
	}

	// Start consumers
	for i := 1; i <= numConsumers; i++ {
		consumerWg.Add(1)
		go Consume(i, buffer, &consumerWg)
	}

	// Wait for all producers to finish
	producerWg.Wait()
	close(buffer) // Close the buffer after all producers are done

	// Wait for all consumers to finish
	consumerWg.Wait()

	fmt.Println("All producers and consumers have finished.")
}
