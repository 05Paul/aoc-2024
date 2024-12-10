package day10

import (
	"aoc/solution"
	"fmt"
	"strconv"
	"strings"
)

func New() solution.Solver {
	return &day{}
}

type day struct{}

func (d *day) SolvePart1(content string) (fmt.Stringer, error) {
	trialMap := parse(content)
	trails := trialMap.trails()

	return solution.New(len(trails)), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func parse(content string) trialMap {
	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]

	grid := make([][]position, 0)
	for y, line := range lines {
		row := make([]position, 0)
		for x, character := range line {
			height, err := strconv.Atoi(string(character))
			if err != nil {
				continue
			}
			row = append(row, position{
				x:      x,
				y:      y,
				height: height,
			})
		}
		grid = append(grid, row)
	}

	return trialMap{
		positions: grid,
	}
}

type trialMap struct {
	positions [][]position
}

func (t *trialMap) trails() [][]position {
	trails := make([][]position, 0)
	for pos := range t.iter() {
		if pos.height == 0 {
			trails = append(trails, t.trailsFrom(pos)...)
		}
	}

	return trails
}

func (t *trialMap) trailsFrom(pos position) [][]position {
	rating := make(map[position]int)
	trails := make(map[position][]position)
	trails[pos] = []position{pos}

	for range 9 {
		newTrails := make(map[position][]position)
		for pos, trail := range trails {
			for neighbor := range t.neighbors(pos) {
				if neighbor.height != pos.height+1 {
					continue
				}

				tmp := make([]position, len(trail)+1)
				copy(tmp, trail)

				tmp[len(trail)] = neighbor
				newTrails[neighbor] = tmp
			}
		}
		trails = newTrails
	}

	tr := make([][]position, 0, len(trails))
	for _, trail := range trails {
		tr = append(tr, trail)
	}

	return tr
}

func (t *trialMap) neighbors(from position) func(func(position) bool) {
	directions := func(moves ...func(position) *position) func(func(position) bool) {
		return func(yield func(position) bool) {
			for _, move := range moves {
				pos := move(from)
				if pos == nil {
					continue
				}

				if !yield(*pos) {
					return
				}
			}
		}
	}
	return directions(t.left, t.right, t.up, t.down)
}

func (t *trialMap) left(from position) *position {
	newX := from.x - 1
	if newX < 0 {
		return nil
	}

	return &t.positions[from.y][newX]
}

func (t *trialMap) right(from position) *position {
	newX := from.x + 1
	if newX >= len(t.positions[from.y]) {
		return nil
	}

	return &t.positions[from.y][newX]
}

func (t *trialMap) up(from position) *position {
	newY := from.y - 1
	if newY < 0 {
		return nil
	}

	return &t.positions[newY][from.x]
}

func (t *trialMap) down(from position) *position {
	newY := from.y + 1
	if newY >= len(t.positions) {
		return nil
	}

	return &t.positions[newY][from.x]
}

func (t *trialMap) iter() func(func(position) bool) {
	return func(yield func(position) bool) {
		for _, row := range t.positions {
			for _, position := range row {
				if !yield(position) {
					return
				}
			}
		}
	}
}

type position struct {
	x      int
	y      int
	height int
}

type trailRoute struct {
	start position
	end   position
}
