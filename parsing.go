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
	inputfile, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	farm := Farm{
		Rooms: make(map[string]Room),
		Adj:   make(map[string][]string),
	}

	content := string(inputfile)
	if content == "" {
		fmt.Println("ERROR: invalid data format")
		return
	}
	lines := strings.Split(strings.TrimRight(content, "\n"), "\n")
	farm.RawLines = lines

	foundAnts := false
	// check if alredy exist
	seenStartCmd, seenEndCmd := false, false
	// next room is start or end room
	expectStartRoom, expectEndRoom := false, false
	// if find this - next line He should is link if She was room... err
	seenLinks := false
	// "Two rooms can't have more than one tunnel connecting them".
	seenTunnels := make(map[string]bool)
	seenCoords := make(map[string]bool)

	for _, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		// Handle Commands
		if strings.HasPrefix(line, "##") {
			if !foundAnts {
				fmt.Println("ERROR: invalid data format")
				return
			}
			if line == "##start" {
				if seenLinks || seenStartCmd {
					fmt.Println("ERROR: invalid data format")
					return
				}
				seenStartCmd, expectStartRoom = true, true
				continue
			} else if line == "##end" {
				if seenLinks || seenEndCmd {
					fmt.Println("ERROR: invalid data format")
					return
				}
				seenEndCmd, expectEndRoom = true, true
				continue
			} else {
				// Ignore unknown commands (like ##anything)
				continue
			}
		}

		// Ignore Comments
		if strings.HasPrefix(line, "#") {
			continue
		}

		// Parse Ants (First valid non-comment line)
		if !foundAnts {
			ants, err := strconv.Atoi(line)
			if err != nil || ants <= 0 || strings.Contains(line, " ") {
				fmt.Println("ERROR: invalid data format")
				return
			}
			farm.Ants = ants
			foundAnts = true
			continue
		}

		// Links Check (A-B)
		if strings.Count(line, "-") == 1 && !strings.Contains(line, " ") {
			// He should To be room after flag ##start/end
			if expectStartRoom || expectEndRoom {
				fmt.Println("ERROR: invalid data format")
				return
			}
			seenLinks = true
			ab := strings.Split(line, "-")
			a, b := ab[0], ab[1]
			if _, ok1 := farm.Rooms[a]; !ok1 {
				fmt.Println("ERROR: invalid data format")
				return
			}
			// if this room a or b Available in map if not err
			if _, ok2 := farm.Rooms[b]; !ok2 {
				fmt.Println("ERROR: invalid data format")
				return
			}
			// Self-linking
			if a == b {
				fmt.Println("ERROR: invalid data format")
				return
			}
			// Duplicate Tunnels Check
			key := normTunnelKey(a, b)
			if seenTunnels[key] {
				fmt.Println("ERROR: invalid data format") 
				return
			}

			// if not available 
			seenTunnels[key] = true
			farm.Adj[a] = append(farm.Adj[a], b)
			farm.Adj[b] = append(farm.Adj[b], a)
			continue
		}

		// Rooms Check (Name X Y)
		parts := strings.Fields(line)
		if len(parts) == 3 {
			if seenLinks {
				fmt.Println("ERROR: invalid data format")
				return
			}
			name := parts[0]
			if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
				fmt.Println("ERROR: invalid data format")
				return
			}
			if _, exists := farm.Rooms[name]; exists {
				fmt.Println("ERROR: invalid data format")
				return
			}

			x, errX := strconv.Atoi(parts[1])
			y, errY := strconv.Atoi(parts[2])
			if errX != nil || errY != nil {
				fmt.Println("ERROR: invalid data format")
				return
			}
			coordKey := fmt.Sprintf("%d,%d", x, y)
			if seenCoords[coordKey] {
				fmt.Println("ERROR: invalid data format")
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

		// Any other line format is an error
		fmt.Println("ERROR: invalid data format")
		return
	}

	// Final Validation
	if farm.Start == "" || farm.End == "" || !foundAnts {
		fmt.Println("ERROR: invalid data format")
		return
	}

	// 1. Print Raw Data (Required by subject)
	for _, l := range farm.RawLines {
		fmt.Println(l)
	}
	fmt.Println()

	// 2. Algorithm
	paths := EdmondsKarp(farm)
	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	Solve(farm.Ants, paths, farm)
}

// Duplicate Tunnels Alphabetically B > A -> A|B
func normTunnelKey(a, b string) string {
	if a > b {
		return b + "|" + a
	}
	return a + "|" + b
}
