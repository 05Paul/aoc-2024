package day01

import (
	"aoc/solution"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func New() solution.Solver {
	return &day{}
}

type day struct{}

func (d *day) SolvePart1(content string) (fmt.Stringer, error) {
	var (
		list1 []int
		list2 []int
		err   error
	)

	if list1, list2, err = parse(content); err != nil {
		return nil, err
	}

	slices.Sort(list1)
	slices.Sort(list2)

	var totalDistance = 0

	for index := range list1 {
		totalDistance += int(math.Abs(float64(list1[index] - list2[index])))
	}

	return solution.New(totalDistance), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	var (
		list1 []int
		list2 []int
		err   error
	)

	if list1, list2, err = parse(content); err != nil {
		return nil, err
	}

	var counts = numberCount(list2)
	var similiarity = 0

	for _, value := range list1 {
		similiarity += value * counts[value]
	}

	return solution.New(similiarity), nil
}

func parse(content string) ([]int, []int, error) {
	var lines = strings.Split(content, "\n")
	var (
		list1 []int = make([]int, len(lines))
		list2 []int = make([]int, len(lines))
	)

	for index, line := range lines[:len(lines)-1] {
		var split = strings.Split(line, "   ")

		if value, err := strconv.Atoi(split[0]); err != nil {
			return nil, nil, err
		} else {
			list1[index] = value
		}

		if value, err := strconv.Atoi(split[1]); err != nil {
			return nil, nil, err
		} else {
			list2[index] = value
		}
	}

	return list1, list2, nil
}

func numberCount(numbers []int) map[int]int {
	var counts = make(map[int]int)

	for _, value := range numbers {
		counts[value] = counts[value] + 1
	}

	return counts
}
