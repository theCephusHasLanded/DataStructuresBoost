package go_data_structures_no_generics

import (
    "encoding/json"
    "fmt"
    "hash/fnv"
)


// MapInterface defines operations for a map data structure.
type MapInterface interface {
    Contains(key string) bool
    Get(key string) (interface{}, error)
    Set(key string, val interface{})
    Removes(key string)
    Empty() bool
    Size() int
    Values() []interface{}
    Keys() []string
    Objects() []Pair
}

// Map implements a map using a slice of binary search trees for collision resolution.
type Map struct {
    arr     []*BinarySearchTree
    maxFill float32
    size    int
}

// NewMap creates a new Map instance.
func NewMap(size int, maxFill float32) *Map {
    m := &Map{
        arr:     make([]*BinarySearchTree, size),
        maxFill: maxFill,
    }
    for i := range m.arr {
        m.arr[i] = NewBinarySearchTree()
    }
    return m
}

// hash generates a hash code for the given key.
func (m *Map) hash(key string) int {
    keyStr, _ := json.Marshal(key)
    h := fnv.New32a()
    h.Write(keyStr)
    return int(h.Sum32()) % len(m.arr)
}

// resize doubles the size of the underlying array and rehashes all keys.
func (m *Map) Resize() {
    newArrSize := len(m.arr) * 2
    newArr := make([]*BinarySearchTree, newArrSize)
    for _, tree := range m.arr {
        if tree != nil {
            for _, pair := range tree.Objects() {
                newIndex := m.hash(pair.Key)
                if newArr[newIndex] == nil {
                    newArr[newIndex] = NewBinarySearchTree()
                }
                newArr[newIndex].Insert(pair.Key, pair.Value)
            }
        }
    }
    m.arr = newArr
}

// Contains checks if the map contains the specified key.
func (m *Map) Contains(key string) bool {
    index := m.hash(key)
    if m.arr[index] == nil {
        return false
    }
    return m.arr[index].Contains(key)
}

// Get retrieves the value associated with the given key.
func (m *Map) Get(key string) (interface{}, error) {
    index := m.hash(key)
    if m.arr[index] == nil {
        return nil, fmt.Errorf("key not found")
    }
    value, found := m.arr[index].Get(key)
    if !found {
        return nil, fmt.Errorf("key not found")
    }
    return value, nil // Successfully found the value
}


func (m *Map) Set(key string, val interface{}) {
    index := m.hash(key)
    if m.arr[index] == nil {
        m.arr[index] = NewBinarySearchTree()
    }
    // Check if the key already exists to decide on incrementing size
    if !m.arr[index].Contains(key) {
        m.size++
    }
    m.arr[index].Insert(key, val)
}

// Removes deletes the key-val pair from the map.
func (m *Map) Removes(key string) {
    index := m.hash(key)
    if m.arr[index] != nil && m.arr[index].Remove(key) {
        m.size--
    }
}

// Empty checks if the map is empty.
func (m *Map) Empty() bool {
    return m.size == 0
}

// Size returns the number of elements in the map.
func (m *Map) Size() int {
    return m.size
}

// Values returns all values in the map.
func (m *Map) Values() []interface{} {
    var values []interface{}
    for _, tree := range m.arr {
        if tree != nil {
            values = append(values, tree.Values()...)
        }
    }
    return values
}

// Keys returns all keys in the map.
func (m *Map) Keys() []string {
    var keys []string
    for _, tree := range m.arr {
        if tree != nil {
            keys = append(keys, tree.Keys()...)
        }
    }
    return keys
}

// Objects returns all key-value pairs in the map.
func (m *Map) Objects() []Pair {
    var objects []Pair
    for _, tree := range m.arr {
        if tree != nil {
            objects = append(objects, tree.Objects()...)
        }
    }
    return objects
}
