package main

// kol mra kankhrj path mn l queue kanchof akhit node wkanzid liha ng jdad bach nwsa3 dak lpath hta twsl l end
func Bfs(farm Farm) {
	// paths incomplete
	queue := [][]string{{farm.Start}}
	// paths complete
	var allpaths [][]string

	for len(queue) > 0 {

		path := queue[0]
		queue = queue[1:]
		last := path[len(path)-1]
		// path already completed zslat l end
		if last == farm.End {
			allpaths = append(allpaths, path)
			continue //continue maghadich nzid nws3 had l path bcause already wslat l end
		}

		for _, ng := range farm.Adj[last] {
			// check if ng is existe in path
			if !Contains(path, ng) {
				newpath := append([]string{}, path...)
				newpath = append(newpath, ng)

				queue = append(queue, newpath)
			}
		}
	}
	var allSets [][][]string
	BuildSets(allpaths, 0, [][]string{}, &allSets)

	Bestset(allSets)
}
