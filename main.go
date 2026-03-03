package main

import (
	"fmt"
	"os"
	"strings"
)

type colony struct{
	rawLines []string
	ants int 
	rooms map[string]Room
	startRoom string
	endRoom string
	adj map[string][]string
}

func main() {
	Parsfile("file.txt")
}

func Parsfile(filename string) {
	inputfile, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println("err")
	}

	input := strings.Split(string(inputfile), "\n")
	fmt.Println(input)


}
