package main

import (
	"fmt"
	"strings"
)

type day6 struct {
	input            []string
	adj              map[string][]string
	bidirectionalAdj map[string][]string
	sanOrbits        map[string]struct{}
}

var d6 day6

func (d *day6) init() {
	d.input = ReadFile("inputs/day6.txt")

	d.adj = make(map[string][]string)
	d.bidirectionalAdj = make(map[string][]string)
	d.sanOrbits = make(map[string]struct{}) // Precompute the that SAN orbits for Part2

	for _, orbit := range d.input {
		split_orbit := strings.Split(orbit, ")")
		out, in := split_orbit[0], split_orbit[1]
		d.adj[out] = append(d.adj[out], in)

		d.bidirectionalAdj[out] = append(d.bidirectionalAdj[out], in)
		d.bidirectionalAdj[in] = append(d.bidirectionalAdj[in], out)

		if in == "SAN" {
			d.sanOrbits[out] = struct{}{}
		}
	}

}

func (d *day6) part2() int {
	seen := make(map[string]struct{})

	// Initialize a queue for BFS
	queue := []string{"YOU"}
	seen["YOU"] = struct{}{}

	minPathCount := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			fmt.Println(cur)
			queue = queue[1:]

			if _, ok := d.sanOrbits[cur]; ok {
				return minPathCount - 1
			}

			for _, o := range d.bidirectionalAdj[cur] {
				if _, ok := seen[o]; ok {
					continue
				}
				queue = append(queue, o)
				seen[o] = struct{}{}
			}
		}

		minPathCount++
	}

	return -1
}

func (d *day6) countOrbits(depth int, orbiting string) int {

	sum := 0

	for _, o := range d.adj[orbiting] {
		sum += d.countOrbits(depth+1, o)
	}

	return depth + sum
}

func (d *day6) part1() int {
	return d.countOrbits(0, "COM")
}

func (d day6) run() (int, int) {
	return d.part1(), d.part2()
}
