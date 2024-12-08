package day06

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Solution struct{}

func (s *Solution) Part1(input []string) any {

	grid := input
	rows := len(grid)
	cols := len(grid[0])

	var guardRow, guardCol int
	var dir Direction

	found := false

	for r := 0; r < rows && !found; r++ {
		for c := 0; c < cols && !found; c++ {
			switch grid[r][c] {
			case '^':
				guardRow, guardCol = r, c
				dir = Up
				found = true
			case 'v':
				guardRow, guardCol = r, c
				dir = Down
				found = true
			case '<':
				guardRow, guardCol = r, c
				dir = Left
				found = true
			case '>':
				guardRow, guardCol = r, c
				dir = Right
				found = true
			}
		}
	}

	dirs := []struct{ dr, dc int }{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	visited := make(map[[2]int]bool)

	visited[[2]int{guardRow, guardCol}] = true

	isObstacle := func(r, c int) bool {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return true
		}
		return grid[r][c] == '#'
	}

	turnRight := func(d Direction) Direction {
		return (d + 1) % 4
	}

	for {
		delta := dirs[dir]
		nextR := guardRow + delta.dr
		nextC := guardCol + delta.dc

		if nextR < 0 || nextR >= rows || nextC < 0 || nextC >= cols {

			break
		}

		if isObstacle(nextR, nextC) {

			dir = turnRight(dir)

			continue
		}

		guardRow, guardCol = nextR, nextC
		visited[[2]int{guardRow, guardCol}] = true
	}

	return len(visited)
}

func (s *Solution) Part2(input []string) any {
	rows := len(input)
	cols := len(input[0])

	mutableGrid := make([][]rune, rows)
	for r := 0; r < rows; r++ {
		mutableGrid[r] = []rune(input[r])
	}

	guardRow, guardCol, dir := findGuardStartAndDirection(mutableGrid)

	dirs := [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	isObstacle := func(r, c int) bool {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return true
		}
		return mutableGrid[r][c] == '#'
	}

	turnRight := func(d int) int {
		return (d + 1) % 4
	}

	stateID := func(r, c, d int) int {
		return (r*cols+c)*4 + d
	}

	candidates := make([][2]int, 0)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if (r != guardRow || c != guardCol) && mutableGrid[r][c] == '.' {
				candidates = append(candidates, [2]int{r, c})
			}
		}
	}

	origStart := mutableGrid[guardRow][guardCol]
	if origStart == '^' || origStart == 'v' || origStart == '<' || origStart == '>' {
		mutableGrid[guardRow][guardCol] = '.'
	}

	canLoopWithObstacle := func(or, oc int) bool {
		mutableGrid[or][oc] = '#'

		visited := make([]bool, rows*cols*4)
		r, c, d := guardRow, guardCol, dir

		for steps := 0; steps < rows*cols*4+1; steps++ {
			sID := stateID(r, c, d)
			if visited[sID] {
				mutableGrid[or][oc] = '.'
				return true
			}
			visited[sID] = true

			dr, dc := dirs[d][0], dirs[d][1]
			nr, nc := r+dr, c+dc

			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				mutableGrid[or][oc] = '.'
				return false
			}

			if isObstacle(nr, nc) {
				d = turnRight(d)
				continue
			}

			r, c = nr, nc
		}

		mutableGrid[or][oc] = '.'
		return false
	}

	count := 0
	for _, cand := range candidates {
		if canLoopWithObstacle(cand[0], cand[1]) {
			count++
		}
	}

	return count
}

func findGuardStartAndDirection(grid [][]rune) (int, int, int) {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			switch grid[r][c] {
			case '^':
				return r, c, 0
			case '>':
				return r, c, 1
			case 'v':
				return r, c, 2
			case '<':
				return r, c, 3
			}
		}
	}
	panic("No guard found")
}
