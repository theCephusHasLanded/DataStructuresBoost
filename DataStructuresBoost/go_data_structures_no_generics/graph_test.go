package go_data_structures_no_generics

import (
    "testing"
	"fmt"
)

func TestGraph(t *testing.T) {

	graph := NewGraph()

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
    found := graph.DepthFirstSearch(4)
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
}
