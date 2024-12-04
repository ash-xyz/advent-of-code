package day04

type Solution struct{}

const xmas = "XMAS"

func getXmasCount(input []string, x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			sXLocation, sYLocation := x+3*i, y+3*j
			if sXLocation >= len(input) || sXLocation < 0 || sYLocation >= len(input[0]) || sYLocation < 0 {
				continue
			}

			isXmas := true
			for idx, c := range xmas {
				if c != rune(input[x+idx*i][y+idx*j]) {
					isXmas = false
					break
				}
			}
			if isXmas {
				count++
			}
		}
	}
	return count
}

func (s *Solution) Part1(input []string) any {
	n, m := len(input), len(input[0])
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sum += getXmasCount(input, i, j)
		}
	}
	return sum
}

func getXmasCountV2(input []string, x, y int) int {
	if input[x][y] != byte('A') {
		return 0
	}

	// Check that all corners fit
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			sXLocation, sYLocation := x+i, y+j
			if sXLocation >= len(input) || sXLocation < 0 || sYLocation >= len(input[0]) || sYLocation < 0 {
				return 0
			}
		}
	}

	if input[x-1][y-1] != byte('S') && input[x-1][y-1] != byte('M') {
		return 0
	}

	if input[x+1][y+1] != byte('S') && input[x+1][y+1] != byte('M') {
		return 0
	}

	if input[x-1][y+1] != byte('S') && input[x-1][y+1] != byte('M') {
		return 0
	}

	if input[x+1][y-1] != byte('S') && input[x+1][y-1] != byte('M') {
		return 0
	}

	if input[x+1][y-1] == input[x-1][y+1] {
		return 0
	}

	if input[x+1][y+1] == input[x-1][y-1] {
		return 0
	}

	return 1
}

func (s *Solution) Part2(input []string) any {
	n, m := len(input), len(input[0])
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sum += getXmasCountV2(input, i, j)
		}
	}
	return sum
}
