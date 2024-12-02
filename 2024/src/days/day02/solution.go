package day02

import (
	"math"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Part1(input []string) any {
	totalSafe := 0
	for _, line := range input {
		parts := strings.Fields(line)

		prev := -1
		decreasing := false
		increasing := false
		isSafe := true
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			if prev == -1 {
				prev = num
			} else {
				diff := int(math.Abs(float64(num - prev)))
				if diff == 0 || diff > 3 {
					isSafe = false
					break
				}
				if !decreasing && !increasing {
					if num-prev > 0 {
						increasing = true
					} else {
						decreasing = true
					}
				}
				if num-prev < 0 && increasing {
					isSafe = false
					break
				}
				if num-prev > 0 && decreasing {
					isSafe = false
					break
				}
			}
			prev = num
		}
		if isSafe {
			totalSafe++
		}
	}

	return totalSafe
}

func (s *Solution) Part2(input []string) any {
	return 0
}
