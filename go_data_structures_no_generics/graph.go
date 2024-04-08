package go_data_structures_no_generics

// Graph represents a graph with nodes of type int.
type Graph struct {
	nodes map[int]struct{}       // Set of nodes to ensure uniqueness.
	edges map[int]map[int]struct{} // Adjacency list to represent edges.
}

// NewGraph creates a new instance of a graph.
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]struct{}),
		edges: make(map[int]map[int]struct{}),
	}
}

// insert adds a node to the graph and initializes its adjacency list if not present.
func (g *Graph) Insert(val int) {
	if _, exists := g.nodes[val]; !exists {
		g.nodes[val] = struct{}{}
		g.edges[val] = make(map[int]struct{})
	}
}

// remove deletes a node from the graph and removes it from other nodes' adjacency lists.
func (g *Graph) Remove(val int) {
	if _, exists := g.nodes[val]; exists {
		delete(g.nodes, val)
		delete(g.edges, val)
		for node := range g.edges {
			delete(g.edges[node], val)
		}
	}
}

// Empty checks if the graph is empty.
func (g *Graph) Empty() bool {
	return len(g.nodes) == 0
}

// DepthFirstTraversal performs a depth-first traversal starting from root.
func (g *Graph) DepthFirstTraversal(root int) []int {
	var result []int
	visited := make(map[int]bool)
	var dfs func(int)
	dfs = func(node int) {
		if visited[node] {
			return
		}
		visited[node] = true
		result = append(result, node)
		for neighbor := range g.edges[node] {
			dfs(neighbor)
		}
	}
	dfs(root)
	return result
}

// BreadthFirstTraversal performs a breadth-first traversal starting from root.
func (g *Graph) BreadthFirstTraversal(root int) []int {
	var result []int
	visited := make(map[int]bool)
	queue := []int{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if visited[node] {
			continue
		}
		visited[node] = true
		result = append(result, node)
		for neighbor := range g.edges[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
	return result
}

// DepthFirstSearch searches for a value in the graph using depth-first search.
func (g *Graph) DepthFirstSearch(val int) bool {
	visited := make(map[int]bool)
	var dfs func(int) bool
	dfs = func(node int) bool {
		if node == val {
			return true
		}
		if visited[node] {
			return false
		}
		visited[node] = true
		for neighbor := range g.edges[node] {
			if dfs(neighbor) {
				return true
			}
		}
		return false
	}

	for node := range g.nodes {
		if !visited[node] && dfs(node) {
			return true
		}
	}
	return false
}

// BreadthFirstSearch searches for a value in the graph using breadth-first search.
func (g *Graph) BreadthFirstSearch(val int) bool {
	visited := make(map[int]bool)
	queue := []int{}
	for node := range g.nodes {
		queue = append(queue, node) // Start search from every node
		break                        // But actually start from the first node only, can be modified if needed
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == val {
			return true
		}
		if visited[node] {
			continue
		}
		visited[node] = true
		for neighbor := range g.edges[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
	return false
}

// Size returns the number of nodes in the graph.
func (g *Graph) Size() int {
	return len(g.nodes)
}
