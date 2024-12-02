package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day02 struct{}

func (d *Day02) SolvePart1(content string) (fmt.Stringer, error) {
	reports, err := d.parse(content)
	if err != nil {
		return nil, err
	}

	safeCount := 0
	for _, report := range reports {
		if d.isSafe(report, 0) {
			safeCount += 1
		}
	}

	return Solve(safeCount), nil
}

func (d *Day02) SolvePart2(content string) (fmt.Stringer, error) {
	reports, err := d.parse(content)
	if err != nil {
		return nil, err
	}

	safeCount := 0
	for _, report := range reports {
		if d.isSafe(report, 1) {
			safeCount += 1
		}
	}

	return Solve(safeCount), nil
}

func (d *Day02) parse(content string) ([][]int, error) {
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

func (d *Day02) isSafe(report []int, dampenerCapacity int) bool {
	for _, rep := range d.permutations(report, dampenerCapacity) {
		differences := d.asDifferences(rep)
		if d.valid(differences) {
			return true
		}
	}

	return false
}

func (d *Day02) valid(report []int) bool {
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

func (d *Day02) asDifferences(report []int) []int {
	differences := make([]int, len(report)-1)
	for index, level := range report[:len(report)-1] {
		differences[index] = report[index+1] - level
	}

	return differences
}

func (d *Day02) permutations(report []int, c int) [][]int {
	perms := make([][]int, 0)
	if c == 0 {
		perms = append(perms, report)
		return perms
	}

	for index := range report {
		n := d.remove(report, index)
		perms = append(perms, n)
	}

	return perms
}

func (d *Day02) remove(report []int, index int) []int {
	tmp := make([]int, len(report))
	copy(tmp, report)
	return append(tmp[:index], tmp[index+1:]...)
}
