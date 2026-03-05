package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name string
	X    int
	Y    int
}

type Farm struct {
	Ants int

	Rooms map[string]Room     // name -> Room
	Adj   map[string][]string // name -> neighbors

	Start string
	End   string

	RawLines []string // keep original lines to print back
}

func main() {
	Parsfile("file.txt")
}

func Parsfile(filename string) {
	foundAnts := false

	seenStartCmd := false
	seenEndCmd := false

	inputfile, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	farm := Farm{
		Rooms: make(map[string]Room),
		Adj:   make(map[string][]string),
	}

	input := strings.Split(strings.TrimRight(string(inputfile), "\n"), "\n")
	farm.RawLines = input

	for _, line := range farm.RawLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// If ##start/##end appear before ants => invalid
		if !foundAnts && (line == "##start" || line == "##end") {
			fmt.Println("ERROR: invalid data format")
			return
		}

		// Commands (after ants)
		if foundAnts && line == "##start" {
			if seenStartCmd {
				fmt.Println("ERROR: invalid data format")
				return
			}
			seenStartCmd = true
			continue
		}
		if foundAnts && line == "##end" {
			if seenEndCmd {
				fmt.Println("ERROR: invalid data format")
				return
			}
			seenEndCmd = true
			continue
		}

		// Ignore normal comments (including unknown ##commands)
		if strings.HasPrefix(line, "#") {
			continue
		}
		// If ants already parsed, a pure integer line is invalid (duplicate ants line)
		if foundAnts {
			if len(strings.Fields(line)) == 1 && !strings.Contains(line, "-") {
				if _, err := strconv.Atoi(line); err == nil {
					fmt.Println("ERROR: invalid data format")
					return
				}
			}
		}
		// Ants: first non-comment non-empty line
		if !foundAnts {
			if strings.Contains(line, "-") || len(strings.Fields(line)) != 1 {
				fmt.Println("ERROR: invalid data format")
				return
			}

			ants, err := strconv.Atoi(line)
			if err != nil || ants <= 0 {
				fmt.Println("ERROR: invalid data format")
				return
			}

			farm.Ants = ants
			foundAnts = true
			continue
		}

		// For now we stop here (rooms/links later)
	}

	// Optional debug
	fmt.Println("Ants:", farm.Ants)
	fmt.Println("Seen ##start:", seenStartCmd)
	fmt.Println("Seen ##end:", seenEndCmd)
}
