package day01

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Part1(input []string) any {
	leftArray := make([]int, 0, len(input))
	rightArray := make([]int, 0, len(input))

	for _, line := range input {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			panic("invalid input")
		}
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])

		leftArray = append(leftArray, left)
		rightArray = append(rightArray, right)
	}

	sort.Ints(leftArray)
	sort.Ints(rightArray)

	total := 0
	for i := 0; i < len(leftArray); i++ {
		total += int(math.Abs(float64(leftArray[i] - rightArray[i])))
	}

	return total
}

func (s *Solution) Part2(input []string) any {
	leftArray := make([]int, 0, len(input))
	rightFreq := make(map[int]int)

	for _, line := range input {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			panic("invalid input")
		}
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])

		leftArray = append(leftArray, left)
		rightFreq[right]++
	}

	similarityScore := 0
	for _, left := range leftArray {
		if count, ok := rightFreq[left]; ok {
			similarityScore += left * count
		}
	}

	return similarityScore
}
