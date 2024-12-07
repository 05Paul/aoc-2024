package day07

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
	equations := parse(content)
	total := solve(equations, []operation{ADD, MUL})

	return solution.New(total), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	equations := parse(content)
	total := solve(equations, []operation{ADD, MUL, CMB})

	return solution.New(total), nil
}

func solve(equations []equation, operations []operation) int {
	total := 0
	for _, equation := range equations {
		combinations := permutate(operations, len(equation.parts)-1)
		combination, last := combinations.next()
		for !last {
			if equation.solves(combination) {
				total += equation.result
				break
			}
			combination, last = combinations.next()
		}
	}

	return total
}

func parse(content string) []equation {
	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]
	equations := make([]equation, len(lines))

	for _, line := range lines {
		line = strings.ReplaceAll(line, ":", "")
		parts := strings.Split(line, " ")
		equations = append(equations, parseEquation(parts))
	}

	return equations
}

type operation = int

const (
	ADD operation = 0
	MUL operation = 1
	CMB operation = 2
)

type equation struct {
	result int
	parts  []int
}

func (e *equation) solves(operations []operation) bool {
	total := e.parts[0]
	for index, number := range e.parts[1:] {
		switch operations[index] {
		case ADD:
			total += number
		case MUL:
			total *= number
		case CMB:
			length := len(strconv.Itoa(number))
			total *= int(math.Pow10(length))
			total += number
		default:
			continue
		}
	}
	return total == e.result
}

func equate(result int, parts []int) equation {
	return equation{
		result: result,
		parts:  parts,
	}
}

func parseEquation(equation []string) equation {
	parts := make([]int, 0, len(equation))
	for _, part := range equation {
		if number, err := strconv.Atoi(part); err == nil {
			parts = append(parts, number)
		}
	}

	return equate(parts[0], parts[1:])
}

type permutation struct {
	possibilities      []operation
	base               int
	length             int
	currentPermutation int
	totalPermutations  int
}

func (p *permutation) next() (permutation []operation, last bool) {
	if p.totalPermutations == p.currentPermutation {
		return nil, true
	}

	iteration := p.currentPermutation
	permutation = make([]operation, p.length)
	for innerIndex := range permutation {
		permutation[innerIndex] = p.possibilities[iteration%p.base]
		iteration /= p.base
	}

	p.currentPermutation += 1

	return permutation, false
}

func permutate(possibilities []operation, length int) permutation {
	return permutation{
		possibilities:      possibilities,
		base:               len(possibilities),
		length:             length,
		currentPermutation: 0,
		totalPermutations:  int(math.Pow(float64(len(possibilities)), float64(length))),
	}
}
