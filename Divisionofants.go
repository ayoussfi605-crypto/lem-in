package main

import (
	"fmt"
	"sort"
)

// Ant represents an ant with its ID, assigned path, and current position.
type Ant struct {
	ID   int      // Unique ID (L1, L2, etc.)
	Path []string // The path this ant follows
	Step int      // Current index in the path (0 = start)
}

// Divisionofants distributes ants across the best set of paths using greedy assignment.
// It assigns each ant to the path that would minimize the total turns.
func Divisionofants(bestSet [][]string, totalAnts int) []Ant {
	// Check for valid input
	if len(bestSet) == 0 {
		fmt.Println("Error: No paths found between start and end.")
		return []Ant{}
	}

	// antsPerPath[i] = number of ants on path i
	antsPerPath := make([][]int, len(bestSet))

	// Distribute ants greedily
	for antID := 1; antID <= totalAnts; antID++ {
		bestIdx := 0
		minScore := len(bestSet[0]) + len(antsPerPath[0]) // Score = path length + current ants

		for i := 1; i < len(bestSet); i++ {
			score := len(bestSet[i]) + len(antsPerPath[i])
			if score < minScore {
				minScore = score
				bestIdx = i
			}
		}
		antsPerPath[bestIdx] = append(antsPerPath[bestIdx], antID)
	}

	// Create Ant structs
	var allAnts []Ant
	for i, antIDs := range antsPerPath {
		for _, id := range antIDs {
			allAnts = append(allAnts, Ant{
				ID:   id,
				Path: bestSet[i],
				Step: 0, // Start at the beginning
			})
		}
	}

	// Sort by ID for consistent output
	sort.Slice(allAnts, func(i, j int) bool {
		return allAnts[i].ID < allAnts[j].ID
	})

	return allAnts
}
