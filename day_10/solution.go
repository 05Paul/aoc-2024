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
	trails, _ := trialMap.trails()

	return solution.New(len(trails)), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	trialMap := parse(content)
	_, ratings := trialMap.trails()
	sum := 0
	for _, rating := range ratings {
		sum += rating
	}

	return solution.New(sum), nil
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
				height = -1
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

func (t *trialMap) trails() ([][]position, []int) {
	trails := make([][]position, 0)
	ratings := make([]int, 0)
	for pos := range t.iter() {
		if pos.height == 0 {
			tr, trr := t.trailsFrom(pos)
			trails = append(trails, tr...)
			ratings = append(ratings, trr...)
		}
	}

	return trails, ratings
}

func (t *trialMap) trailsFrom(pos position) ([][]position, []int) {
	ratings := make(map[position]int)
	ratings[pos] = 1
	trails := make(map[position][]position)
	trails[pos] = []position{pos}

	for range 9 {
		newTrails := make(map[position][]position)
		newRatings := make(map[position]int)
		for pos, trail := range trails {
			for neighbor := range t.neighbors(pos) {
				if neighbor.height != pos.height+1 {
					continue
				}

				tmp := make([]position, len(trail)+1)
				copy(tmp, trail)

				tmp[len(trail)] = neighbor
				newTrails[neighbor] = tmp

				newRatings[neighbor] += ratings[pos]

			}
		}
		trails = newTrails
		ratings = newRatings
	}

	tr := make([][]position, 0, len(trails))
	trr := make([]int, 0, len(trails))
	for key, trail := range trails {
		tr = append(tr, trail)
		trr = append(trr, ratings[key])
	}

	return tr, trr
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
