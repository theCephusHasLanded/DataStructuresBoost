package go_data_structures_no_generics

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
m := NewMap(10, 0.75)

// Test Set
fmt.Println("Setting key-value pairs in the map.")
m.Set("apple", 1)
m.Set("banana", 2)
m.Set("cherry", 3)

// Test Contains
fmt.Println("Contains 'apple':", m.Contains("apple")) // Expected: true
fmt.Println("Contains 'dragonfruit':", m.Contains("dragonfruit")) // Expected: false

// Test Get
value, err := m.Get("banana")
if err != nil {
	fmt.Println("Error retrieving 'banana':", err)
} else {
	fmt.Println("Value for 'banana':", value) // Expected: 2
}

// Test Removes
fmt.Println("Removing 'banana' from the map.")
m.Removes("banana")
fmt.Println("Contains 'banana' after removal:", m.Contains("banana")) // Expected: false

// Test Empty
fmt.Println("Is the map empty?:", m.Empty()) // Expected: false after inserting and removing

// Test Size
fmt.Println("Size of the map:", m.Size()) // Expected: 2 after removing 'banana'

// Test Keys
fmt.Println("Keys in the map:", m.Keys()) // Expected to contain 'apple' and 'cherry'

// Test Values
fmt.Println("Values in the map:", m.Values()) // Expected to contain 1 and 3

// Test Objects
fmt.Println("Objects in the map:", m.Objects()) // Expected to contain Pairs for 'apple' and 'cherry'

}
