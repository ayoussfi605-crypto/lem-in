package main

import (
	"fmt"
	"strings"
)

func Solve(totalAnts int, paths [][]string, farm Farm) {
	// Distribute ants to paths
	antPathIdx := make([]int, totalAnts+1)
	pathAntsCount := make([]int, len(paths))

	for i := 1; i <= totalAnts; i++ {
		bestP := 0
		minScore := len(paths[0]) + pathAntsCount[0]
		for p := 1; p < len(paths); p++ {
			score := len(paths[p]) + pathAntsCount[p]
			if score < minScore {
				minScore = score
				bestP = p
			}
		}
		antPathIdx[i] = bestP
		pathAntsCount[bestP]++
	}

	// Simulation
	antPosInPath := make([]int, totalAnts+1) // Index in the path slice
	finished := 0
	released := 0

	for finished < totalAnts {
		var moves []string
		occupied := make(map[string]bool)

		for i := 1; i <= totalAnts; i++ {
			pIdx := antPathIdx[i]
			path := paths[pIdx]

			// Move ants that are already "out" or release a new one
			if antPosInPath[i] > 0 && antPosInPath[i] < len(path)-1 {
				antPosInPath[i]++
				room := path[antPosInPath[i]]
				moves = append(moves, fmt.Sprintf("L%d-%s", i, room))
				if room == farm.End {
					finished++
				}
			} else if antPosInPath[i] == 0 && released < totalAnts {
				// Only release if it's this ant's turn (greedy release)
				if canRelease(i, antPathIdx, released) {
					nextRoom := path[1]
					if !occupied[nextRoom] || nextRoom == farm.End {
						antPosInPath[i] = 1
						released++
						moves = append(moves, fmt.Sprintf("L%d-%s", i, nextRoom))
						occupied[nextRoom] = true
						if nextRoom == farm.End {
							finished++
						}
					}
				}
			}
		}
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}

func canRelease(antID int, antPathIdx []int, released int) bool {
	// Simple logic: ants are released in order 1, 2, 3...
	return antID == released+1
}
