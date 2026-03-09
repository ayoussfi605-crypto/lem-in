package main

import (
	"fmt"
)

func Bfs(farm Farm, room Room) []string {
	queue := []string{farm.Start}

	visite := make(map[string]bool)
	visite[farm.Start] = true

	parent := make(map[string]string)

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		if current == farm.End {
			break
		}

		for _, ng := range farm.Adj[current] {
			if !visite[ng] {

				visite[ng] = true
				parent[ng] = current

				queue = append(queue, ng)
			}
		}

	}
	if !visite[farm.End] {
		return nil
	}
	fmt.Println("parent :",parent)
	rev := []string{}
	d := farm.End

	for {
		rev = append(rev, d)
		if d == farm.Start {
			break
		}
		d = parent[d]
	}
	path := []string{}
	for i := len(rev) - 1; i >= 0; i-- {
		path = append(path, rev[i])
	}

	return path
}
