package main

import (
	"sort"
)

func GetBestSet(allPaths [][]string, totalAnts int) [][]string {
	// 1. Sort paths by length (Dima qsar homa l-uwwal)
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	var bestSet [][]string
	minTurns := 1<<31 - 1

	// 2. Greedy Loop: Jarreb n-bniw sets m-khtalfin
	// Koul mra n-bdaw b-path m-khtalfa (i) u n-choufou chkun li mzyan m3aha
	limit := len(allPaths)
	if limit > 100 {
		limit = 100
	} // Bach n-kunu sra3

	for i := 0; i < limit; i++ {
		currentSet := [][]string{allPaths[i]}

		for j := 0; j < len(allPaths); j++ {
			if i == j {
				continue
			}

			// Ila kant had l-triq (j) ma-m-charkach f l-biutan m3a l-set l-7aliya
			if !HasCollision(allPaths[j], currentSet) {
				currentSet = append(currentSet, allPaths[j])
			}
		}

		// 3. 7seb ch7al mn turn ghadi n-7taju b had l-set
		turns := CalculateTurns(currentSet, totalAnts)
		if turns < minTurns {
			minTurns = turns
			bestSet = currentSet
		}
	}
	return bestSet
}
