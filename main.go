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
	flag.StringVar(&path1, "f1", "inputs/day_%v/part1.txt", "Path to part 1")
	flag.StringVar(&path2, "f2", "inputs/day_%v/part2.txt", "Path to part 2")
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

	var input1 []byte

	if input1, err = os.ReadFile(path1); err != nil {
		fmt.Fprintln(os.Stderr, "Could not read input 1")
		os.Exit(1)
	}

	if solution, err := solver.SolvePart1(string(input1)); err != nil {
		fmt.Fprintf(os.Stderr, "Could not solve part 1: %v\n", err)
	} else {
		fmt.Printf("Solution for part 1: %v\n", solution)
	}

	var input2 []byte

	if input2, err = os.ReadFile(path2); err != nil {
		fmt.Fprintln(os.Stderr, "Could not read input 2")
		os.Exit(1)
	}

	if solution, err := solver.SolvePart2(string(input2)); err != nil {
		fmt.Fprintf(os.Stderr, "Could not solve part 2: %v\n", err)
	} else {
		fmt.Printf("Solution for part 2: %v\n", solution)
	}
}

type Solver interface {
	SolvePart1(puzzle string) (fmt.Stringer, error)
	SolvePart2(puzzle string) (fmt.Stringer, error)
}

func getSolver(day uint8) (Solver, error) {
	switch day {
	default:
		return nil, fmt.Errorf("Undefined day: %v", day)
	}
}
