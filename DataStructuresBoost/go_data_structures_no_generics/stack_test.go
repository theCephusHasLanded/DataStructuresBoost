package go_data_structures_no_generics

import (
    "testing"
	"fmt"
)

func TestStack(t *testing.T) {

    stack := NewStack()
// Test Push
fmt.Println("Pushing items onto the stack")
stack.Push("Hello")
stack.Push(42)
stack.Push([3]int{1, 2, 3})

// Test Top
topItem := stack.Top()
fmt.Println("Item at the top of the stack:", topItem) // Expected: [3]int{1, 2, 3}

// Test Size
fmt.Println("Current size of the stack:", stack.Size()) // Expected: 3

// Test Empty
fmt.Println("Is the stack empty?", stack.Empty()) // Expected: false

// Test Pop
fmt.Println("Popping items from the stack:")
for !stack.Empty() {
	top := stack.Top()
	stack.Pop()
	fmt.Println("Popped:", top)
	fmt.Println("New top after pop:", stack.Top())
}

// Check if stack is empty after popping all items
fmt.Println("Is the stack empty after all pops?", stack.Empty()) // Expected: true

// Test Empty and Size after clearing stack
fmt.Println("Final check - Is the stack empty?", stack.Empty()) // Expected: true
fmt.Println("Final check - Size of the stack:", stack.Size())   // Expected: 0
}
