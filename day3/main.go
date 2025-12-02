package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type BatteryBank []int

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

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

type battery struct {
	pos   int
	power int
}

func getHighestCombo(bb BatteryBank, batteryCount int) int {
	sb := make([]battery, batteryCount) // selectedBatteries
	for i, val := range bb {
		for batNum, bat := range sb {
			minDistFromEnd := batteryCount - batNum
			// Check that this battery is at least n positions from end
			if i > len(bb)-minDistFromEnd {
				continue
			}
			// Validate this battery's position is greater than previous's; move it
			// forward if not. Otherwise, move it forward if this position's power
			// is greater than current selection.
			if (batNum > 0 && bat.pos <= sb[batNum-1].pos) || val > bat.power {
				bat.pos = i
				bat.power = val
				sb[batNum] = bat
			}
		}
	}

	sum := 0
	for i, bat := range sb {
		sum += bat.power * powInt(10, batteryCount-i-1)
	}
	return sum
}

func sumHighestJoltages(bbanks []BatteryBank, batteryCount int) int {
	sum := 0
	this := 0
	for _, bb := range bbanks {
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
