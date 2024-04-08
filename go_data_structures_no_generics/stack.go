package go_data_structures



type StackInterface[T any] interface {
	Top() T      // returns top item in stack. O(1)
	Push(val T)  // adds val to top of stack, increases size by 1. O(1)
	Pop()        // removes item from top of stack, decreases size by 1. O(1)
	Empty() bool // returns whether stack is empty. O(1)
	Size() int   // returns number of elements in stack. O(1) or O(n) depending on implementation
}

type Stack struct {
    list SinglyLinkedListInterface // Adapted to work with interface{}
}

func NewStack() *Stack {
    return &Stack{list: NewSinglyLinkedList()} // Assuming NewSinglyLinkedList is adapted to use interface{}
}

func (s *Stack) Top() interface{} {
    return s.list.PeekFront() // Assuming PeekFront returns the element at the front of the list. O(1)
}

func (s *Stack) Push(val interface{}) {
    s.list.PushFront(val) // Assuming PushFront adds an element to the front of the list. O(1)
}

func (s *Stack) Pop() {
    s.list.PopFront() // Assuming PopFront removes the element from the front of the list. O(1)
}

func (s *Stack) Empty() bool {
    return s.list.IsEmpty() // Assuming IsEmpty checks if the list is empty. O(1)
}

func (s *Stack) Size() int {
    return s.list.Count() // Assuming Count returns the number of elements in the list. O(1) or O(n)
}

