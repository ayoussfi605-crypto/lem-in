package helper

func Dfs(farm Farm) [][]string {
	var allPaths [][]string
	limit := 100

	visited := make(map[string]bool)
	path := []string{}

	var dfs func(node string)

	dfs = func(node string) {
		if len(allPaths) >= limit {
			return
		}

		// add node to path
		visited[node] = true
		path = append(path, node)

		// reached end
		if node == farm.End {
			// copy path before saving
			cp := make([]string, len(path))
			copy(cp, path)
			allPaths = append(allPaths, cp)
		} else {
			// explore neighbors
			for _, ng := range farm.Adj[node] {
				if !visited[ng] {
					dfs(ng)
				}
			}
		}

		// backtrack (key for low RAM)
		path = path[:len(path)-1]
		visited[node] = false
	}

	dfs(farm.Start)
	return allPaths
}
