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

	Rooms map[string]Room
	Adj   map[string][]string

	Start string
	End   string

	RawLines []string
}

func main() {
	Parsfile("file.txt")
}

func Parsfile(filename string) {
	foundAnts := false

	seenStartCmd := false
	seenEndCmd := false

	expectStartRoom := false
	expectEndRoom := false

	seenLinks := false
	seenTunnels := make(map[string]bool)

	inputfile, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	farm := Farm{
		Rooms: make(map[string]Room),
		Adj:   make(map[string][]string),
	}

	lines := strings.Split(strings.TrimRight(string(inputfile), "\n"), "\n")
	farm.RawLines = lines

	for _, raw := range farm.RawLines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		// Commands before ants => invalid
		if !foundAnts && (line == "##start" || line == "##end") {
			fmt.Println("ERROR: invalid data format")
			return
		}

		// Handle commands (after ants) - also disallow after links started
		if foundAnts && line == "##start" {
			if seenLinks {
				fmt.Println("ERROR: invalid data format")
				return
			}
			if seenStartCmd {
				fmt.Println("ERROR: invalid data format")
				return
			}
			seenStartCmd = true
			expectStartRoom = true
			expectEndRoom = false
			continue
		}
		if foundAnts && line == "##end" {
			if seenLinks {
				fmt.Println("ERROR: invalid data format")
				return
			}
			if seenEndCmd {
				fmt.Println("ERROR: invalid data format")
				return
			}
			seenEndCmd = true
			expectEndRoom = true
			expectStartRoom = false
			continue
		}

		// Ignore comments
		if strings.HasPrefix(line, "#") {
			continue
		}

		// Parse ants (first valid non-comment line)
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

		// Duplicate ants line (a pure integer after ants) => invalid
		if len(strings.Fields(line)) == 1 && !strings.Contains(line, "-") {
			if _, err := strconv.Atoi(line); err == nil {
				fmt.Println("ERROR: invalid data format")
				return
			}
		}

		// LINKS: "A-B" (single token, exactly one '-')
		if strings.Count(line, "-") == 1 && len(strings.Fields(line)) == 1 {
			// if we were waiting for start/end room but got a link => invalid
			if expectStartRoom || expectEndRoom {
				fmt.Println("ERROR: invalid data format")
				return
			}

			seenLinks = true

			ab := strings.Split(line, "-")
			a := strings.TrimSpace(ab[0])
			b := strings.TrimSpace(ab[1])

			if a == "" || b == "" {
				fmt.Println("ERROR: invalid data format")
				return
			}

			// self-link invalid
			if a == b {
				fmt.Println("ERROR: invalid data format")
				return
			}

			// unknown rooms invalid
			if _, ok := farm.Rooms[a]; !ok {
				fmt.Println("ERROR: invalid data format")
				return
			}
			if _, ok := farm.Rooms[b]; !ok {
				fmt.Println("ERROR: invalid data format")
				return
			}

			// duplicate tunnel invalid (A-B == B-A)
			key := normTunnelKey(a, b)
			if seenTunnels[key] {
				fmt.Println("ERROR: invalid data format")
				return
			}
			seenTunnels[key] = true

			// build graph (bidirectional)
			farm.Adj[a] = append(farm.Adj[a], b)
			farm.Adj[b] = append(farm.Adj[b], a)

			continue
		}

		// ROOMS: "name x y"
		parts := strings.Fields(line)
		if len(parts) == 3 {
			// rooms after links => invalid
			if seenLinks {
				fmt.Println("ERROR: invalid data format")
				return
			}

			name := parts[0]

			// room name rules
			if name == "" || strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
				fmt.Println("ERROR: invalid data format")
				return
			}

			x, err1 := strconv.Atoi(parts[1])
			y, err2 := strconv.Atoi(parts[2])
			if err1 != nil || err2 != nil {
				fmt.Println("ERROR: invalid data format")
				return
			}

			// duplicate room
			if _, exists := farm.Rooms[name]; exists {
				fmt.Println("ERROR: invalid data format")
				return
			}

			farm.Rooms[name] = Room{Name: name, X: x, Y: y}

			// bind start/end
			if expectStartRoom {
				if farm.Start != "" {
					fmt.Println("ERROR: invalid data format")
					return
				}
				farm.Start = name
				expectStartRoom = false
			}
			if expectEndRoom {
				if farm.End != "" {
					fmt.Println("ERROR: invalid data format")
					return
				}
				farm.End = name
				expectEndRoom = false
			}

			continue
		}

		// If we were waiting for a room after ##start/##end but got something else => invalid
		if expectStartRoom || expectEndRoom {
			fmt.Println("ERROR: invalid data format")
			return
		}

		// Unknown line => invalid (better for tests)
		fmt.Println("ERROR: invalid data format")
		return
	}

	// Final checks after loop
	if expectStartRoom || expectEndRoom {
		fmt.Println("ERROR: invalid data format")
		return
	}
	if farm.Start == "" || farm.End == "" {
		fmt.Println("ERROR: invalid data format")
		return
	}
	allPaths := Bfs(farm)
	// 1. Lqa l-paths
	bestSet := GetBestSet(allPaths, farm.Ants)

	// 2. Farraq n-nmel 3la l-paths
	antsReady := Divisionofants(bestSet, farm.Ants)

	// 1. Print input data (Ants, Rooms, Links)
	fmt.Println(farm.RawLines)
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
