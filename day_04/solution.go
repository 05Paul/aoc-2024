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
	grid := parse(content)
	count := searchMas(grid)
	return solution.New(count), nil
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
		if current := getOffset(grid, x, y, offset*xDir, offset*yDir); current == nil || character != *current {
			return 0
		}
	}

	return 1
}

func getOffset(grid [][]rune, x int, y int, xOff int, yOff int) *rune {
	newX := x + xOff
	newY := y + yOff
	if !inbounds(grid, newY) || !inbounds(grid[0], newX) {
		return nil
	}

	return &grid[newY][newX]
}

func inbounds[T any](axis []T, index int) bool {
	return index >= 0 && index < len(axis)
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

func searchMas(grid [][]rune) int {
	count := 0
	for y, row := range grid {
		for x, character := range row {
			if character == 'A' {
				count += isXMas(grid, x, y)
			}
		}
	}

	return count
}

func isXMas(grid [][]rune, x int, y int) int {
	topLeft := getOffset(grid, x, y, -1, 1)
	topRight := getOffset(grid, x, y, 1, 1)
	bottomLeft := getOffset(grid, x, y, -1, -1)
	bottomRight := getOffset(grid, x, y, 1, -1)

	if topLeft == nil || topRight == nil || bottomLeft == nil || bottomRight == nil {
		return 0
	}

	if masAxis(*topLeft, *bottomRight) && masAxis(*topRight, *bottomLeft) {
		return 1
	} else {
		return 0
	}
}

func masAxis(a rune, b rune) bool {
	return (a == 'M' && b == 'S') || (a == 'S' && b == 'M')
}
