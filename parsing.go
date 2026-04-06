package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name string
}

type Farm struct {
	Ants     int
	Rooms    map[string]Room
	Adj      map[string][]string
	Start    string
	End      string
	RawLines []string
}

func Parsfile(filename string) {
	// Read the input file
	inputfile, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR: invalid data format, cannot read file")
		return
	}

	farm := Farm{
		Rooms: make(map[string]Room),
		Adj:   make(map[string][]string),
	}

	content := string(inputfile)
	if content == "" {
		fmt.Println("ERROR: invalid data format, empty file")
		return
	}
	lines := strings.Split(strings.TrimRight(content, "\n"), "\n")
	farm.RawLines = lines

	// Flags and maps for validation
	foundAnts := false
	seenStartCmd, seenEndCmd := false, false
	expectStartRoom, expectEndRoom := false, false
	seenLinks := false
	seenTunnels := make(map[string]bool)
	seenCoords := make(map[string]bool)

	for _, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		// Handle Commands (##start, ##end)
		if strings.HasPrefix(line, "##") {
			if !foundAnts {
				fmt.Println("ERROR: invalid data format, commands before ant count")
				return
			}
			if line == "##start" {
				if seenLinks || seenStartCmd || expectStartRoom || expectEndRoom {
					fmt.Println("ERROR: invalid data format, duplicate or misplaced start command")
					return
				}
				seenStartCmd, expectStartRoom = true, true
				continue
			} else if line == "##end" {
				if seenLinks || seenEndCmd || expectStartRoom || expectEndRoom {
					fmt.Println("ERROR: invalid data format, duplicate or misplaced end command")
					return
				}
				seenEndCmd, expectEndRoom = true, true
				continue
			} else {
				// Ignore unknown commands
				continue
			}
		}

		// Ignore Comments
		if strings.HasPrefix(line, "#") {
			continue
		}

		// Parse Ants (first valid non-comment line)
		if !foundAnts {
			ants, err := strconv.Atoi(line)
			if err != nil || ants <= 0 || strings.Contains(line, " ") {
				fmt.Println("ERROR: invalid data format, invalid number of ants")
				return
			}
			farm.Ants = ants
			foundAnts = true
			continue
		}

		// Parse Links (A-B)
		if strings.Count(line, "-") == 1 && !strings.Contains(line, " ") {
			if expectStartRoom || expectEndRoom {
				fmt.Println("ERROR: invalid data format, link instead of room after command")
				return
			}
			seenLinks = true
			ab := strings.Split(line, "-")
			a, b := ab[0], ab[1]
			if _, ok1 := farm.Rooms[a]; !ok1 {
				fmt.Println("ERROR: invalid data format, link to unknown room")
				return
			}
			if _, ok2 := farm.Rooms[b]; !ok2 {
				fmt.Println("ERROR: invalid data format, link to unknown room")
				return
			}
			if a == b {
				fmt.Println("ERROR: invalid data format, self-linking room")
				return
			}
			key := normTunnelKey(a, b)
			if seenTunnels[key] {
				fmt.Println("ERROR: invalid data format, duplicate tunnel")
				return
			}
			seenTunnels[key] = true
			farm.Adj[a] = append(farm.Adj[a], b)
			farm.Adj[b] = append(farm.Adj[b], a)
			continue
		}

		// Parse Rooms (Name X Y)
		parts := strings.Fields(line)
		if len(parts) == 3 {
			if seenLinks {
				fmt.Println("ERROR: invalid data format, rooms after links")
				return
			}
			if expectStartRoom && expectEndRoom {
				fmt.Println("ERROR: invalid data format, both start and end expected simultaneously")
				return
			}
			name := parts[0]
			if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
				fmt.Println("ERROR: invalid data format, invalid room name")
				return
			}
			if _, exists := farm.Rooms[name]; exists {
				fmt.Println("ERROR: invalid data format, duplicate room name")
				return
			}
			x, errX := strconv.Atoi(parts[1])
			y, errY := strconv.Atoi(parts[2])
			if errX != nil || errY != nil {
				fmt.Println("ERROR: invalid data format, invalid coordinates")
				return
			}
			coordKey := fmt.Sprintf("%d,%d", x, y)
			if seenCoords[coordKey] {
				fmt.Println("ERROR: invalid data format, duplicate coordinates")
				return
			}
			seenCoords[coordKey] = true
			farm.Rooms[name] = Room{Name: name}
			if _, exists := farm.Adj[name]; !exists {
				farm.Adj[name] = []string{}
			}
			if expectStartRoom {
				farm.Start = name
				expectStartRoom = false
			}
			if expectEndRoom {
				farm.End = name
				expectEndRoom = false
			}
			continue
		}

		// Unknown line format
		fmt.Println("ERROR: invalid data format, unknown line format")
		return
	}

	// Final Validation
	if farm.Start == "" {
		fmt.Println("ERROR: invalid data format, no start room found")
		return
	}
	if farm.End == "" {
		fmt.Println("ERROR: invalid data format, no end room found")
		return
	}
	if !foundAnts {
		fmt.Println("ERROR: invalid data format, no ant count found")
		return
	}

	allPaths := Dfs(farm)
	// 1. Find all paths from start to end
	bestSet := GetBestSet(allPaths, farm.Ants)

	// 2. Distribute ants across the best set of paths
	antsReady := Divisionofants(bestSet, farm.Ants)
	if len(bestSet) == 0 {
		fmt.Println("ERROR: invalid data format, no valid paths found")
		return
	}
	// 1. Print input data (Ants, Rooms, Links)
	for _, r := range farm.RawLines {
		fmt.Println(r)
	}
	fmt.Println() // Khlliw s-tr khawi bin l-data u l-moves

	// 3. Simuli l-movement (Solve)
	Solve(antsReady, farm)
}

func normTunnelKey(a, b string) string {
	if a > b {
		a, b = b, a
	}
	return a + "|" + b
}
