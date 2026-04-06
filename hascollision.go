package main

// HasCollision checks if a new path conflicts with an existing set of paths.
// Paths conflict if they share any intermediate rooms (excluding start and end).
func HasCollision(path []string, set [][]string) bool {
	if len(path) <= 2 {
		return false // No intermediate rooms to collide
	}

	// Collect intermediate rooms of the new path
	visited := make(map[string]bool, len(path)-2)
	for _, room := range path[1 : len(path)-1] {
		visited[room] = true
	}

	// Check against each path in the set
	for _, p := range set {
		for _, room := range p[1 : len(p)-1] {
			if visited[room] {
				return true // Collision found
			}
		}
	}
	return false
}
