package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const roll rune = '@'

// Pos represents a row (i) and column (j) in a grid
type Pos struct {
	i int
	j int
}

// Grid represents a 2D grid of runes as rows and columns
type Grid [][]rune

// valueAt simply returns the rune at the given position in the grid
func (g Grid) valueAt(pos Pos) rune {
	return g[pos.i][pos.j]
}

// getNeighbors returns a set (map) of all neighbors of the given position,
// including diagonals.
func (g Grid) getNeighbors(pos Pos) map[Pos]bool {
	neighbors := map[Pos]bool{}
	nJ := 0
	nI := 0
	for _, iShift := range []int{-1, 0, 1} {
		nI = pos.i + iShift
		if nI < 0 || nI >= len(g) {
			continue
		}
		for _, jShift := range []int{-1, 0, 1} {
			nJ = pos.j + jShift
			if nJ < 0 || nJ >= len(g[0]) {
				continue
			}
			if nI == pos.i && nJ == pos.j {
				// Skip own position
				continue
			}
			neighbors[Pos{nI, nJ}] = true
		}
	}
	return neighbors
}

// parseGrid parses the input string into a 2D grid.
func parseGrid(
	input string,
) Grid {
	rows := strings.Split(input, "\n")
	grid := make(Grid, len(rows))

	for i, row := range rows {
		grid[i] = make([]rune, len(row))
		for j, rune := range row {
			grid[i][j] = rune
		}
	}
	return grid
}

// findAccessibleRolls counts rolls with fewer than 4 neighboring rolls,
// optionally removing them as they are counted
func findAccessibleRolls(g *Grid, remove bool) int {
	nRolls := 0
	accessible := 0
	for i, row := range *g {
		for j, v := range row {
			if v != roll {
				continue
			}

			// Check neighbors and count rolls
			nRolls = 0
			for nPos := range g.getNeighbors(Pos{i, j}) {
				if g.valueAt(nPos) == roll {
					nRolls += 1
				}
			}

			// If neighboring rolls < 4, count this roll accessible
			if nRolls < 4 {
				accessible += 1
				if remove {
					(*g)[i][j] = 'x'
				}
			}
		}
	}
	return accessible
}

// findRemovableRolls runs recursively removes accessible rolls until no
// more rolls can be removed, returning the total count removed
func findRemovableRolls(g *Grid) int {
	removed := 0
	rN := 0
	for {
		rN = findAccessibleRolls(g, true)
		if rN == 0 {
			break
		}
		removed += rN
	}
	return removed
}

func main() {
	start := time.Now()

	if len(os.Args) < 3 {
		fmt.Println("Please provide {part1,part2} and {input_file} as arguments")
		os.Exit(1)
	}
	file := os.Args[2]
	raw_input, _ := os.ReadFile(file)
	input := string(raw_input)
	input = strings.TrimSpace(input)
	grid := parseGrid(input)

	switch os.Args[1] {
	case "part1":
		fmt.Println(findAccessibleRolls(&grid, false))
	case "part2":
		fmt.Println(findRemovableRolls(&grid))
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
