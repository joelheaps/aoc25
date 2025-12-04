package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

// BatteryBank represents a set of (unconnected) batteries
type BatteryBank []int

// SelectedBattery represents a connected ("on") battery
type SelectedBattery struct {
	pos   int
	power int
}

// getBanks parses battery banks from the input, returning them as a list of lists
func getBanks(input string) []BatteryBank {
	banks := []BatteryBank{}
	for line := range strings.SplitSeq(input, "\n") {
		bank := BatteryBank{}
		for _, r := range line {
			bank = append(bank, int(r-'0'))
		}
		banks = append(banks, bank)
	}
	return banks
}

// powInt is a simple utility function for raising integers to a power (used to calculate total "joltages")
func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// getHighestCombo returns the highest possible "joltage" produceable from the given bank using the number
// of batteries specified.
func getHighestCombo(bBank BatteryBank, batteryCount int) int {
	sel := make([]SelectedBattery, batteryCount)
	minDistFromEnd := 0
	for i, val := range bBank {
		for batNum, bat := range sel {
			minDistFromEnd = batteryCount - batNum
			if i > len(bBank)-minDistFromEnd {
				// Leave room for remaining batteries
				continue
			}
			// If current selection for this battery is left of the previous battery, it needs
			// to be moved forward. Otherwise, move it forward if this position's power
			// is greater than current selection.
			if (batNum > 0 && bat.pos <= sel[batNum-1].pos) || val > bat.power {
				bat.pos = i
				bat.power = val
				sel[batNum] = bat
			}
		}
	}

	joltageSum := 0
	for i, bat := range sel {
		joltageSum += bat.power * powInt(10, batteryCount-i-1)
	}
	return joltageSum
}

// sumHighestJoltages returns the sum of best-case joltages across all battery banks
func sumHighestJoltages(bBanks []BatteryBank, batteryCount int) int {
	sum := 0
	this := 0
	for _, bb := range bBanks {
		this = getHighestCombo(bb, batteryCount)
		sum += this
	}
	return sum
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
	bbanks := getBanks(input)

	switch os.Args[1] {
	case "part1":
		fmt.Println(sumHighestJoltages(bbanks, 2))
	case "part2":
		fmt.Println(sumHighestJoltages(bbanks, 12))
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
