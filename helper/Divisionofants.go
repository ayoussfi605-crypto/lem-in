package helper

type Ant struct {
	ID   int
	Path []string
	Step int
}

func Divisionofants(bestSet [][]string, totalAnts int) []Ant {
	if len(bestSet) == 0 {
		return []Ant{}
	}

	allAnts := make([]Ant, totalAnts)
	// pathUsage k-t-7seb ch7al mn nemla dkhlat l-koll triq
	pathUsage := make([]int, len(bestSet))

	for id := 1; id <= totalAnts; id++ {
		// 1. Dima bda u tkhayel blli triq l-uwwla hiya l-best f had l-marra
		bestIdx := 0
		minScore := len(bestSet[0]) + pathUsage[0]

		// 2. Qaren m3a l-triqat khrin
		for i := 1; i < len(bestSet); i++ {
			score := len(bestSet[i]) + pathUsage[i]
			if score < minScore {
				minScore = score
				bestIdx = i
			}
		}

		// 3. Assign n-nemla l-dik l-triq li 3ndha aqal score
		allAnts[id-1] = Ant{
			ID:   id,
			Path: bestSet[bestIdx],
			Step: 0,
		}

		// 4. Update usage: Zid nemla f l-7ssab dyal dik l-triq
		pathUsage[bestIdx]++
	}

	return allAnts
}
