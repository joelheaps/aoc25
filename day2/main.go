package main

import (
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
)

// Range simply represents a range of product IDs, inclusive
type Range struct {
	first int
	last  int
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
func getRanges(input string) []Range {
	rawRanges := strings.Split(input, ",")
	ranges := make([]Range, len(rawRanges))
	for _, rr := range rawRanges {
		split := strings.Split(rr, "-")
		first, _ := strconv.Atoi(split[0])
		last, _ := strconv.Atoi(split[1])
		ranges = append(ranges, Range{first, last})
	}
	return ranges
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide {part1,part2} and {input_file} as arguments")
		os.Exit(1)
	}
	file := os.Args[2]
	raw_input, _ := os.ReadFile(file)
	input := string(raw_input)
	input = strings.TrimSpace(input)
	ranges := getRanges(input)

	switch os.Args[1] {
	case "part1":
		fmt.Println(sumInvalids(ranges, isValidProductId))
	case "part2":
		fmt.Println(sumInvalids(ranges, isValidProductId2))
	}
}

// isValidProductId returns false when a product IDs first half is the same as its second half
func isValidProductId(i int) bool {
	asStr := strconv.Itoa(i)
	if halfLen := len(asStr) / 2; asStr[:halfLen] == asStr[halfLen:] {
		return false
	}
	return true
}

// isValidProductId2 returns false when a product ID consists entirely of repeated number patterns
func isValidProductId2(i int) bool {
	asStr := strconv.Itoa(i)
	length := len(asStr)
	secLength := 0 // Section length
	firstSection := ""

	// For each evenly divisible section size, check that each section matches the first
Outer:
	for sectionCount := 2; sectionCount <= length; sectionCount++ {
		if length%sectionCount != 0 {
			// If remainder, then this isn't a valid section size
			continue
		}
		secLength = length / sectionCount
		index := secLength
		firstSection = asStr[:index]

		// Check each subsequent section until the end
		for {
			section := asStr[index : index+secLength]
			if section != firstSection {
				// Valid if not repeating; check next section size
				continue Outer
			}
			index += secLength
			if index >= length {
				// Reached end; all sections matched the firstSection
				break
			}
		}
		return false // For this section size, all sections matched
	}
	return true // For all section sizes, at least one section mismatched
}

// sumInvalids returns the sum of all invalid product IDs within the ranges
// provided
func sumInvalids(ranges []Range, isValid func(int) bool) int {
	invalids := []int{}
	for _, r := range ranges {
		for num := range r.enumerate() {
			if !isValid(num) {
				invalids = append(invalids, num)
			}
		}
	}
	sum := 0
	for _, num := range invalids {
		sum += num
	}
	return sum
}
