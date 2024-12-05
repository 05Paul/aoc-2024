package day05

import (
	"aoc/solution"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func New() solution.Solver {
	return &day{}
}

type day struct{}

func (d *day) SolvePart1(content string) (fmt.Stringer, error) {
	sum := 0
	constraints, orderings := parse(content)

	for _, ordering := range orderings {
		valid := correctOrder(ordering, constraints)
		if valid {
			sum += ordering[len(ordering)/2]
		}
	}

	return solution.New(sum), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	sum := 0
	constraints, orderings := parse(content)

	for _, ordering := range orderings {
		valid := correctOrder(ordering, constraints)
		if valid {
			continue
		}
		newOrder := ordered(ordering, constraints)
		sum += newOrder[len(newOrder)/2]
	}

	return solution.New(sum), nil

}

func ordered(ordering []int, constraints map[int][]int) []int {
	newOrder := make([]int, 0)
Outer:
	for _, item := range ordering {
		currentContraints, ok := constraints[item]
		if !ok {
			newOrder = append(newOrder, item)
			continue
		}

		for index, placed := range newOrder {
			for _, constraint := range currentContraints {
				if placed == constraint {
					newOrder = slices.Insert(newOrder, index, item)
					continue Outer
				}
			}
		}

		newOrder = append(newOrder, item)
	}

	return newOrder
}

func correctOrder(ordering []int, constraints map[int][]int) bool {
	for index, item := range ordering[1:] {
		currentContraints, ok := constraints[item]
		if !ok {
			continue
		}

		for _, constraint := range currentContraints {
			for _, preceeding := range ordering[:index+1] {
				if constraint == preceeding {
					return false
				}
			}
		}
	}

	return true
}

func parse(content string) (map[int][]int, [][]int) {
	lines := strings.Split(content, "\n")
	constraint := true

	constraints := make(map[int][]int)
	orderings := make([][]int, 0)

	for _, line := range lines {
		if len(line) < 2 {
			constraint = false
			continue
		}

		if constraint {
			n := nums(strings.Split(line, "|"))
			con, ok := constraints[n[0]]
			if !ok {
				con = make([]int, 0)
			}
			con = append(con, n[1])
			constraints[n[0]] = con
		} else {
			orderings = append(orderings, nums(strings.Split(line, ",")))
		}
	}
	return constraints, orderings
}

func nums(strings []string) []int {
	nums := make([]int, 0)
	for _, s := range strings {
		num, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}

	return nums
}
