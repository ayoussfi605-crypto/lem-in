// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// type Room struct {
// 	Name string
// 	X    int
// 	Y    int
// }

// type Farm struct {
// 	Ants int

// 	Rooms map[string]Room     // name -> Room
// 	Adj   map[string][]string // name -> neighbors

// 	Start string
// 	End   string

// 	RawLines []string // keep original lines to print back
// }

// func main() {
// 	Parsfile("file.txt")
// }

// func Parsfile(filename string) {
// 	isAnts := false
// 	inputfile, err := os.ReadFile(filename)
// 	if err != nil {
// 		fmt.Println("ERROR: invalid data format")
// 		return
// 	}
// 	farm := Farm{
// 		Rooms: make(map[string]Room),
// 		Adj:   make(map[string][]string),
// 	}
// 	input := strings.Split(string(inputfile), "\n")
// 	farm.RawLines = input

// 	for _, line := range farm.RawLines {

// 		ants, err := strconv.Atoi(line)
// 		if err != nil {
// 			fmt.Println("error for convert.Atoi")
// 		}
// 		if !isAnts {
// 			farm.Ants = ants
// 			isAnts = true
// 			break
// 		}
// 	}
// 	fmt.Println(farm.Ants)
// }

package main

import (
	"bufio"
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
	Ants  int
	Rooms map[string]Room
	Adj   map[string][]string
	Start string
	End   string
}

// ParseFile reads a lem-in file and validates:
// - ants line
// - rooms (including ##start / ##end)
// - links
// - duplicate rooms
// - unknown rooms in links
// - room name starting with 'L' or '#'
// - rooms after links
// - duplicate tunnels (A-B and B-A count as same)
func ParseFile(filename string) (*Farm, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ERROR: invalid data format")
	}
	defer file.Close()

	farm := &Farm{
		Rooms: make(map[string]Room),
		Adj:   make(map[string][]string),
	}

	scanner := bufio.NewScanner(file)

	expectStart := false
	expectEnd := false
	foundAnts := false
	seenLinks := false

	seenStartCmd := false
	seenEndCmd := false

	// Track tunnels to detect duplicates:
	// key is normalized like "A|B" (sorted)
	seenTunnels := make(map[string]bool)

	for scanner.Scan() {
		raw := scanner.Text()
		line := strings.TrimSpace(raw)

		// Ignore empty lines
		if line == "" {
			continue
		}

		// Handle comments and commands
		if strings.HasPrefix(line, "#") {
			if line == "##start" {
				if seenStartCmd {
					return nil, fmt.Errorf("ERROR: invalid data format")
				}
				seenStartCmd = true
				expectStart = true
				expectEnd = false
			} else if line == "##end" {
				if seenEndCmd {
					return nil, fmt.Errorf("ERROR: invalid data format")
				}
				seenEndCmd = true
				expectEnd = true
				expectStart = false
			}
			// Any other comment is ignored
			continue
		}

		// 1) Parse ants (first non-comment, non-empty line)
		if !foundAnts {
			// ants line must be a single integer token (no '-', no spaces)
			if strings.Contains(line, "-") || len(strings.Fields(line)) != 1 {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}
			ants, err := strconv.Atoi(line)
			if err != nil || ants <= 0 {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}
			farm.Ants = ants
			foundAnts = true
			continue
		}

		// 2) Link line: must be a single token with exactly one '-'
		if isLinkLine(line) {
			seenLinks = true

			a, b, ok := splitLink(line)
			if !ok {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}
			if a == b {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}

			// Unknown rooms in links => ERROR
			if _, ok := farm.Rooms[a]; !ok {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}
			if _, ok := farm.Rooms[b]; !ok {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}

			// Duplicate tunnels => ERROR (A-B equals B-A)
			key := normTunnelKey(a, b)
			if seenTunnels[key] {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}
			seenTunnels[key] = true

			// Add to adjacency list (bidirectional)
			farm.Adj[a] = append(farm.Adj[a], b)
			farm.Adj[b] = append(farm.Adj[b], a)
			continue
		}

		// 3) Room line: "name x y"
		if isRoomLine(line) {
			// Rooms after links => ERROR (format is ants -> rooms -> links)
			if seenLinks {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}

			room, err := parseRoom(line)
			if err != nil {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}

			// Duplicate rooms => ERROR
			if _, exists := farm.Rooms[room.Name]; exists {
				return nil, fmt.Errorf("ERROR: invalid data format")
			}

			farm.Rooms[room.Name] = room

			// Initialize adjacency key (optional but useful)
			if _, ok := farm.Adj[room.Name]; !ok {
				farm.Adj[room.Name] = nil
			}

			// Apply ##start / ##end flags
			if expectStart {
				if farm.Start != "" {
					return nil, fmt.Errorf("ERROR: invalid data format")
				}
				farm.Start = room.Name
				expectStart = false
			} else if expectEnd {
				if farm.End != "" {
					return nil, fmt.Errorf("ERROR: invalid data format")
				}
				farm.End = room.Name
				expectEnd = false
			}
			continue
		}

		// Any other unknown line is invalid
		return nil, fmt.Errorf("ERROR: invalid data format")
	}

	// Scanner error
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ERROR: invalid data format")
	}

	// Must have ants, start, end
	if !foundAnts || farm.Start == "" || farm.End == "" {
		return nil, fmt.Errorf("ERROR: invalid data format")
	}

	// ##start/##end must be followed by a room
	if expectStart || expectEnd {
		return nil, fmt.Errorf("ERROR: invalid data format")
	}

	return farm, nil
}

func isRoomLine(line string) bool {
	return len(strings.Fields(line)) == 3
}

func parseRoom(line string) (Room, error) {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return Room{}, fmt.Errorf("bad room")
	}

	name := parts[0]

	// Room name rules:
	// - cannot start with 'L' or '#'
	// - cannot contain spaces (already handled by Fields)
	if name == "" || strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
		return Room{}, fmt.Errorf("bad room name")
	}

	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return Room{}, fmt.Errorf("bad x")
	}
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return Room{}, fmt.Errorf("bad y")
	}

	return Room{Name: name, X: x, Y: y}, nil
}

func isLinkLine(line string) bool {
	// Must be ONE token and contain exactly one '-'
	if len(strings.Fields(line)) != 1 {
		return false
	}
	return strings.Count(line, "-") == 1
}

func splitLink(line string) (string, string, bool) {
	ab := strings.Split(line, "-")
	if len(ab) != 2 {
		return "", "", false
	}
	a := strings.TrimSpace(ab[0])
	b := strings.TrimSpace(ab[1])
	if a == "" || b == "" {
		return "", "", false
	}
	// No spaces allowed inside names for links
	if strings.ContainsAny(a, " \t") || strings.ContainsAny(b, " \t") {
		return "", "", false
	}
	return a, b, true
}

func normTunnelKey(a, b string) string {
	// Normalize so A-B and B-A are considered the same tunnel
	if a > b {
		a, b = b, a
	}
	return a + "|" + b
}


func main() {
	farm, err := ParseFile("file.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Ants:", farm.Ants)
	fmt.Println("Start:", farm.Start)
	fmt.Println("End:", farm.End)
	fmt.Println("Rooms:", len(farm.Rooms))
}