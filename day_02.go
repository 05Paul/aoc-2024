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
		if d.isSafe(report) {
			safeCount += 1
		}
	}

	return Solve(safeCount), nil
}

func (d *Day02) SolvePart2(content string) (fmt.Stringer, error) {
	return nil, fmt.Errorf("Not yet implemented")
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

func (d *Day02) isSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	if len(report) < 3 {
		return int(math.Abs(float64(report[0]-report[1]))) <= 3
	}

	diff := report[0] - report[1]

	if diff == 0 {
		return false
	}

	increasing := diff < 0

	return d.isSafeInner(0, 1, increasing, report)
}

func (d *Day02) isSafeInner(firstIndex int, secondIndex int, increasing bool, values []int) bool {
	if secondIndex >= len(values) {
		return true
	}

	diff := values[firstIndex] - values[secondIndex]

	if increasing && (diff >= 0 || diff < -3) {
		return false
	}

	if !increasing && (diff <= 0 || diff > 3) {
		return false
	}

	return d.isSafeInner(firstIndex+1, secondIndex+1, increasing, values)
}
