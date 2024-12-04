package day03

import (
	"aoc/parser"
	"aoc/solution"
	"fmt"
)

func New() solution.Solver {
	return &day{}
}

type day struct{}

func (d *day) SolvePart1(content string) (fmt.Stringer, error) {
	parser := parser.New(func(captures []int) parser.Operation[int] {
		return parser.Multiply(captures...)
	},
		parser.CaptureString("mul"),
		parser.CaptureString("("),
		parser.CaptureInt(1, 3),
		parser.CaptureString(","),
		parser.CaptureInt(1, 3),
		parser.CaptureString(")"),
	)

	total := 0
	for _, character := range content {
		operation, complete := parser.Parse(character)
		if complete {
			result := *operation.Apply()
			total += result
		}
	}

	return solution.New(total), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	parser := parser.New(func(captures []int) parser.Operation[int] {
		return parser.Multiply(captures...)
	},
		parser.CaptureBetween("don't()", "do()"),
		parser.CaptureString("mul"),
		parser.CaptureString("("),
		parser.CaptureInt(1, 3),
		parser.CaptureString(","),
		parser.CaptureInt(1, 3),
		parser.CaptureString(")"),
	)

	total := 0
	for _, character := range content {
		operation, complete := parser.Parse(character)
		if complete {
			result := *operation.Apply()
			total += result
		}
	}

	return solution.New(total), nil
}
