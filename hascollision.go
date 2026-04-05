package main

// Bach n-choufou wach triq jdiida k-t-t-qata3 m3a triqat li aslan 3ndna f l-set
func HasCollision(path []string, set [][]string) bool {
	for _, p := range set {
		// Intersection check (mashi start/end)
		for _, r1 := range path[1 : len(path)-1] {
			for _, r2 := range p[1 : len(p)-1] {
				if r1 == r2 {
					return true
				}
			}
		}
	}
	return false
}
