package main

import (
	day01 "aoc/day_01"
	day02 "aoc/day_02"
	day03 "aoc/day_03"
	day04 "aoc/day_04"
	day05 "aoc/day_05"
	day06 "aoc/day_06"
	day07 "aoc/day_07"
	day08 "aoc/day_08"
	day09 "aoc/day_09"
	day10 "aoc/day_10"
	day11 "aoc/day_11"
	day12 "aoc/day_12"
	day13 "aoc/day_13"
	day14 "aoc/day_14"
	dbg "aoc/debug"
	"aoc/solution"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		d          uint64
		path1      string
		path2      string
		part1      bool
		part2      bool
		debugLevel int
	)

	flag.Uint64Var(&d, "day", 0, "Day to solve")
	flag.StringVar(&path1, "f1", "", "Path to part 1")
	flag.StringVar(&path2, "f2", "", "Path to part 2")
	flag.BoolVar(&part1, "p1", false, "Run part 1")
	flag.BoolVar(&part2, "p2", false, "Run part 2")
	flag.IntVar(&debugLevel, "debug", 0, "Set debug level")
	flag.Parse()

	var day = uint8(d)
	dbg.SetLevel(debugLevel)

	var (
		solver solution.Solver
	)
	if solver = getSolver(day); solver == nil {
		fmt.Fprintf(os.Stderr, "Could not get solver: Day %v\n", day)
		os.Exit(1)
	}

	fmt.Printf("Day %v:\n", day)

	if !part1 && !part2 {
		part1 = true
		part2 = true
	}

	if part1 {
		part(day, solver, path1, one)
	}

	if part2 {
		part(day, solver, path2, two)
	}
}

type Part int

const (
	one Part = 1
	two Part = 2
)

func part(day uint8, solver solution.Solver, path string, part Part) {
	if path == "" {
		path = fmt.Sprintf("inputs/day_%02d/input.txt", day)
	}

	var input []byte
	var err error
	if input, err = os.ReadFile(path); err != nil {
		fmt.Fprintf(os.Stderr, "Could not read input for part %v: %v\n", part, err)
		os.Exit(1)
	}

	if solution, err := solvePart(solver, string(input), part); err != nil {
		fmt.Fprintf(os.Stderr, "Could not solve part %v: %v\n", part, err)
	} else {
		fmt.Printf(" * Part %v: %v\n", part, solution)
	}
}

func solvePart(solver solution.Solver, puzzle string, part Part) (fmt.Stringer, error) {
	if part == one {
		return solver.SolvePart1(puzzle)
	} else {
		return solver.SolvePart2(puzzle)
	}
}

func getSolver(day uint8) solution.Solver {
	switch day {
	case 1:
		return day01.New()
	case 2:
		return day02.New()
	case 3:
		return day03.New()
	case 4:
		return day04.New()
	case 5:
		return day05.New()
	case 6:
		return day06.New()
	case 7:
		return day07.New()
	case 8:
		return day08.New()
	case 9:
		return day09.New()
	case 10:
		return day10.New()
	case 11:
		return day11.New()
	case 12:
		return day12.New()
	case 13:
		return day13.New()
	case 14:
		return day14.New()
	default:
		return nil
	}
}
