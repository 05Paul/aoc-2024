package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		d     uint64
		path1 string
		path2 string
	)

	flag.Uint64Var(&d, "day", 0, "Day to solve")
	flag.StringVar(&path1, "f1", "", "Path to part 1")
	flag.StringVar(&path2, "f2", "", "Path to part 2")
	flag.Parse()

	var day = uint8(d)

	var (
		solver Solver
		err    error
	)
	if solver, err = getSolver(day); err != nil {
		fmt.Fprintf(os.Stderr, "Could not get solver: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Day %v:\n", day)

	if path1 == "" {
		path1 = fmt.Sprintf("inputs/day_%02d/part1.txt", day)
	}

	var input1 []byte
	if input1, err = os.ReadFile(path1); err != nil {
		fmt.Fprintf(os.Stderr, "Could not read input 1: %v\n", err)
		os.Exit(1)
	}

	if solution, err := solver.SolvePart1(string(input1)); err != nil {
		fmt.Fprintf(os.Stderr, "Could not solve part 1: %v\n", err)
	} else {
		fmt.Printf(" * Part 1: %v\n", solution)
	}

	if path2 == "" {
		path2 = fmt.Sprintf("inputs/day_%02d/part2.txt", day)
	}

	var input2 []byte

	if input2, err = os.ReadFile(path2); err != nil {
		fmt.Fprintf(os.Stderr, "Could not read input 2: %v\n", err)
		os.Exit(1)
	}

	if solution, err := solver.SolvePart2(string(input2)); err != nil {
		fmt.Fprintf(os.Stderr, "Could not solve part 2: %v\n", err)
	} else {
		fmt.Printf(" * Part 2: %v\n", solution)
	}
}

type Solver interface {
	SolvePart1(puzzle string) (fmt.Stringer, error)
	SolvePart2(puzzle string) (fmt.Stringer, error)
}

type Solution[T any] struct {
	value T
}

func Solve[T any](value T) *Solution[T] {
	return &Solution[T]{
		value: value,
	}
}

func (s *Solution[any]) String() string {
	return fmt.Sprintf("%v", s.value)
}

func getSolver(day uint8) (Solver, error) {
	switch day {
	case 1:
		return &Day01{}, nil
	default:
		return nil, fmt.Errorf("Undefined day: %v", day)
	}
}
