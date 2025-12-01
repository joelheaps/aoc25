package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide {part1,part2} and {input_file} as arguments")
		os.Exit(1)
	}
	file := os.Args[2]
	raw_input, _ := os.ReadFile(file)
	input := string(raw_input)
	input = strings.TrimSpace(input)
	rotations := getRotations(input)

	switch os.Args[1] {
	case "part1":
		part1(rotations)
	case "part2":
		part2(rotations)
	}
}

func getRotations(input string) []int {
	result := []int{}
	val := 0
	for line := range strings.SplitSeq(input, "\n") {
		switch line[0] {
		case 'R':
			val, _ = strconv.Atoi(line[1:])
			result = append(result, val)
		case 'L':
			val, _ = strconv.Atoi(line[1:])
			result = append(result, val*-1)
		}
	}
	return result
}

func part1(rotations []int) {
	selected := 50
	zeroCount := 0
	for _, r := range rotations {
		selected += r
		if selected < 0 {
			selected = selected + 100*(selected/-100+1)
		}
		if selected > 99 {
			selected = selected % 100
		}
		if selected == 0 {
			zeroCount += 1
		}
	}
	fmt.Println(zeroCount)
}

// I started on a "smarter" implementation of part 2, but remainders and passing-zero counts were getting
// tricky.  The click-by-one for loop was easier to reason about
func part2(rotations []int) {
	selected := 50
	dir := -1
	zeroCount := 0
	for _, r := range rotations {
		if r < 0 {
			dir = -1
		} else {
			dir = 1
		}

		for i := 0; i != r; i += dir {
			selected += dir
			if selected == 100 {
				selected = 0
			}
			if selected == -1 {
				selected = 99
			}
			if selected == 0 {
				zeroCount += 1
			}
		}
	}
	fmt.Println(zeroCount)
}
