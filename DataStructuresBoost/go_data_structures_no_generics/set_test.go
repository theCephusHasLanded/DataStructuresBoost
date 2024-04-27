package go_data_structures_no_generics

import (
    "testing"
	"fmt"
)

func TestSet(t *testing.T) {
set := NewSet(10) // Initialize set with 10 buckets

// Test Insert
fmt.Println("Inserting values into the set")
set.Insert(10)
set.Insert(20)
set.Insert(30)
set.Insert(40)
set.Insert(10) // Duplicate, should not increase the size

// Test Contains
fmt.Println("Does the set contain 10?", set.Contains(10)) // Expected: true
fmt.Println("Does the set contain 50?", set.Contains(50)) // Expected: false

// Test Remove
fmt.Println("Removing 20 from the set")
set.Remove(20)
fmt.Println("Does the set contain 20 after removal?", set.Contains(20)) // Expected: false

// Test Empty
fmt.Println("Is the set empty?", set.Empty()) // Expected: false

// Test Size
fmt.Println("Size of the set:", set.Size()) // Expected: 3, after one removal

// Test Values
fmt.Println("Values in the set:", set.Values())

// Remove all items
set.Remove(10)
set.Remove(30)
set.Remove(40)
fmt.Println("Is the set empty after removing all items?", set.Empty()) // Expected: true

// Re-insert and test resize
for i := 0; i < 50; i++ {
	set.Insert(i)
}
fmt.Println("Size of the set after inserting 50 items:", set.Size()) // Expected: 50
fmt.Println("Testing resize by checking new values presence:")
fmt.Println("Does the set contain 25?", set.Contains(25)) // Expected: true
}
