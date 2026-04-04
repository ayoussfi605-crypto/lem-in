package main

func EdmondsKarp(farm Farm) [][]string {
	var allPaths [][]string
	usedRooms := make(map[string]bool)

	for {
		parent := make(map[string]string)
		queue := []string{farm.Start}
		parent[farm.Start] = "START_NODE"
		found := false

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			if curr == farm.End {
				found = true
				break
			}

			for _, next := range farm.Adj[curr] {
				if _, visited := parent[next]; !visited && (next == farm.End || !usedRooms[next]) {
					parent[next] = curr
					queue = append(queue, next)
				}
			}
			if found {
				break
			}
		}

		if !found {
			break
		}

		var path []string
		curr := farm.End
		for curr != "START_NODE" {
			path = append([]string{curr}, path...)
			if curr != farm.Start && curr != farm.End {
				usedRooms[curr] = true
			}
			curr = parent[curr]
		}
		allPaths = append(allPaths, path)
	}
	return allPaths
}
