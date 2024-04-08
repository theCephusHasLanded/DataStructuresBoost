package go_data_structures_no_generics


type QueueInterface interface {
    Front() interface{}      // Returns the first item in queue. O(1)
    Enqueue(val interface{}) // Adds val to the end of queue, increases size by 1. O(1)
    Dequeue()                // Removes the item from the front of the queue, decreases size by 1. O(1)
    Empty() bool             // Returns whether queue is empty. O(1)
    Size() int               // Returns the number of elements in queue. O(1) or O(n) depending on implementation
}


type Queue struct {
    list SinglyLinkedListInterface
}

func NewQueue() *Queue {
    return &Queue{list: NewSinglyLinkedList()} // Assuming NewSinglyLinkedList is adapted accordingly
}

func (q *Queue) Front() interface{} {
    val, ok := q.list.PeekFront()
    if !ok {
        return nil
    }
    return val
}

func (q *Queue) Enqueue(val interface{}) {
    intValue, ok := val.(int) // Type assertion
    if !ok {
        // Handle the error: val is not of type int
        return
    }
    q.list.PushBack(intValue)
}


func (q *Queue) Dequeue() {
    q.list.PopFront()
}

func (q *Queue) Empty() bool {
    return q.list.IsEmpty()
}

func (q *Queue) Size() int {
    return q.list.Count()
}

type AlternateQueue struct {
    arr []interface{}
}

func NewAlternateQueue() *AlternateQueue {
    return &AlternateQueue{arr: make([]interface{}, 0)}
}

func (q *AlternateQueue) Front() interface{} {
    if len(q.arr) == 0 {
        return nil // Return nil if queue is empty
    }
    return q.arr[0]
}

func (q *AlternateQueue) Enqueue(val interface{}) {
    q.arr = append(q.arr, val)
}

func (q *AlternateQueue) Dequeue() {
    if len(q.arr) > 0 {
        q.arr = q.arr[1:]
    }
}

func (q *AlternateQueue) Empty() bool {
    return len(q.arr) == 0
}

func (q *AlternateQueue) Size() int {
    return len(q.arr)
}
