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
	direction := 0

	positions := []int{linear(x, y, len(grid))}
	for {
		newX, newY := move(x, y, direction)
		if !inbounds(newX, newY, len(grid)) {
			break
		}

		if grid[newY][newX] {
			direction = (direction + 1) % 4
			continue
		}
		x, y = newX, newY
		positions = append(positions, linear(x, y, len(grid)))
	}

	count := len(removeDuplicates(positions))

	return solution.New(count), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	return nil, fmt.Errorf("Not yet implemented")
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

func linear(x int, y int, size int) int {
	return y*size + x
}

func inbounds(x int, y int, size int) bool {
	inbound := func(coordinate int) bool {
		return coordinate >= 0 && coordinate < size
	}

	return inbound(x) && inbound(y)
}

func removeDuplicates(positions []int) []int {
	values := make(map[int]bool)
	list := make([]int, 0)
	for _, value := range positions {
		if _, exists := values[value]; !exists {
			values[value] = true
			list = append(list, value)
		}
	}

	return list
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
