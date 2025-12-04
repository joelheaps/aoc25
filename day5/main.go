package main

import (
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
	"time"
)

// Range simply represents a range of product IDs, inclusive
type Range struct {
	first int
	last  int
}

func (r *Range) isInRange(i int) bool {
	return i >= r.first && i <= r.last
}

// enumerate yields all product IDs contained in the Range
func (r *Range) enumerate() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := r.first; i <= r.last; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// getRanges parses ranges of product IDs from input data
func parseRangesAndIds(input string) ([]Range, []int) {
	parts := strings.Split(input, "\n\n") // First part ranges, second part IDs

	// Split ranges
	rawRanges := strings.Split(parts[0], "\n")
	ranges := make([]Range, len(rawRanges))
	ranges = ranges[:0]
	for _, rr := range rawRanges {
		split := strings.Split(rr, "-")
		first, _ := strconv.Atoi(split[0])
		last, _ := strconv.Atoi(split[1])
		ranges = append(ranges, Range{first, last})
	}

	// Split IDs
	rawIds := strings.Split(parts[1], "\n")
	ids := make([]int, len(rawIds))
	ids = ids[:0]
	for _, ri := range rawIds {
		id, _ := strconv.Atoi(strings.TrimSpace(ri))
		ids = append(ids, id)
	}

	return ranges, ids
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
	ranges, ids := parseRangesAndIds(input)

	switch os.Args[1] {
	case "part1":
		fmt.Println(sumFresh(ranges, ids))
	case "part2":
		fmt.Println(sumAllFresh(ranges))
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}

// sumInvalids returns the sum of all invalid product IDs within the ranges
// provided
func sumFresh(ranges []Range, ids []int) int {
	freshCount := 0
Outer:
	for _, id := range ids {
		for _, r := range ranges {
			if r.isInRange(id) {
				freshCount += 1
				continue Outer

			}
		}
	}
	return freshCount
}

func combineOverlapping(ranges []Range) []Range {
	combined := false
	for {
		combined = false // Reset every loop

		for i := 0; i < len(ranges); i++ {
			for j := i + 1; j < len(ranges); j++ {
				r := ranges[i]
				rCompare := ranges[j]

				if rCompare.isInRange(r.first) || rCompare.isInRange(r.last) ||
					rCompare.isInRange(r.last+1) || rCompare.isInRange(r.first-1) ||
					r.isInRange(rCompare.first) || r.isInRange(rCompare.last) {

					first := min(r.first, rCompare.first)
					last := max(r.last, rCompare.last)

					// Create new slice with combined range
					newRanges := make([]Range, 0, len(ranges)-1)
					newRanges = append(newRanges, Range{first, last})

					// Add all ranges except i and j
					for k := 0; k < len(ranges); k++ {
						if k != i && k != j {
							newRanges = append(newRanges, ranges[k])
						}
					}

					ranges = newRanges
					combined = true
					break
				}
			}
			if combined {
				break
			}
		}

		if !combined {
			// No more overlaps found
			break
		}
	}

	return ranges
}

func sumAllFresh(ranges []Range) int {
	count := 0
	ranges = combineOverlapping(ranges)
	for _, r := range ranges {
		count += r.last - r.first + 1
	}
	return count
}
