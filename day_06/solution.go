package day06

import (
	"aoc/solution"
	"fmt"
	"strings"
)

func New() solution.Solver {
	return &day{}
}

type day struct{}

func (d *day) SolvePart1(content string) (fmt.Stringer, error) {
	x, y, grid := parse(content)
	_, visited := solve(x, y, 0, grid)

	count := len(visited)
	return solution.New(count), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	x, y, grid := parse(content)
	_, visited := solve(x, y, 0, grid)

	count := 0
	for position := range visited {
		if position.x == x && position.y == y {
			continue
		}

		grid[position.y][position.x] = true

		loop, _ := solve(x, y, 0, grid)
		if loop {
			count += 1
		}

		grid[position.y][position.x] = false
	}

	return solution.New(count), nil
}

func parse(content string) (x int, y int, grid [][]bool) {
	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]
	grid = make([][]bool, 0)

	for row, line := range lines {
		currentRow := make([]bool, 0)
		for column, character := range line {
			if character == '^' {
				x = column
				y = row
			}
			currentRow = append(currentRow, character == '#')
		}

		grid = append(grid, currentRow)
	}

	return x, y, grid
}

func inbounds(x int, y int, size int) bool {
	inbound := func(coordinate int) bool {
		return coordinate >= 0 && coordinate < size
	}

	return inbound(x) && inbound(y)
}

func move(x int, y int, direction int) (newX int, newY int) {
	switch direction {
	case 0:
		return x, y - 1
	case 1:
		return x + 1, y
	case 2:
		return x, y + 1
	case 3:
		return x - 1, y
	default:
		return x, y
	}
}

func solve(x int, y int, direction int, grid [][]bool) (loop bool, positions map[position]bool) {
	traversed := make(map[traversal]bool)
	traversed[traverse(x, y, direction)] = true
	for {
		newX, newY := move(x, y, direction)
		if !inbounds(newX, newY, len(grid)) {
			break
		}

		if _, exists := traversed[traverse(newX, newY, direction)]; exists {
			return true, nil
		}

		if grid[newY][newX] {
			direction = (direction + 1) % 4
			continue
		}

		x, y = newX, newY

		traversed[traverse(x, y, direction)] = true
	}

	return false, simpleMap(traversed)
}

func simpleMap(visited map[traversal]bool) map[position]bool {
	out := make(map[position]bool, len(visited))
	for traversal := range visited {
		out[traversal.pos] = true
	}

	return out
}

type position struct {
	x int
	y int
}

func at(x int, y int) position {
	return position{
		x: x,
		y: y,
	}
}

type traversal struct {
	pos position
	dir int
}

func traverse(x int, y int, dir int) traversal {
	return traversal{
		pos: position{
			x: x,
			y: y,
		},
		dir: dir,
	}
}
