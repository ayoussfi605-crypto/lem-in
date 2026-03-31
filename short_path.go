package main

func bestset(allpath [][][]string) [][][]string {
	best := allpath[0]
	for i := 1; i < len(allpath); i++ {
		if checklen(best, allpath[i]) {

		}
	}
	return allpath
}

func checklen(set1, set2 [][]string) bool {

}
