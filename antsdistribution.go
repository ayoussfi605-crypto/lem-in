package main

import (
	"fmt"
	"strings"
)

// Solve simulates the ant movement turn by turn.
// It moves ants along their paths, ensuring no two ants occupy the same room except the end.
func Solve(ants []Ant, farm Farm) {
	finished := 0
	totalAnts := len(ants)

	// occupied tracks rooms that have ants in them this turn (except end, which allows multiple)
	for finished < totalAnts {
		moves := make([]string, 0, len(ants))
		occupied := make(map[string]bool, len(ants))

		// Process each ant
		for i := 0; i < len(ants); i++ {
			// Skip if ant has finished
			if ants[i].Step >= len(ants[i].Path)-1 {
				continue
			}

			nextIdx := ants[i].Step + 1
			nextRoom := ants[i].Path[nextIdx]

			// Move if the next room is free or is the end
			if !occupied[nextRoom] || nextRoom == farm.End {
				ants[i].Step = nextIdx
				moves = append(moves, fmt.Sprintf("L%d-%s", ants[i].ID, nextRoom))

				if nextRoom == farm.End {
					finished++
				} else {
					occupied[nextRoom] = true
				}
			}
		}

		// Print moves for this turn
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}
