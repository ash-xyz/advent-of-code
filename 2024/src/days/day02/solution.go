package day02

import (
	"math"
	"slices"
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
	totalSafe := 0
	for _, line := range input {
		parts := strings.Fields(line)
		nums := make([]int, len(parts))
		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			nums[i] = num
		}

		if IsSafe(nums) {
			totalSafe++
			continue
		}

		isSafe := false
		for i := 0; i < len(nums); i++ {
			newNums := slices.Delete(slices.Clone(nums), i, i+1)
			if IsSafe(newNums) {
				isSafe = true
				break
			}
		}

		if isSafe {
			totalSafe++
		}
	}

	return totalSafe
}

func IsSafe(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	dir := 0 // 1 for increasing, -1 for decreasing
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if abs(diff) == 0 || abs(diff) > 3 {
			return false
		}
		if dir == 0 {
			if diff > 0 {
				dir = 1
			} else if diff < 0 {
				dir = -1
			}
		} else {
			if diff*dir <= 0 {
				return false
			}
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
