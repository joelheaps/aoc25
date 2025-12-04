/*
	INCOMPLETE!

I bailed for a Python solution after some tricky bugs :( I'll circle back eventually
*/
package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Junction struct {
	X float64
	Y float64
	Z float64
}

type empty struct{}

type Circuit map[Junction]bool // Set of junctions
func (c Circuit) add(junction Junction) {
	c[junction] = true
}

// Connect junction boxes closes together
// Work through junction box pairs closes to farthest

func getJunctions(input string) []Junction {
	juncs := []Junction{}
	x := 0
	y := 0
	z := 0
	for l := range strings.SplitSeq(input, "\n") {
		coords := strings.Split(l, ",")
		x, _ = strconv.Atoi(coords[0])
		y, _ = strconv.Atoi(coords[1])
		z, _ = strconv.Atoi(coords[2])
		juncs = append(juncs, Junction{float64(x), float64(y), float64(z)})
	}
	return juncs
}

func initializeCircuits(juncs []Junction) map[Junction]Circuit {
	circuits := map[Junction]Circuit{}
	for _, j := range juncs {
		circuits[j] = Circuit{j: true}
	}
	return circuits
}

func initializeMembers(juncs []Junction) map[Junction]Junction {
	members := map[Junction]Junction{}
	for _, j := range juncs {
		members[j] = j
	}
	return members
}

func calcDist(start, end Junction) float64 {
	return math.Sqrt(math.Pow(start.X-end.X, 2) + math.Pow(start.Y-end.Y, 2) + math.Pow(start.Z-end.Z, 2))
}

func getClosestPairs(junctions []Junction) [][2]Junction {
	nearestNeighbor := map[Junction]Junction{}
	nearestCost := map[Junction]float64{}

	// For every junction, find its nearest neighbor
	for i, ji := range junctions {
		for j, jj := range junctions {
			if i == j {
				continue
			}
			if dist := calcDist(ji, jj); nearestCost[ji] == 0 || dist < nearestCost[ji] {
				nearestNeighbor[ji] = jj
				nearestCost[ji] = dist
			}
		}
	}
	// Deduplicate pairs and convert to slice
	jPairs := [][2]Junction{}
	for junc, nearest := range nearestNeighbor {
		delete(nearestNeighbor, nearest)
		jPairs = append(jPairs, [2]Junction{junc, nearest})
	}
	// Sort pairs
	slices.SortFunc(jPairs, func(pair1, pair2 [2]Junction) int {
		if nearestCost[pair1[0]] < nearestCost[pair2[0]] {
			return -1
		} else {
			return 1
		}
	})
	return jPairs
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
	juncs := getJunctions(input)

	// circuits holds sets of connected junctions
	circuits := initializeCircuits(juncs) // map[Junction]Circuit
	fmt.Println(circuits)
	count := map[int]int{}
	for circuit := range circuits {
		count[len(circuits[circuit])] += 1
	}
	fmt.Println(count)

	// Need to be able determine which circuit a junction is a member of -- members stores this
	circuitMember := initializeMembers(juncs)

	// Sort pairs by nearest -> farthest
	closestPairs := getClosestPairs(juncs)

	// switch os.Args[1] {
	// case "part1":
	// 	fmt.Println(juncs)
	// 	// case "part2":
	// 	// 	fmt.Println(findRemovableRolls(&grid))
	// }

	for _, jPair := range closestPairs[:10] {
		// Combine circuits of poth pairs, putting them in the current circuit of jPair[0]
		leftCircuit := circuitMember[jPair[0]]
		rightCircuit := circuitMember[jPair[1]]
		for junction := range circuits[rightCircuit] {
			// Whatever circuit the right side (jPair[1]) is a member of, combine it with the left's circuit (jPair[0])
			circuits[leftCircuit].add(junction)
			// And mark each as a member of the new circuit
			circuitMember[junction] = leftCircuit
		}
		circuitMember[jPair[1]] = leftCircuit // Add jPair[1] as a member of the destination circuit as well
		delete(circuits, rightCircuit)        // These have been combined, so delete the old right side
	}

	count = map[int]int{}
	for circuit := range circuits {
		count[len(circuits[circuit])] += 1
	}
	fmt.Println(count)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
