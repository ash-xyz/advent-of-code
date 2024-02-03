package main

import (
	"strconv"
)

func recursiveFuel(fuel int) int {
	if fuel <= 0 {
		return 0
	}

	return fuel + recursiveFuel((fuel/3)-2)
}

func part2(input []string) {
	ans := 0

	for _, line := range input {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		ans += recursiveFuel((num / 3) - 2)
	}

	println("Part 2: ", ans)
}

func part1(input []string) {
	ans := 0

	for _, line := range input {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		ans += (num / 3) - 2
	}

	println("Part 1: ", ans)
}

func main() {
	// Take its mass, divide by 3; round down and substract 2;
	input := ReadFile("inputs/day1.txt")
	part1(input)
	part2(input)
}
