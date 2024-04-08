package go_data_structures_no_generics

import (
	"encoding/json"
	"hash/fnv"
)

type SetInterface interface {
    hash(val int) int      // helper function to assign val to an index in the underlying array
    resize()               // helper function to double the size of the underlying array
    Contains(val int) bool // returns whether set contains element val. O(1)
    Insert(val int)        // inserts val in set if not present, increases size by 1. O(1)
    Remove(val int)        // removes item from set if present, decreases size by 1. O(1)
    Empty() bool           // returns whether set is empty. O(1)
    Size() int             // returns number of elements in set. O(1)
    Values() []int         // returns all values in the set.
}


// Node represents a node in a linked list.
type Node struct {
	Value int
	Next  *Node
}

// Bucket is a linked list, used for handling collisions in the hash table.
type Bucket struct {
	head *Node
	size int
}

// NewBucket initializes a new empty bucket.
func NewBucket() *Bucket {
	return &Bucket{}
}

// Insert adds a new value to the bucket if it's not already present.
func (b *Bucket) Insert(val int) bool {
	if b.Contains(val) { // Avoid inserting duplicates.
		return false
	}
	newNode := &Node{Value: val, Next: b.head}
	b.head = newNode
	b.size++
	return true
}

// Contains checks if a value is present in the bucket.
func (b *Bucket) Contains(val int) bool {
	current := b.head
	for current != nil {
		if current.Value == val {
			return true
		}
		current = current.Next
	}
	return false
}

// Remove deletes a value from the bucket, if present.
func (b *Bucket) Remove(val int) bool {
	if b.head == nil {
		return false
	}
	if b.head.Value == val {
		b.head = b.head.Next
		b.size--
		return true
	}
	current := b.head
	for current.Next != nil {
		if current.Next.Value == val {
			current.Next = current.Next.Next
			b.size--
			return true
		}
		current = current.Next
	}
	return false
}

// Size returns the number of elements in the bucket.
func (b *Bucket) Size() int {
	return b.size
}

// Values collects and returns all values stored in the bucket.
func (b *Bucket) Values() []int {
	var values []int
	current := b.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}

// Set represents a set data structure implemented using a hash table.
type Set struct {
	buckets []*Bucket
	size    int
}

// NewSet initializes a new Set with a given number of buckets (for the hash table).
func NewSet(bucketCount int) *Set {
	buckets := make([]*Bucket, bucketCount)
	for i := range buckets {
		buckets[i] = NewBucket()
	}
	return &Set{buckets: buckets}
}

// hash generates a hash code for the given value.
func (s *Set) hash(val int) int {
	valBytes, _ := json.Marshal(val)
	h := fnv.New32a()
	h.Write(valBytes)
	return int(h.Sum32()) % len(s.buckets)
}

// resize doubles the size of the underlying array and rehashes all keys.
func (s *Set) resize() {
	newBuckets := make([]*Bucket, len(s.buckets)*2)
	for i := range newBuckets {
		newBuckets[i] = NewBucket()
	}
	for _, bucket := range s.buckets {
		for _, val := range bucket.Values() {
			newIndex := int(fnv.New32a().Sum32()) % len(newBuckets)
			newBuckets[newIndex].Insert(val)
		}
	}
	s.buckets = newBuckets
}

// Contains returns whether the set contains the given value.
func (s *Set) Contains(val int) bool {
	index := s.hash(val)
	return s.buckets[index].Contains(val)
}

// Insert adds a new value to the set if it's not already present.
func (s *Set) Insert(val int) {
	index := s.hash(val)
	if s.buckets[index].Insert(val) {
		s.size++
		if float32(s.size)/float32(len(s.buckets)) > 0.75 { // Check load factor
			s.resize()
		}
	}
}

// Remove deletes a value from the set, if present.
func (s *Set) Remove(val int) {
	index := s.hash(val)
	if s.buckets[index].Remove(val) {
		s.size--
	}
}

// Empty returns whether the set is empty.
func (s *Set) Empty() bool {
	return s.size == 0
}

// Size returns the number of elements in the set.
func (s *Set) Size() int {
	return s.size
}

// Values returns all values in the set.
func (s *Set) Values() []int {
	var values []int
	for _, bucket := range s.buckets {
		values = append(values, bucket.Values()...)
	}
	return values
}
