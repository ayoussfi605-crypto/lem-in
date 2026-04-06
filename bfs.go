package main

// Bfs performs a breadth-first search to find all possible paths from start to end.
// It uses a visited set per path to allow revisiting nodes in different paths.
func Bfs(farm Farm) [][]string {
	type pathState struct {
		path    []string
		visited map[string]bool
	}

	// Initialize queue with start room
	queue := []pathState{{path: []string{farm.Start}, visited: map[string]bool{farm.Start: true}}}
	var allPaths [][]string
	limit := 500 // Safety limit to prevent excessive computation

	for len(queue) > 0 && len(allPaths) < limit {
		state := queue[0]
		queue = queue[1:]
		last := state.path[len(state.path)-1]

		// If reached end, save the path
		if last == farm.End {
			allPaths = append(allPaths, state.path)
			continue
		}

		// Explore neighbors
		for _, ng := range farm.Adj[last] {
			if !state.visited[ng] {
				// Create new path and visited set
				newPath := make([]string, len(state.path)+1)
				copy(newPath, state.path)
				newPath[len(state.path)] = ng

				newVisited := make(map[string]bool, len(state.visited)+1)
				for node, seen := range state.visited {
					newVisited[node] = seen
				}
				newVisited[ng] = true

				queue = append(queue, pathState{path: newPath, visited: newVisited})
			}
		}
	}
	return allPaths
}
