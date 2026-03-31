package main

// compatible selection
func BuildSets(allpaths [][]string, index int, current [][]string, result *[][][]string) {
	if index == len(allpaths) {
		if len(current) > 0 {
			temp := make([][]string, len(current))
			copy(temp, current)
			*result = append(*result, temp)
		}
		return
	}

	BuildSets(allpaths, index+1, current, result)

	if canAdd(current, allpaths[index]) {
		newCurrent := append([][]string{}, current...)
		newCurrent = append(newCurrent, allpaths[index])

		BuildSets(allpaths, index+1, newCurrent, result)
	}

}

func canAdd(current [][]string, path []string) bool {
	for _, p := range current {
		if !compatible(p, path) {
			return false
		}
	}
	return true
}

func compatible(path1, path2 []string) bool {
	for i := 1; i < len(path1)-1; i++ {
		for j := 1; j < len(path2)-1; j++ {
			if path1[i] == path2[j] {
				return false
			}
		}
	}
	return true
}

// package main

// // compatible selection
// func Compatiblepaths(allpaths [][]string){
// 	var newallpath [][]string

// 	newallpath = append(newallpath, allpaths[0])

// 	for i := 1; i < len(allpaths); i++ {
// 		// if allpath[i] != newallpath '(canadd)' append allpath[i] to newallpath
// 		if canadd(newallpath, allpaths[i]) {
// 			newallpath = append(newallpath, allpaths[i])
// 		}
// 	}
// }

// func canadd(newallpath [][]string, path []string) bool {
// 	for _, p := range newallpath {
// 		if !compatible(p, path) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func compatible(path1, path2 []string) bool {
// 	for i := 1; i < len(path1)-1; i++ {
// 		for j := 1 ; j < len(path2)-1; j++{
// 			if path1[i] == path2[j]{
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
