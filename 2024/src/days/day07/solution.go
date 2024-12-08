package day07

import (
	"strconv"
	"strings"
)

type Solution struct{}

func isPossible(target int, cur int, input []int) bool {
	if cur > target {
		return false
	}

	if len(input) == 0 {
		return target == cur
	}

	return isPossible(target, cur*input[0], input[1:]) || isPossible(target, cur+input[0], input[1:])
}

func intLen(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func concat(num1, num2 int) int {
	n := intLen(num2)
	pow := 1
	for i := 0; i < n; i++ {
		pow *= 10
	}
	return num1*pow + num2
}

func isPossibleWithConcat(target, cur int, nums []int) bool {
	if cur > target {
		return false
	}
	if len(nums) == 0 {
		return cur == target
	}

	if isPossibleWithConcat(target, cur*nums[0], nums[1:]) {
		return true
	}

	if isPossibleWithConcat(target, cur+nums[0], nums[1:]) {
		return true
	}

	return isPossibleWithConcat(target, concat(cur, nums[0]), nums[1:])
}

func (s *Solution) Part1(input []string) any {
	total := 0
	for _, equation := range input {
		split := strings.Split(equation, ":")
		value, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}

		arr := strings.Fields(split[1])
		nums := make([]int, len(arr))
		for i, num := range arr {
			target, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			nums[i] = target
		}

		if len(nums) > 0 && isPossible(value, nums[0], nums[1:]) {
			total += value
		}
	}

	return total
}

func (s *Solution) Part2(input []string) any {
	total := 0
	for _, equation := range input {
		split := strings.Split(equation, ":")
		target, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}

		arr := strings.Fields(split[1])
		nums := make([]int, len(arr))
		for i, num := range arr {
			value, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			nums[i] = value
		}

		if len(nums) > 0 && isPossibleWithConcat(target, nums[0], nums[1:]) {
			total += target
		}
	}

	return total
}
