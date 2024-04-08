package go_data_structures_no_generics

// DequeInterface defines the operations for a deque.
type DequeInterface interface {
    Front() interface{}            // Returns the first item in deque. O(1)
    Back() interface{}             // Returns the last item in deque. O(1)
    PushFront(val interface{})     // Adds val to the front of the queue, increases size by 1. O(1)
    PushBack(val interface{})      // Adds val to the back of the queue, increases size by 1. O(1)
    PopFront() (interface{}, bool) // Removes and returns the item from the front of the queue. Returns false if empty.
    PopBack() (interface{}, bool)  // Removes and returns the item from the back of the queue. Returns false if empty.
    Empty() bool                   // Returns whether the queue is empty. O(1)
    Size() int                     // Returns the number of elements in the queue. O(1)
}


// Deque implements a double-ended queue using a doubly linked list.
type Deque struct {
  list *DoublyLinkedList // Adjust DoublyLinkedList to use interface{} as well
}

// NewDeque creates a new Deque instance.
func NewDeque() *Deque {
  return &Deque{list: NewDoublyLinkedList()}
}

func (dq *Deque) PopFront() (interface{}, bool) {
  if dq.Empty() {
      return nil, false // Indicate failure due to empty deque
  }
  return dq.list.PopFront()
}

func (dq *Deque) Front() interface{} {
  if dq.Empty() {
      return nil // Return nil directly as the deque is empty
  }
  value, ok := dq.list.Front()
  if !ok {
      return nil // If for some reason Front() fails, return nil
  }
  return value // Return the value from the front of the deque
}


func (dq *Deque) Back() interface{} {
  if dq.Empty() {
      return nil // Return nil directly as the deque is empty
  }
  value, ok := dq.list.Back()
  if !ok {
      return nil // If for some reason Back() fails, return nil
  }
  return value // Return the value from the back of the deque
}


func (dq *Deque) PushFront(val interface{}) {
  dq.list.InsertAtFront(val)
}

func (dq *Deque) PushBack(val interface{}) {
  dq.list.InsertAtEnd(val)
}

func (dq *Deque) PopBack() (interface{}, bool) {
  if dq.Empty() {
      return nil, false
  }
  return dq.list.PopBack()
}

func (dq *Deque) Empty() bool {
  return dq.list.Empty()
}

func (dq *Deque) Size() int {
  return dq.list.Size()
}
