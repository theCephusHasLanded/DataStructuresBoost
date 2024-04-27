package go_data_structures_no_generics

import (
    "testing"
	"fmt"
)

func TestHeap(t *testing.T) {

maxHeap :=NewMaxHeap()

// Test Insert
fmt.Println("Inserting 42, 29, 18 into the max heap.")
maxHeap.Insert(42)
maxHeap.Insert(29)
maxHeap.Insert(18)

// Test Peek
top := maxHeap.Peek()
fmt.Println("Top of the max heap (should be 42):", top)

// Test Size
fmt.Println("Size of the max heap (should be 3):", maxHeap.Size())

// Test RemoveTop
fmt.Println("Removing the top element (should remove 42).")
maxHeap.RemoveTop()
fmt.Println("New top of the max heap (should be 29):", maxHeap.Peek())

// Test Sorted (Heap Sort)
fmt.Println("Getting sorted elements (should not alter the heap structure):")
sorted := maxHeap.Sorted()
fmt.Println("Sorted elements:", sorted)

// Test that Sorted did not alter the original heap
fmt.Println("Top of the max heap after sorting (should be 29):", maxHeap.Peek())

// Test Remove specific element
fmt.Println("Inserting 15 into the max heap.")
maxHeap.Insert(15)
fmt.Println("Removing 15 from the max heap.")
maxHeap.Remove(15)
fmt.Println("Top of the max heap after removing 15 (should be 29):", maxHeap.Peek())
fmt.Println("Size of the max heap after removing 15 (should be 2):", maxHeap.Size())

// Test Empty
empty := maxHeap.Empty()
fmt.Println("Is the max heap empty? (should be false):", empty)

// Removing remaining elements to test empty scenario
maxHeap.RemoveTop()
maxHeap.RemoveTop() // Assuming this is the last element

// Test Empty after removal
empty = maxHeap.Empty()
fmt.Println("Is the max heap empty after removals? (should be true):", empty)

}
