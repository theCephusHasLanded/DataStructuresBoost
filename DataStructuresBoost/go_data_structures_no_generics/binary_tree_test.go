package go_data_structures_no_generics

import (
    "testing"
	"fmt"
)

func TestBinaryTree(t *testing.T) {
// Create a new binary search tree
bst := NewBinarySearchTree()

// Test Insert
fmt.Println("Inserting values...")
bst.Insert("apple", 1)
bst.Insert("banana", 2)
bst.Insert("cherry", 3)

// Test Contains
fmt.Println("Contains 'banana':", bst.Contains("banana")) // Expected: true
fmt.Println("Contains 'durian':", bst.Contains("durian")) // Expected: false

// Test Remove
fmt.Println("Removing 'banana'...")
bst.Remove("banana")
fmt.Println("Contains 'banana' after removal:", bst.Contains("banana")) // Expected: false

// Test Empty
fmt.Println("Is the tree empty?:", bst.Empty()) // Expected: false

// // Test Size
// fmt.Println("Size of the tree:", bst.Size()) // Expected: Number of inserted elements

// Test InOrderTraversal
fmt.Println("InOrder Traversal:", bst.InOrderTraversal())

// Test PreOrderTraversal
fmt.Println("PreOrder Traversal:", bst.PreOrderTraversal())

// Test PostOrderTraversal
fmt.Println("PostOrder Traversal:", bst.PostOrderTraversal())

// Test LevelOrderTraversal
fmt.Println("LevelOrder Traversal:", bst.LevelOrderTraversal())

// Test Get
value, found := bst.Get("apple")
fmt.Printf("Get 'apple': %v, Found: %v\n", value, found)

// Test Objects
fmt.Println("Objects in the tree:", bst.Objects())

// Test Values
fmt.Println("Values in the tree:", bst.Values())

// Test Keys
fmt.Println("Keys in the tree:", bst.Keys())
}
