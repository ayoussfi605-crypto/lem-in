package main

// L-mo7arrik li k-i-goul lik chkun hwa l-Best Set
func CalculateTurns(paths [][]string, ants int) int {
	if len(paths) == 0 {
		return 1<<31 - 1
	}

	// Water-filling logic
	pathAnts := make([]int, len(paths))
	for a := 0; a < ants; a++ {
		best := 0
		for i := 1; i < len(paths); i++ {
			// score = path_length + ants_already_on_it
			if len(paths[i])+pathAnts[i] < len(paths[best])+pathAnts[best] {
				best = i
			}
		}
		pathAnts[best]++
	}

	maxTurns := 0
	for i := range paths {
		turns := len(paths[i]) + pathAnts[i] - 1
		if turns > maxTurns {
			maxTurns = turns
		}
	}
	return maxTurns
}
