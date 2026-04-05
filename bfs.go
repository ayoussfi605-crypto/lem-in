package main

// kol mra kankhrj path mn l queue kanchof akhit node wkanzid liha ng jdad bach nwsa3 dak lpath hta twsl l end
func Bfs(farm Farm) [][]string {
	queue := [][]string{{farm.Start}}
	var allPaths [][]string
	limit := 500 // Safety limit bach may-t-bloquach l-program f maps kbar

	for len(queue) > 0 && len(allPaths) < limit {
		path := queue[0]
		queue = queue[1:]
		last := path[len(path)-1]

		if last == farm.End {
			allPaths = append(allPaths, path)
			continue
		}

		for _, ng := range farm.Adj[last] {
			if !Contains(path, ng) {
				newPath := make([]string, len(path)+1)
				copy(newPath, path)
				newPath[len(path)] = ng
				queue = append(queue, newPath)
			}
		}
	}
	return allPaths
}
