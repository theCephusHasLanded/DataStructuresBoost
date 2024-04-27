package go_data_structures_no_generics

import (
    "testing"
	"fmt"
)

func TestDeque(t *testing.T) {

	// Create a new deque
    deque := NewDeque()

    // Test PushFront
    fmt.Println("Pushing 10 to the front...")
    deque.PushFront(10)

    // Test Front
    fmt.Println("The front item of deque:", deque.Front()) // Expected: 10

    // Test PushBack
    fmt.Println("Pushing 20 to the back...")
    deque.PushBack(20)

    // Test Back
    fmt.Println("The back item of deque:", deque.Back()) // Expected: 20

    // Test PopFront
    frontValue, frontOk := deque.PopFront()
    fmt.Printf("Popped from front: %v, Successful: %v\n", frontValue, frontOk) // Expected: 10, true

    // Test PopBack
    backValue, backOk := deque.PopBack()
    fmt.Printf("Popped from back: %v, Successful: %v\n", backValue, backOk) // Expected: 20, true

    // Test Empty
    fmt.Println("Is the deque empty?:", deque.Empty()) // Expected: true after popping both items

    // Test Size
    fmt.Println("Size of deque:", deque.Size()) // Expected: 0 after popping both items

    // Additional Push and Pop operations to test behavior
    fmt.Println("Pushing more items...")
    deque.PushFront(30)
    deque.PushBack(40)
    deque.PushFront(50)

    fmt.Println("Current front item:", deque.Front()) // Expected: 50
    fmt.Println("Current back item:", deque.Back())   // Expected: 40
    fmt.Println("Current size:", deque.Size())        // Expected: 3

    // Pop operations
    fmt.Println("Popping items...")
    deque.PopFront()
    deque.PopBack()

    fmt.Println("Current front item after pops:", deque.Front()) // Expected: 30
    fmt.Println("Current size after pops:", deque.Size())        // Expected: 1
    fmt.Println("Is deque empty after pops?:", deque.Empty())    // Expected: false

    // Final Pop to clear the deque
    deque.PopFront()
    fmt.Println("Is deque empty after final pop?:", deque.Empty())
}
