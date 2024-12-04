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
	parser := parser.New(func(captures []any) parser.Operation[int] {
		op1, ok1 := captures[2].(*int)
		op2, ok2 := captures[4].(*int)

		if !ok1 || !ok2 {
			panic(fmt.Sprintf("Could not covert operand"))
		}

		return parser.Multiply(*op1, *op2)
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

	return solution.Solve(total), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	parser := parser.New(func(captures []any) parser.Operation[int] {
		op1, ok1 := captures[3].(*int)
		op2, ok2 := captures[5].(*int)

		if !ok1 || !ok2 {
			panic(fmt.Sprintf("Could not covert operand"))
		}

		return parser.Multiply(*op1, *op2)
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

	return solution.Solve(total), nil
}
