package main

import (
	"fmt"
	"sort"
)

type Ant struct {
	ID   int      // L1, L2, L3...
	Path []string // Triq li ghadi t-ched
	Step int      // Fin wsset f-l-index d l-path
}

func Divisionofants(bestSet [][]string, totalAnts int) []Ant {
	// 1. Check ila l-BFS ma-lqat walo (Critical Check)
	if len(bestSet) == 0 {
		fmt.Println("Error: No paths found between start and end.")
		return []Ant{}
	}
	// 1. Distribution Logic (L-7sab)
	// antsPerPath[i] k-t-goul lina ch7al mn nemla f-l-path i
	antsPerPath := make([][]int, len(bestSet))

	for antID := 1; antID <= totalAnts; antID++ {
		bestIdx := 0
		// Score = toul d l-path + ch7al mn nemla aslan fiha
		minScore := len(bestSet[0]) + len(antsPerPath[0])

		for i := 1; i < len(bestSet); i++ {
			score := len(bestSet[i]) + len(antsPerPath[i])
			if score < minScore {
				minScore = score
				bestIdx = i
			}
		}
		antsPerPath[bestIdx] = append(antsPerPath[bestIdx], antID)
	}

	// 2. Map IDs to Ant structs
	// Ghadi n-creaw ga3 n-nmel u n-3tiw l-koll nemla l-path dyalha
	var allAnts []Ant
	for i, antIDs := range antsPerPath {
		for _, id := range antIDs {
			allAnts = append(allAnts, Ant{
				ID:   id,
				Path: bestSet[i],
				Step: 0, // Dima k-i-bdaw mn Start
			})
		}
	}

	// 3. Sort by ID (Bach n-releasiw L1, L2, L3... b-tartib)
	sort.Slice(allAnts, func(i, j int) bool {
		return allAnts[i].ID < allAnts[j].ID
	})

	return allAnts
}
