package go_data_structures_no_generics

import (
    "testing"
	"fmt"
)

func TestLinkedLists(t *testing.T) {

sll := NewSinglyLinkedList()
dll := NewDoublyLinkedList()

// Test Singly Linked List InsertAtFront
fmt.Println("Singly Linked List: Inserting at front")
sll.InsertAtFront(1)
sll.InsertAtFront(2)

// Test Singly Linked List RemoveAtFront
fmt.Println("Singly Linked List: Removing from front")
sll.RemoveAtFront()

// Test Singly Linked List InsertAtEnd
fmt.Println("Singly Linked List: Inserting at end")
sll.InsertAtEnd(3)

// Test Singly Linked List RemoveAtEnd
fmt.Println("Singly Linked List: Removing from end")
sll.RemoveAtEnd()

// Test Singly Linked List Size
fmt.Println("Singly Linked List: Size -", sll.Size())

// Test Doubly Linked List InsertAtFront
fmt.Println("Doubly Linked List: Inserting at front")
dll.InsertAtFront(1)
dll.InsertAtFront(2)

// Test Doubly Linked List RemoveAtFront
fmt.Println("Doubly Linked List: Removing from front")
dll.RemoveAtFront()

// Test Doubly Linked List InsertAtEnd
fmt.Println("Doubly Linked List: Inserting at end")
dll.InsertAtEnd(3)

// Test Doubly Linked List RemoveAtEnd
fmt.Println("Doubly Linked List: Removing from end")
dll.RemoveAtEnd()

// Test Doubly Linked List Size
fmt.Println("Doubly Linked List: Size -", dll.Size())
}
