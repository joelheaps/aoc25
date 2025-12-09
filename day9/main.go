package main

// the Elves would like to find the largest rectangle that uses red tiles for two of its opposite corners

import (
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

type Pos [2]int
type Extents [2]int
type empty struct{}

func getCorners(input string) ([]Pos, Extents) {
	corners := []Pos{}
	i := 0
	j := 0
	maxI, maxJ := 0, 0

	for l := range strings.SplitSeq(input, "\n") {
		coords := strings.Split(l, ",")
		i, _ = strconv.Atoi(coords[1])
		maxI = max(maxI, i)
		j, _ = strconv.Atoi(coords[0])
		maxJ = max(maxJ, j)
		corners = append(corners, [2]int{i, j})
	}
	return corners, Extents{maxI, maxJ}
}

func getBiggestRect(corners []Pos) int {
	biggest := 0
	for _, r1 := range corners {
		for _, r2 := range corners {
			area := int(math.Abs(float64((r1[0]-r2[0])+1) * math.Abs(float64(r1[1]-r2[1])+1)))
			if area > biggest {
				biggest = area
			}
		}
	}
	return biggest
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
	corners, extents := getCorners(input)

	switch os.Args[1] {
	case "part1":
		fmt.Println(getBiggestRect(corners))
	case "part2":
		part2(corners, extents)
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}

// Verify that exiting midpoints of rectangle towards edge of grid
// crosses a border tile in every direction, exactly once
func rectIsInsideBorder(topL, btmR Pos, border map[Pos]empty, extents Extents) bool {
	mid := Pos{randRange(topL[0], btmR[0]), randRange(topL[1], btmR[1])}

	// Check exiting from top crosses border
	crossed := 0
	for i := mid[0]; i > 0; i-- {
		if _, ok := border[Pos{i, mid[1]}]; ok {
			crossed += 1
		}
	}
	if crossed%2 != 1 {
		return false
	}

	// Check exiting from left crosses border
	for j := mid[1]; j > 0; j-- {
		if _, ok := border[Pos{mid[0], j}]; ok {
			crossed += 1
		}
	}
	if crossed%2 != 0 {
		return false
	}

	// Check exiting from bottom crosses border
	for i := mid[0]; i < extents[0]; i++ {
		if _, ok := border[Pos{i, mid[1]}]; ok {
			crossed += 1
		}
	}
	if crossed%2 != 1 {
		return false
	}

	// Check exiting from right crosses border
	for j := mid[1]; j < extents[1]; j++ {
		if _, ok := border[Pos{mid[0], j}]; ok {
			crossed += 1
		}
	}
	if crossed%2 != 0 {
		return false
	}
	return true
}

func intersects(corner1, corner2 Pos, border map[Pos]empty, extents Extents) bool {
	topL := Pos{min(corner1[0], corner2[0]), min(corner1[1], corner2[1])} // top left corner
	btmR := Pos{max(corner1[0], corner2[0]), max(corner1[1], corner2[1])} // bottom right corner
	// Inset by one space; rectangle and shape border can share a line
	for p := range border { // border position
		if p[0] > topL[0] && p[0] < btmR[0] && p[1] > topL[1] && p[1] < btmR[1] {
			// Border crosses into rectangle (not just sharing a border)
			return true
		}
	}
	if topL[0] == btmR[0] || topL[1] == btmR[1] {
		// Just don't want a skinny box
		return true
	}
	// Add additional check that rectangle seems to be contained by shape
	return !rectIsInsideBorder(topL, btmR, border, extents)
}

func part2(corners []Pos, extents Extents) {
	border := map[Pos]empty{}

	// Mark corners and connect them
	last := corners[len(corners)-1]
	for _, c := range corners {
		border[c] = empty{}
		// Fill i (vertical)
		if result := last[0] - c[0]; result != 0 {
			step := 1 // Move down
			if result > 0 {
				step = -1 // Move up
			}
			for i := last[0]; i != c[0]; i += step {
				border[Pos{i, c[1]}] = empty{}
			}

		}

		// Or, fill j (horizontal)
		if result := last[1] - c[1]; result != 0 {
			step := 1 // Move right
			if result > 0 {
				step = -1 // Move left
			}
			for j := last[1]; j != c[1]; j += step {
				border[Pos{c[0], j}] = empty{}
			}
		}
		last = c
	}

	biggest := 100
	for _, r1 := range corners {
		for _, r2 := range corners {
			area := int(math.Abs(float64((r1[0]-r2[0])+1) * math.Abs(float64(r1[1]-r2[1])+1)))
			if area > biggest && !intersects(r1, r2, border, extents) {
				fmt.Println(area)
				biggest = area
			}
		}
	}
	fmt.Println(biggest)
}
