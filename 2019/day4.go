package main

type day4 struct {
	lower int
	upper int
}

var d4 day4

func (d *day4) init() {
	d.lower = 367479 // 367789
	d.upper = 893698 // 899999
}

func neverDecreases(num int) bool {
	prev := MaxInt

	for num > 0 {
		x := num % 10
		if x > prev {
			return false
		}
		num /= 10
		prev = x
	}

	return true
}

func hasAdjacent(num int) bool {
	prev := 0

	for num > 0 {
		x := num % 10
		if x == prev {
			return true
		}
		num /= 10
		prev = x
	}
	return false
}

func hasAdjacent2(num int) bool {
	prev := 0
	consecutiveCount := make(map[int]int)

	for num > 0 {
		x := num % 10
		if x == prev {
			consecutiveCount[prev]++
		} else {
			if consecutiveCount[prev] == 1 {
				return true
			}
			consecutiveCount[prev] = 0
		}
		num /= 10
		prev = x
	}
	if consecutiveCount[prev] == 1 {
		return true
	}
	return false
}

func (d *day4) part2() int {
	count := 0
	for num := d.lower; num <= d.upper; num++ {
		if neverDecreases(num) && hasAdjacent2(num) {
			count++
		}
	}
	return count
}

func (d *day4) part1() int {
	count := 0
	for num := d.lower; num <= d.upper; num++ {
		if neverDecreases(num) && hasAdjacent(num) {
			count++
		}
	}
	return count
}

func (d *day4) run() (int, int) {
	return d.part1(), d.part2()
}
