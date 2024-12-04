package day02

import (
	"aoc/solution"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func New() solution.Solver {
	return &day{}
}

type day struct{}

func (d *day) SolvePart1(content string) (fmt.Stringer, error) {
	reports, err := parse(content)
	if err != nil {
		return nil, err
	}

	safeCount := 0
	for _, report := range reports {
		if isSafe(report, 0) {
			safeCount += 1
		}
	}

	return solution.Solve(safeCount), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	reports, err := parse(content)
	if err != nil {
		return nil, err
	}

	safeCount := 0
	for _, report := range reports {
		if isSafe(report, 1) {
			safeCount += 1
		}
	}

	return solution.Solve(safeCount), nil
}

func parse(content string) ([][]int, error) {
	reports := strings.Split(content, "\n")
	var lines = make([][]int, 0, len(reports))
	for _, report := range reports[:len(reports)-1] {
		line := make([]int, 0, len(report)*2/3)
		levels := strings.Split(report, " ")
		for _, level := range levels {
			if l, err := strconv.Atoi(level); err != nil {
				return nil, err
			} else {
				line = append(line, l)
			}
		}
		lines = append(lines, line)
	}

	return lines, nil
}

func isSafe(report []int, dampenerCapacity int) bool {
	for _, rep := range permutations(report, dampenerCapacity) {
		differences := asDifferences(rep)
		if valid(differences) {
			return true
		}
	}

	return false
}

func valid(report []int) bool {
	inc := report[0] > 0
	for _, diff := range report {
		if diff == 0 {
			return false
		}
		if int(math.Abs(float64(diff))) > 3 {
			return false
		}

		if diff > 0 != inc {
			return false
		}
	}

	return true
}

func asDifferences(report []int) []int {
	differences := make([]int, len(report)-1)
	for index, level := range report[:len(report)-1] {
		differences[index] = report[index+1] - level
	}

	return differences
}

func permutations(report []int, c int) [][]int {
	perms := make([][]int, 0)
	if c == 0 {
		perms = append(perms, report)
		return perms
	}

	for index := range report {
		n := remove(report, index)
		perms = append(perms, n)
	}

	return perms
}

func remove(report []int, index int) []int {
	tmp := make([]int, len(report))
	copy(tmp, report)
	return append(tmp[:index], tmp[index+1:]...)
}
