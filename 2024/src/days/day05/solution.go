package day05

import (
	"container/heap"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Part1(input []string) any {
	rules, updates := parse(input)
	total := 0
	for _, update := range updates {
		if isCorrectOrder(update, rules) {
			midIndex := len(update) / 2
			total += update[midIndex]
		}
	}
	return total
}

func (s *Solution) Part2(input []string) any {
	rules, updates := parse(input)
	total := 0

	for _, update := range updates {
		if !isCorrectOrder(update, rules) {
			corrected := reorderUpdate(update, rules)
			midIndex := len(corrected) / 2
			total += corrected[midIndex]
		}
	}

	return total
}

func parse(lines []string) (rules [][2]int, updates [][]int) {
	var ruleLines []string
	var updateLines []string
	collectingRules := true
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			collectingRules = false
			continue
		}
		if collectingRules {
			if strings.Contains(line, "|") {
				ruleLines = append(ruleLines, line)
			} else {
				collectingRules = false
				if line != "" {
					updateLines = append(updateLines, line)
				}
			}
		} else {
			updateLines = append(updateLines, line)
		}
	}

	for _, r := range ruleLines {
		parts := strings.Split(r, "|")
		if len(parts) == 2 {
			X, errX := strconv.Atoi(strings.TrimSpace(parts[0]))
			Y, errY := strconv.Atoi(strings.TrimSpace(parts[1]))
			if errX == nil && errY == nil {
				rules = append(rules, [2]int{X, Y})
			}
		}
	}

	for _, u := range updateLines {
		pagesStr := strings.Split(u, ",")
		var pages []int
		for _, p := range pagesStr {
			num, err := strconv.Atoi(strings.TrimSpace(p))
			if err == nil {
				pages = append(pages, num)
			}
		}
		if len(pages) > 0 {
			updates = append(updates, pages)
		}
	}

	return rules, updates
}

func isCorrectOrder(update []int, rules [][2]int) bool {
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}

	for _, rule := range rules {
		X, Y := rule[0], rule[1]
		idxX, okX := position[X]
		idxY, okY := position[Y]
		if okX && okY {
			if idxX >= idxY {
				return false
			}
		}
	}
	return true
}

func reorderUpdate(update []int, rules [][2]int) []int {
	pagesSet := make(map[int]bool)
	for _, p := range update {
		pagesSet[p] = true
	}

	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	for _, p := range update {
		inDegree[p] = 0
		graph[p] = []int{}
	}

	for _, rule := range rules {
		X, Y := rule[0], rule[1]
		if pagesSet[X] && pagesSet[Y] {
			graph[X] = append(graph[X], Y)
			inDegree[Y]++
		}
	}

	minHeap := &IntHeap{}
	heap.Init(minHeap)

	for p, deg := range inDegree {
		if deg == 0 {
			heap.Push(minHeap, p)
		}
	}

	var result []int
	for minHeap.Len() > 0 {
		x := heap.Pop(minHeap).(int)
		result = append(result, x)
		for _, nbr := range graph[x] {
			inDegree[nbr]--
			if inDegree[nbr] == 0 {
				heap.Push(minHeap, nbr)
			}
		}
	}

	return result
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
