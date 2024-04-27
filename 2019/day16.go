package main

import (
	"strconv"
)

type day16 struct {
	input string
}

var d16 day16

func (d *day16) init() {
	input := ReadFile("inputs/day16.txt")
	d.input = input[0]
}

func (d *day16) part2() int {
	offset, _ := strconv.Atoi(d.input[:7])
	digits := make([]int, len(d.input)*10000)
	for i := range digits {
		digits[i], _ = strconv.Atoi(string(d.input[i%len(d.input)]))
	}

	for phase := 0; phase < 100; phase++ {
		sum := 0
		for i := len(digits) - 1; i >= offset; i-- {
			sum += digits[i]
			digits[i] = sum % 10
		}
	}

	result := 0
	for i := offset; i < offset+8; i++ {
		result = result*10 + digits[i]
	}
	return result
}

func (d *day16) part1() int {
	digits := make([]int, len(d.input))
	for i, c := range d.input {
		digits[i], _ = strconv.Atoi(string(c))
	}

	for phase := 0; phase < 100; phase++ {
		newDigits := make([]int, len(digits))
		for i := range digits {
			sum := 0
			for j := range digits {
				multiplier := getMultiplier(i, j)
				sum += digits[j] * multiplier
			}
			newDigits[i] = abs(sum) % 10
		}
		digits = newDigits
	}

	result := 0
	for i := 0; i < 8; i++ {
		result = result*10 + digits[i]
	}
	return result
}

func getMultiplier(i, j int) int {
	pattern := []int{0, 1, 0, -1}
	index := ((j + 1) / (i + 1)) % len(pattern)
	return pattern[index]
}

func (d day16) run() (int, int) {
	return d.part1(), d.part2()
}
