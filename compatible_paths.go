package main

import (
	"math"
	"sort"
)

func GetBestSet(allPaths [][]string, totalAnts int) [][]string {
	// 1. Sort paths by length (shortest first)
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	var bestSet [][]string
	minTurns := math.MaxInt32

	// 2. Greedy algorithm: Try starting with each path and add compatible ones
	// Limit to first 100 paths to avoid excessive computation
	limit := len(allPaths)
	if limit > 100 {
		limit = 100
	}

	for i := 0; i < limit; i++ {
		// Start with the i-th path
		currentSet := [][]string{allPaths[i]}

		// Add other paths that don't collide with the current set
		for j := 0; j < len(allPaths); j++ {
			if i == j {
				continue
			}
			if !HasCollision(allPaths[j], currentSet) {
				currentSet = append(currentSet, allPaths[j])
			}
		}

		// 3. Calculate the number of turns needed for this set
		turns := CalculateTurns(currentSet, totalAnts)
		if turns < minTurns {
			minTurns = turns
			bestSet = currentSet
		}
	}
	return bestSet
}
