package go_data_structures_no_generics

// BinaryTreeNode represents a node in the binary search tree.
type BinaryTreeNode struct {
    Data  Pair
    Left  *BinarySearchTreeNode
    Right *BinarySearchTreeNode
}

// BinaryTreeInterface defines the operations for a binary tree.
type BinaryTreeInterface interface {
    Root() *BinarySearchTreeNode         // Adjusted to return the correct node type
    Insert(key string, val interface{})  // Key type corrected to string
    Contains(key string) bool            // Key type corrected to string
    Remove(key string)                   // Key type corrected to string
    Empty() bool
    InOrderTraversal() []Pair
    PreOrderTraversal() []Pair
    PostOrderTraversal() []Pair
    LevelOrderTraversal() []Pair
    Get(key string) (interface{}, bool)  // Added to fetch a value by its key
    Objects() []Pair                      // Added to get all key-value pairs
    Values() []interface{}                // Added to get all values
    Keys() []string                       // Added to get all keys
}


// BinarySearchTree implements a binary search tree.
type BinarySearchTree struct {
    root *BinarySearchTreeNode
    size int
}

type BinarySearchTreeNode struct {
    Key   int
    Data  Pair
    Left  *BinarySearchTreeNode
    Right *BinarySearchTreeNode
}


// NewBinarySearchTree creates a new instance of a binary search tree.
func NewBinarySearchTree() *BinarySearchTree {
    return &BinarySearchTree{}
}

// Insert inserts a new key-value pair into the binary search tree.
func (bst *BinarySearchTree) Insert(key string, value interface{}) {
    bst.root = bst.insertNode(bst.root, Pair{Key: key, Value: value})
}

func (bst *BinarySearchTree) insertNode(node *BinarySearchTreeNode, data Pair) *BinarySearchTreeNode {
    if node == nil {
        return &BinarySearchTreeNode{Data: data}
    }
    if data.Key < node.Data.Key {
        node.Left = bst.insertNode(node.Left, data)
    } else if data.Key > node.Data.Key {
        node.Right = bst.insertNode(node.Right, data)
    } // else: Key already exists, update the value or ignore based on specific requirements
    return node
}

func insertNode(node *BinarySearchTreeNode, pair Pair) *BinarySearchTreeNode {
    if node == nil {
        return &BinarySearchTreeNode{Data: pair}
    }
    // Assume string keys should be ordered lexicographically
    if pair.Key < node.Data.Key {
        node.Left = insertNode(node.Left, pair)
    } else if pair.Key > node.Data.Key {
        node.Right = insertNode(node.Right, pair)
    }
    // Note: This example does not handle the case where pair.Key == node.Data.Key
    return node
}


// Contains checks if a key exists in the binary search tree.
func (bst *BinarySearchTree) Contains(key string) bool {
    return bst.containsNode(bst.root, key)
}

func (bst *BinarySearchTree) containsNode(node *BinarySearchTreeNode, key string) bool {
    if node == nil {
        return false
    }
    if key == node.Data.Key {
        return true
    }
    if key < node.Data.Key {
        return bst.containsNode(node.Left, key)
    }
    return bst.containsNode(node.Right, key)
}


// Remove removes a node with the specified key from the binary search tree.
func (bst *BinarySearchTree) Remove(key string) bool {
    var wasRemoved bool
    bst.root, wasRemoved = bst.removeNode(bst.root, key)
    return wasRemoved
}

func (bst *BinarySearchTree) removeNode(node *BinarySearchTreeNode, key string) (*BinarySearchTreeNode, bool) {
    if node == nil {
        return nil, false // Node not found, nothing removed
    }
    var wasRemoved bool
    if key < node.Data.Key {
        node.Left, wasRemoved = bst.removeNode(node.Left, key)
    } else if key > node.Data.Key {
        node.Right, wasRemoved = bst.removeNode(node.Right, key)
    } else {
        // Key found, remove this node
        wasRemoved = true
        if node.Left == nil {
            return node.Right, wasRemoved
        }
        if node.Right == nil {
            return node.Left, wasRemoved
        }
        // Node has two children, find inorder successor
        minRightNode := bst.findMinNode(node.Right)
        node.Data = minRightNode.Data
        node.Right, _ = bst.removeNode(node.Right, minRightNode.Data.Key)
    }
    return node, wasRemoved
}

func (bst *BinarySearchTree) findMinNode(node *BinarySearchTreeNode) *BinarySearchTreeNode {
    current := node
    for current.Left != nil {
        current = current.Left
    }
    return current
}


// Empty checks if the binary search tree is empty.
func (bst *BinarySearchTree) Empty() bool {
    return bst.size == 0
}

// Root returns the root node of the binary search tree.
func (bst *BinarySearchTree) Root() *BinarySearchTreeNode {
    return bst.root
}

// Traversal methods: InOrderTraversal, PreOrderTraversal, PostOrderTraversal, LevelOrderTraversal.
// These methods would be implemented to traverse the tree in various orders and collect pairs into slices.


func traverse(node *BinarySearchTreeNode, order string, result *[]Pair) {
    if node == nil {
        return
    }
    switch order {
    case "inorder":
        traverse(node.Left, "inorder", result)
        *result = append(*result, node.Data)
        traverse(node.Right, "inorder", result)
    case "preorder":
        *result = append(*result, node.Data)
        traverse(node.Left, "preorder", result)
        traverse(node.Right, "preorder", result)
    case "postorder":
        traverse(node.Left, "postorder", result)
        traverse(node.Right, "postorder", result)
        *result = append(*result, node.Data)
    }
}

func (bst *BinarySearchTree) InOrderTraversal() []Pair {
    var result []Pair
    traverse(bst.root, "inorder", &result)
    return result
}

func (bst *BinarySearchTree) PreOrderTraversal() []Pair {
    var result []Pair
    traverse(bst.root, "preorder", &result)
    return result
}

func (bst *BinarySearchTree) PostOrderTraversal() []Pair {
    var result []Pair
    traverse(bst.root, "postorder", &result)
    return result
}

func (bst *BinarySearchTree) LevelOrderTraversal() []Pair {
    var result []Pair
    if bst.root == nil {
        return result
    }
    queue := []*BinarySearchTreeNode{bst.root}
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        result = append(result, current.Data)
        if current.Left != nil {
            queue = append(queue, current.Left)
        }
        if current.Right != nil {
            queue = append(queue, current.Right)
        }
    }
    return result
}

func (bst *BinarySearchTree) Objects() []Pair {
    var objects []Pair
    bst.inOrder(bst.root, &objects)
    return objects
}

func (bst *BinarySearchTree) inOrder(node *BinarySearchTreeNode, objects *[]Pair) {
    if node == nil {
        return
    }
    bst.inOrder(node.Left, objects)
    *objects = append(*objects, node.Data)
    bst.inOrder(node.Right, objects)
}

func (bst *BinarySearchTree) Values() []interface{} {
    var values []interface{}
    bst.inOrderValues(bst.root, &values)
    return values
}

func (bst *BinarySearchTree) inOrderValues(node *BinarySearchTreeNode, values *[]interface{}) {
    if node == nil {
        return
    }
    bst.inOrderValues(node.Left, values)
    *values = append(*values, node.Data.Value)
    bst.inOrderValues(node.Right, values)
}

func (bst *BinarySearchTree) Keys() []string {
    var keys []string
    bst.inOrderKeys(bst.root, &keys)
    return keys
}

func (bst *BinarySearchTree) inOrderKeys(node *BinarySearchTreeNode, keys *[]string) {
    if node == nil {
        return
    }
    bst.inOrderKeys(node.Left, keys)
    *keys = append(*keys, node.Data.Key)
    bst.inOrderKeys(node.Right, keys)
}

func (bst *BinarySearchTree) Get(key string) (interface{}, bool) {
    node := bst.getNode(bst.root, key)
    if node != nil {
        return node.Data.Value, true
    }
    return nil, false
}

func (bst *BinarySearchTree) getNode(node *BinarySearchTreeNode, key string) *BinarySearchTreeNode {
    if node == nil || node.Data.Key == key {
        return node
    }
    if key < node.Data.Key {
        return bst.getNode(node.Left, key)
    }
    return bst.getNode(node.Right, key)
}
