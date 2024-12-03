package day03

import (
	"regexp"
	"strconv"
)

type Solution struct{}

func (s *Solution) Part1(input []string) any {
	regex := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	sum := 0
	for _, line := range input {
		matches := regex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			left, _ := strconv.Atoi(match[1])
			right, _ := strconv.Atoi(match[2])
			sum += left * right
		}
	}
	return sum
}

func (s *Solution) Part2(input []string) any {
	regex := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)|don't\(\)|do\(\)`)
	sum := 0
	enabled := true

	for _, line := range input {
		matches := regex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[0] == "don't()" {
				enabled = false
			}
			if match[0] == "do()" {
				enabled = true
			} else if enabled {
				left, _ := strconv.Atoi(match[1])
				right, _ := strconv.Atoi(match[2])
				sum += left * right
			}
		}
	}
	return sum
}
