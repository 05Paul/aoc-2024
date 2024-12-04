package day04

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
	grid := parse(content)
	count := search([]rune("XMAS"), grid)
	return solution.New(count), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func parse(content string) [][]rune {
	grid := make([][]rune, 0)
	rowLength := 0
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		row := make([]rune, 0, rowLength)
		for _, character := range line {
			row = append(row, character)
		}

		rowLength = len(row)
		if rowLength != 0 {
			grid = append(grid, row)
		}
	}

	return grid
}

func search(word []rune, grid [][]rune) int {
	count := 0
	for y, row := range grid {
		for x, character := range row {
			if character == word[0] {
				count += dfs(word, grid, x, y)
			}
		}
	}

	return count
}

func dfs(word []rune, grid [][]rune, x int, y int) int {
	sum := 0
	for direction := range 8 {
		sum += dfsDirection(word, grid, x, y, direction)
	}

	return sum
}

func dfsDirection(word []rune, grid [][]rune, x int, y int, direction int) int {
	xDir, yDir := directionPair(direction)
	for offset, character := range word {
		newX := x + offset*xDir
		newY := y + offset*yDir
		if newY < 0 || newY >= len(grid) {
			return 0
		}

		if newX < 0 || newX >= len(grid[0]) {
			return 0
		}

		if character != grid[newY][newX] {
			return 0
		}
	}

	return 1
}

func directionPair(direction int) (xDir int, yDir int) {
	switch direction {
	case 0:
		return 1, 1
	case 1:
		return 0, 1
	case 2:
		return -1, 1
	case 3:
		return -1, 0
	case 4:
		return -1, -1
	case 5:
		return 0, -1
	case 6:
		return 1, -1
	case 7:
		return 1, 0
	default:
		return 0, 0
	}
}
