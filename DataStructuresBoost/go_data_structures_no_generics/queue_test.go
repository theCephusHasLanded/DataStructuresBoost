package go_data_structures_no_generics

import (
    "testing"
	"fmt"
)

func TestQueue(t *testing.T) {

	queue := NewQueue()
	altQueue := NewAlternateQueue()

	// Testing standard Queue
	fmt.Println("Testing standard Queue")
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	fmt.Println("Front of queue:", queue.Front()) // Expected: 10
	fmt.Println("Queue size:", queue.Size())      // Expected: 3
	queue.Dequeue()
	fmt.Println("Front of queue after one dequeue:", queue.Front()) // Expected: 20
	fmt.Println("Queue size after one dequeue:", queue.Size())      // Expected: 2
	fmt.Println("Is queue empty?", queue.Empty())                   // Expected: false
	queue.Dequeue()
	queue.Dequeue()
	fmt.Println("Is queue empty after clearing?", queue.Empty())    // Expected: true

	// Testing AlternateQueue
	fmt.Println("\nTesting AlternateQueue")
	altQueue.Enqueue("apple")
	altQueue.Enqueue("banana")
	altQueue.Enqueue("cherry")
	fmt.Println("Front of alternate queue:", altQueue.Front()) // Expected: "apple"
	fmt.Println("Alternate queue size:", altQueue.Size())      // Expected: 3
	altQueue.Dequeue()
	fmt.Println("Front of alternate queue after one dequeue:", altQueue.Front()) // Expected: "banana"
	fmt.Println("Alternate queue size after one dequeue:", altQueue.Size())      // Expected: 2
	fmt.Println("Is alternate queue empty?", altQueue.Empty())                   // Expected: false
	altQueue.Dequeue()
	altQueue.Dequeue()
	fmt.Println("Is alternate queue empty after clearing?", altQueue.Empty())    // Expected: true
	}
