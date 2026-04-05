package main

import (
	"fmt"
	"strings"
)

func Solve(ants []Ant, farm Farm) {
	finished := 0
	totalAnts := len(ants)

	// Kan-st3mlo had l-map bach n-choufu achmen bit m-occupied f koll turn
	// Bach n-fadaou l-izdi7am (Collisions)
	for finished < totalAnts {
		var moves []string
		occupied := make(map[string]bool)

		// Koll turn, n-jerrbou n-7rrkou koll nemla
		for i := 0; i < len(ants); i++ {
			// Ila n-nemla aslan وصلات l-End, n-nagzouha
			if ants[i].Step >= len(ants[i].Path)-1 {
				continue
			}

			// N-choufou l-bit li jaya f l-path dyal had n-nemla
			nextIdx := ants[i].Step + 1
			nextRoom := ants[i].Path[nextIdx]

			// Qā3ida: N-7rrkouha illa ila l-bit khawya WALA hiya l-End room
			// (L-End room tqder t-hizz n-nmel kamel f-daqqa)
			if !occupied[nextRoom] || nextRoom == farm.End {

				// N-7eydou l-nemla mn l-bit l-qdima (ila ma-kantch Start)
				// u n-markiw l-bit jdiida blli m-occupied
				ants[i].Step = nextIdx

				// Format d l-output: Lx-y
				moves = append(moves, fmt.Sprintf("L%d-%s", ants[i].ID, nextRoom))

				if nextRoom == farm.End {
					finished++
				} else {
					occupied[nextRoom] = true
				}
			}
		}

		// Print l-moves dyal had l-turn f s-str wahed
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}
