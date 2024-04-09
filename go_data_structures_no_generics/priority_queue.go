package go_data_structures_no_generics

import (
	"container/heap"
)

type PriorityQueueInterface interface {
    Front() Pair          // Returns the first item in queue. O(1)
    Enqueue(priority int, val interface{}) // Adds val to queue with priority, increases size by 1. O(1)
    Dequeue()              // Removes the highest priority item from the queue, decreases size by 1. O(1)
    Empty() bool           // Returns whether queue is empty. O(1)
    Size() int             // Returns the number of elements in queue. O(1)
}

// PriorityQueue implements a priority queue using a heap.
type PriorityQueue struct {
    items []Pair
}

// Ensure PriorityQueue implements heap.Interface
var _ heap.Interface = (*PriorityQueue)(nil)

// Implement heap.Interface for PriorityQueue.
func (pq *PriorityQueue) Len() int { return len(pq.items) }

func (pq *PriorityQueue) Less(i, j int) bool {
    // We want Pop to give us the highest, not lowest, priority so we use greater than here.
    return pq.items[i].Priority > pq.items[j].Priority
}
func (pq *PriorityQueue) Swap(i, j int) {
    pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(Pair)
    pq.items = append(pq.items, item)
}
func (pq *PriorityQueue) Pop() interface{} {
    old := pq.items
    n := len(old)
    item := old[n-1]
    pq.items = old[:n-1]
    return item
}

// NewPriorityQueue creates a new PriorityQueue.
func NewPriorityQueue() *PriorityQueue {
    pq := &PriorityQueue{}
    heap.Init(pq)
    return pq
}

// Front returns the first item in the queue without removing it.
func (pq *PriorityQueue) Front() Pair {
    if len(pq.items) > 0 {
        return pq.items[0]
    }
    // Assuming a default return if empty; adjust based on usage
    return Pair{}
}

// Enqueue adds an item to the queue with a given priority.
func (pq *PriorityQueue) Enqueue(priority string, val interface{}) {
    heap.Push(pq, Pair{Priority: priority, Value: val})
}

// Dequeue removes and returns the highest priority item from the queue.
func (pq *PriorityQueue) Dequeue() {
    heap.Pop(pq)
}

// Empty returns whether the queue is empty.
func (pq *PriorityQueue) Empty() bool {
    return pq.Len() == 0
}

// Size returns the number of elements in the queue.
func (pq *PriorityQueue) Size() int {
    return pq.Len()
}


