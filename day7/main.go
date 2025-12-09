package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

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

func (g Grid) set(pos Pos, r rune) {
	g[pos.i][pos.j] = r
}

func (g Grid) print() {
	for _, row := range g {
		fmt.Println(string(row))
	}
}

// getNeighbors returns a set (map) of all neighbors of the given position,
// including diagonals.
func (g *Grid) getLRNeighbors(pos Pos) []Pos {
	neighbors := []Pos{}
	if nJ := pos.j + 1; nJ < len((*g)[0]) {
		neighbors = append(neighbors, Pos{pos.i, nJ})
	}
	if nJ := pos.j - 1; nJ > 0 {
		neighbors = append(neighbors, Pos{pos.i, nJ})
	}
	return neighbors
}

func (g *Grid) getDNeighbor(pos Pos) (Pos, error) {
	if pos.i+1 < len(*g) {
		return Pos{pos.i + 1, pos.j}, nil
	}
	return Pos{}, fmt.Errorf("exited grid")
}

// parseGrid parses the input string into a 2D grid.
func parseGrid(
	input string,
) (Grid, Pos) {
	rows := strings.Split(input, "\n")
	grid := make(Grid, len(rows))
	start := Pos{}

	for i, row := range rows {
		grid[i] = make([]rune, len(row))
		for j, rune := range row {
			grid[i][j] = rune
			if rune == 'S' {
				start = Pos{i, j}
			}
		}
	}
	return grid, start
}

// func beam(g *Grid, here Pos, acc int, cache map[Pos]int) int {
// 	if val, ok := cache[here]; ok {
// 		return val
// 	}
// 	next, err := g.getDNeighbor(here)
// 	if err != nil {
// 		return acc
// 	}
// 	// if g.valueAt(next) == '|' {
// 	// 	return acc
// 	// }
// 	// g.set(here, '|')
// 	if g.valueAt(next) == '^' {
// 		acc += 1
// 		split := g.getLRNeighbors(next)
// 		for _, pos := range split {
// 			// if g.valueAt(pos) != '|' {
// 			// g.set(pos, '|')
// 			acc += beam(g, pos, )
// 			// }
// 		}
// 		return acc
// 	}
// 	return beam(g, next, acc)
// }

func beam(g *Grid, here Pos, acc int, cache map[Pos]int) int {
	if val, ok := cache[here]; ok {
		return val
	}
	next, err := g.getDNeighbor(here)
	if err != nil {
		return acc
	}
	if g.valueAt(next) == '^' {
		acc += 1
		split := g.getLRNeighbors(next)
		for _, pos := range split {
			acc += beam(g, pos, 0, cache)
		}
		cache[here] = acc
		return acc
	}
	result := beam(g, next, acc, cache)
	cache[here] = result
	return result
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
	grid, startPos := parseGrid(input)

	switch os.Args[1] {
	case "part1":
		cache := map[Pos]int{}
		fmt.Println(beam(&grid, startPos, 1, cache))
		grid.print()
		// case "part2":
		// 	fmt.Println(findRemovableRolls(&grid))
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
