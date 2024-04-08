package go_data_structures_no_generics

// SinglyLinkedListInterface defines the operations for a singly linked list of integers.
type SinglyLinkedListInterface interface {
    InsertAfter(val int, prev *SingleLinkNode)
    RemoveAfter(prev *SingleLinkNode)
    InsertAtFront(val int)
    RemoveAtFront()
    InsertAtEnd(val int)
    RemoveAtEnd()
    Head() *SingleLinkNode
    Tail() *SingleLinkNode
    Empty() bool
    Size() int
	PushFront(val interface{})
    PeekFront() (interface{}, bool) // Assuming it returns the first item and its existence.
    PushBack(val int)       // Method to add an item at the end.
    PopFront() (int, bool)  // Method to remove and return the first item.
    IsEmpty() bool          // Method to check if the list is empty.
    Count() int             // Method to get the number of items.
}


// SingleLinkNode represents a node in a singly linked list.
type SingleLinkNode struct {
	Data int
	Next *SingleLinkNode
}

// SinglyLinkedList represents a singly linked list of integers.
type SinglyLinkedList struct {
	head *SingleLinkNode
	tail *SingleLinkNode
	size int
}

// NewSinglyLinkedList initializes a new empty singly linked list.
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}


// InsertAfter inserts a new node with the given value after the specified predecessor node.
func (sll *SinglyLinkedList) InsertAfter(val int, prev *SingleLinkNode) {
	newNode := &SingleLinkNode{Data: val, Next: prev.Next}
	prev.Next = newNode
	if prev.Next == nil {
		sll.tail = newNode
	}
}

// RemoveAfter removes the node immediately following the specified predecessor node.
func (sll *SinglyLinkedList) RemoveAfter(prev *SingleLinkNode) {
	if prev == nil || prev.Next == nil {
		return
	}
	removing := prev.Next
	prev.Next = removing.Next
	if removing.Next == nil {
		sll.tail = prev
	}
}

// InsertAtFront adds a new node with the given value to the start of the list.
func (sll *SinglyLinkedList) InsertAtFront(val int) {
	newNode := &SingleLinkNode{Data: val, Next: sll.head}
	sll.head = newNode
	if sll.tail == nil {
		sll.tail = newNode
	}
}

// RemoveAtFront removes the first node of the list.
func (sll *SinglyLinkedList) RemoveAtFront() {
	if sll.head == nil {
		return
	}
	sll.head = sll.head.Next
	if sll.head == nil {
		sll.tail = nil
	}
}

// InsertAtEnd adds a new node with the given value to the end of the list.
func (sll *SinglyLinkedList) InsertAtEnd(val int) {
	newNode := &SingleLinkNode{Data: val, Next: nil}
	if sll.tail == nil {
		sll.head = newNode
		sll.tail = newNode
		return
	}
	sll.tail.Next = newNode
	sll.tail = newNode
}

// RemoveAtEnd removes the last node of the list.
func (sll *SinglyLinkedList) RemoveAtEnd() {
	if sll.head == nil {
		return
	}
	if sll.head.Next == nil {
		sll.head = nil
		sll.tail = nil
		return
	}
	curr := sll.head
	for curr.Next != sll.tail {
		curr = curr.Next
	}
	curr.Next = nil
	sll.tail = curr
}

// Head returns the first node of the list.
func (sll *SinglyLinkedList) Head() *SingleLinkNode {
	return sll.head
}

// Tail returns the last node of the list.
func (sll *SinglyLinkedList) Tail() *SingleLinkNode {
	return sll.tail
}

// Empty checks if the list is empty.
func (sll *SinglyLinkedList) Empty() bool {
	return sll.head == nil
}

func (sll *SinglyLinkedList) PushFront(val interface{}) {
	valInt, ok := val.(int)
	if !ok {
		// Handle the error if val is not of type int
		panic("val is not an int") // Or return an error, depending on your function signature
	}
	newNode := &SingleLinkNode{Data: valInt, Next: nil}
    sll.head = newNode
    // Update any other necessary state (e.g., size if tracked)
}


func (sll *SinglyLinkedList) PeekFront() (interface{}, bool) {
    if sll.head == nil {
        return nil, false
    }
    return sll.head.Data, true
}


func (sll *SinglyLinkedList) PushBack(val int) {
    sll.InsertAtEnd(val) // Utilize your existing InsertAtEnd method.
}

func (sll *SinglyLinkedList) PopFront() (int, bool) {
    if sll.head == nil {
        return 0, false
    }
    val := sll.head.Data
    sll.RemoveAtFront() // Utilize your existing RemoveAtFront method.
    return val, true
}

func (sll *SinglyLinkedList) IsEmpty() bool {
    return sll.Empty() // Utilize your existing Empty method.
}

func (sll *SinglyLinkedList) Size() int {
    return sll.size // Utilize your existing Size method.
}

func (sll *SinglyLinkedList) Count() int {
    return sll.Size() // Assuming you have a Size method that returns the list size
}


// DoublyLinkedListInterface defines the operations for a doubly linked list of integers.
type DoublyLinkedListInterface interface {
	InsertAtFront(val interface{})                     // Adds a new node with the given value to the start of the list.
	InsertAtEnd(val interface{})                       // Adds a new node with the given value to the end of the list.
	RemoveAtFront()                            // Removes the first node of the list.
	RemoveAtEnd()                              // Removes the last node of the list.
	Head() *DoubleLinkNode                     // Returns the first node of the list.
	Tail() *DoubleLinkNode                     // Returns the last node of the list.
	Empty() bool                               // Checks if the list is empty.
	Size() int                                 // Returns the number of nodes in the list.
}

// DoubleLinkNode represents a node in the doubly linked list.
type DoubleLinkNode struct {
	Value interface{}
	Next  *DoubleLinkNode
	Prev  *DoubleLinkNode
}

// DoublyLinkedList represents a doubly linked list.
type DoublyLinkedList struct {
	head *DoubleLinkNode
	tail *DoubleLinkNode
	size int
}

// NewDoublyLinkedList initializes a new doubly linked list.
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// InsertAtFront adds a new node with the given value to the start of the list.
func (dll *DoublyLinkedList) InsertAtFront(val interface{}) {
	newNode := &DoubleLinkNode{Value: val, Next: dll.head}
	if dll.head != nil {
		dll.head.Prev = newNode
	} else {
		dll.tail = newNode // If the list was empty
	}
	dll.head = newNode
	dll.size++
}

// InsertAtEnd adds a new node with the given value to the end of the list.
func (dll *DoublyLinkedList) InsertAtEnd(val interface{}) {
	newNode := &DoubleLinkNode{Value: val, Prev: dll.tail}
	if dll.tail != nil {
		dll.tail.Next = newNode
	} else {
		dll.head = newNode // If the list was empty
	}
	dll.tail = newNode
	dll.size++
}

// RemoveAtFront removes the first node of the list.
func (dll *DoublyLinkedList) RemoveAtFront() {
	if dll.head == nil {
		return
	}
	dll.head = dll.head.Next
	if dll.head != nil {
		dll.head.Prev = nil
	} else {
		dll.tail = nil // If the list becomes empty
	}
	dll.size--
}

// RemoveAtEnd removes the last node of the list.
func (dll *DoublyLinkedList) RemoveAtEnd() {
	if dll.tail == nil {
		return
	}
	dll.tail = dll.tail.Prev
	if dll.tail != nil {
		dll.tail.Next = nil
	} else {
		dll.head = nil // If the list becomes empty
	}
	dll.size--
}

// Head returns the first node of the list.
func (dll *DoublyLinkedList) Head() *DoubleLinkNode {
	return dll.head
}

// Tail returns the last node of the list.
func (dll *DoublyLinkedList) Tail() *DoubleLinkNode {
	return dll.tail
}

// Empty checks if the list is empty.
func (dll *DoublyLinkedList) Empty() bool {
	return dll.head == nil
}

// Size returns the number of nodes in the list.
func (dll *DoublyLinkedList) Size() int {
	return dll.size
}


func (dll *DoublyLinkedList) PopFront() (interface{}, bool) {
    if dll.head == nil {
        return nil, false
    }
    val := dll.head.Value
    dll.head = dll.head.Next
    if dll.head != nil {
        dll.head.Prev = nil
    } else {
        dll.tail = nil
    }
    dll.size--
    return val, true
}

func (dll *DoublyLinkedList) PopBack() (interface{}, bool) {
    if dll.tail == nil {
        return nil, false
    }
    val := dll.tail.Value
    dll.tail = dll.tail.Prev
    if dll.tail != nil {
        dll.tail.Next = nil
    } else {
        dll.head = nil
    }
    dll.size--
    return val, true
}

func (dll *DoublyLinkedList) Front() (interface{}, bool) {
    if dll.head == nil {
        return nil, false
    }
    return dll.head.Value, true
}

func (dll *DoublyLinkedList) Back() (interface{}, bool) {
    if dll.tail == nil {
        return nil, false
    }
    return dll.tail.Value, true
}
