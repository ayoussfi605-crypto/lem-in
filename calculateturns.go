package main

// CalculateTurns computes the minimum number of turns needed to move all ants
// using the given set of paths. It uses a water-filling approach to distribute ants.
func CalculateTurns(paths [][]string, ants int) int {
	if len(paths) == 0 {
		return 1<<31 - 1 // Max int, indicating invalid
	}

	// pathAnts[i] = number of ants assigned to path i
	pathAnts := make([]int, len(paths))
	for a := 0; a < ants; a++ {
		// Find the path with the lowest current score (length + ants already on it)
		best := 0
		for i := 1; i < len(paths); i++ {
			if len(paths[i])+pathAnts[i] < len(paths[best])+pathAnts[best] {
				best = i
			}
		}
		pathAnts[best]++
	}

	// The maximum turns is the max over all paths of (path length + ants on it - 1)
	maxTurns := 0
	for i := range paths {
		turns := len(paths[i]) + pathAnts[i] - 1
		if turns > maxTurns {
			maxTurns = turns
		}
	}
	return maxTurns
}
