package main

func Bestset(allset [][][]string) [][]string {
	best := allset[0]

	for i := 1; i < len(allset); i++ {
		if len(allset[i]) > len(best) {
			best = allset[i]
		} else if len(allset[i]) == len(best) {
			if scorlen(allset[i]) < scorlen(best) {
				best = allset[i]
			}
		}
	}
	return best
}

func scorlen(set [][]string) int {
	total := 0
	for _, p := range set {
		total += len(p)
	}
	return total
}
