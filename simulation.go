package main

import "fmt"

func Simulation(path []string, ants int) []int {
	porsition := make([]int, ants)

	for i := 0; i < len(porsition); i++ {
		
		porsition[len(porsition)-1] = i-1
		

	}
	fmt.Println(porsition)
	return porsition
}

// path := []string{"1", "3", "4", "0"}
// positions := []int{0, 0, 0}
