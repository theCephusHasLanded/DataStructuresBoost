package main

import (
    "fmt"
    "go_data_structures_no_generics"
)

func main() {
    // Create a new binary search tree
    bst := go_data_structures_no_generics.NewBinarySearchTree()

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

    // Test Size
    fmt.Println("Size of the tree:", bst.Size()) // Expected: Number of inserted elements

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

    // TODO: Add additional fmt tests for any methods not covered here

	// Create a new deque
    deque := go_data_structures_no_generics.NewDeque()

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

	// TODO: Add additional fmt tests for any methods not covered here
	graph := go_data_structures_no_generics.NewGraph()

    // Test Insert
    fmt.Println("Inserting nodes 1, 2, 3 into the graph.")
    graph.Insert(1)
    graph.Insert(2)
    graph.Insert(3)

    // Test Remove
    fmt.Println("Removing node 2 from the graph.")
    graph.Remove(2)

    // Test Empty
    empty := graph.Empty()
    fmt.Println("Is the graph empty?", empty) // Expected: false

    // Re-insert node for traversal tests
    graph.Insert(2)
    graph.Insert(4)
    graph.edges[1][2] = struct{}{} // Assuming edges are bidirectional for this example
    graph.edges[2][1] = struct{}{}
    graph.edges[2][3] = struct{}{}
    graph.edges[3][2] = struct{}{}
    graph.edges[4][2] = struct{}{}
    graph.edges[2][4] = struct{}{}

    // Test DepthFirstTraversal
    fmt.Println("Depth First Traversal starting from node 1:")
    dft := graph.DepthFirstTraversal(1)
    fmt.Println(dft) // Expected: 1 2 3 4 or a variation depending on edge setup

    // Test BreadthFirstTraversal
    fmt.Println("Breadth First Traversal starting from node 1:")
    bft := graph.BreadthFirstTraversal(1)
    fmt.Println(bft) // Expected: 1 2 3 4 or a variation depending on edge setup

    // Test DepthFirstSearch
    found = graph.DepthFirstSearch(4)
    fmt.Println("Depth First Search for node 4:", found) // Expected: true

    // Test BreadthFirstSearch
    found = graph.BreadthFirstSearch(4)
    fmt.Println("Breadth First Search for node 4:", found) // Expected: true

    // Test Size
    size := graph.Size()
    fmt.Println("Size of the graph (number of nodes):", size) // Expected: 3

    // Removing node 4 to test empty scenario
    graph.Remove(4)
    graph.Remove(3)
    graph.Remove(2)
    graph.Remove(1)

    // Test Empty after removal
    empty = graph.Empty()
    fmt.Println("Is the graph empty after removals?", empty) // Expected: true
	// TODO: Add additional fmt tests for any methods not covered here
	maxHeap := go_data_structures_no_generics.NewMaxHeap()

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
	empty = maxHeap.Empty()
	fmt.Println("Is the max heap empty? (should be false):", empty)

	// Removing remaining elements to test empty scenario
	maxHeap.RemoveTop()
	maxHeap.RemoveTop() // Assuming this is the last element

	// Test Empty after removal
	empty = maxHeap.Empty()
	fmt.Println("Is the max heap empty after removals? (should be true):", empty)
	// TODO: Add additional fmt tests for any methods not covered here
	sll := go_data_structures_no_generics.NewSinglyLinkedList()
    dll := go_data_structures_no_generics.NewDoublyLinkedList()

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
	// TODO: Add additional fmt tests for any methods not covered here
	m := go_data_structures_no_generics.NewMap(10, 0.75)

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
	// TODO: Add additional fmt tests for any methods not covered here
	// Create an instance of Pair from the go_data_structures_no_generics package
	pair := go_data_structures_no_generics.Pair{
		Key:      "exampleKey",
		Priority: "high",
		Value:    "This is the value.",
	}

	// Print out the Pair
	fmt.Println("Pair contents:", pair)
}

// TODO: Add additional fmt tests for any methods not covered here
    pq := go_data_structures_no_generics.NewPriorityQueue()

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
		item := pq.Pop().(go_data_structures_no_generics.Pair) // Make sure to cast the returned interface{} to Pair
		fmt.Printf("Dequeued: %v with priority %s\n", item.Value, item.Priority)
	}

	// Check if queue is empty after dequeuing everything
	fmt.Println("Is the queue empty after dequeuing everything?", pq.Empty())

    // TODO: Add additional fmt tests for any methods not covered here

    queue := go_data_structures_no_generics.NewQueue()
    altQueue := go_data_structures_no_generics.NewAlternateQueue()

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

// TODO: Add additional fmt tests for any methods not covered here

set := go_data_structures_no_generics.NewSet(10) // Initialize set with 10 buckets

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
    // TODO: Add additional fmt tests for any methods not covered here
    stack := go_data_structures_no_generics.NewStack()

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
