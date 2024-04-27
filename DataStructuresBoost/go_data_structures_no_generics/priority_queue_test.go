package go_data_structures_no_generics

import (
    "testing"
	"fmt"
)

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()

	// Test Enqueue
	fmt.Println("Enqueuing items with priority")
	pq.Enqueue("high", "Pay bills")
	pq.Enqueue("medium", "Read book")
	pq.Enqueue("low", "Watch TV")

	// Test Front
	front := pq.Front()
	fmt.Printf("Front item: %v with priority %s\n", front.Value, front.Priority)

	// Test Size
	fmt.Println("Size of the queue:", pq.Size())

	// Test Empty
	fmt.Println("Is the queue empty?", pq.Empty())

	// Test Dequeue by emptying the queue
	fmt.Println("Dequeuing all items:")
	for !pq.Empty() {
		item := pq.Pop().(Pair) // Make sure to cast the returned interface{} to Pair
		fmt.Printf("Dequeued: %v with priority %s\n", item.Value, item.Priority)
	}

	// Check if queue is empty after dequeuing everything
	fmt.Println("Is the queue empty after dequeuing everything?", pq.Empty())
	}
