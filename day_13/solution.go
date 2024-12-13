package day13

import (
	dbg "aoc/debug"
	"aoc/solution"
	"fmt"
	"strconv"
	"strings"
)

func New() solution.Solver {
	return &day{}
}

type day struct{}

const (
	buttonAKey string = "Button A"
	buttonBKey string = "Button B"
	prizeKey   string = "Prize"
)

func (d *day) SolvePart1(content string) (fmt.Stringer, error) {
	machines := parse(content)
	total := 0
	for index, machine := range machines {
		dbg.Printf(2, "%2d: %v\n", index, machine)
		solved, a, b := machine.solution()

		if max(a, b) > 100 {
			dbg.Println(3, "Too many presses")
			continue
		}

		if solved {
			total += a*3 + b*1
			dbg.Printf(3, "Solution #%02d: a=%v, b=%v\n", index, a, b)
		} else {
			dbg.Printf(3, "#%02d: unsolvable\n", index)
		}
	}
	return solution.New(total), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	offset := 10000000000000
	machines := parse(content)
	total := 0
	for index, machine := range machines {
		dbg.Printf(2, "%2d: %v\n", index, machine)
		machine.prize.x += offset
		machine.prize.y += offset
		solved, a, b := machine.solution()

		if solved {
			total += a*3 + b*1
			dbg.Printf(3, "Solution #%02d: a=%v, b=%v\n", index, a, b)
		} else {
			dbg.Printf(3, "#%02d: unsolvable\n", index)
		}
	}
	return solution.New(total), nil
}

func parse(content string) []clawMachine {
	blocks := strings.Split(content, "\n\n")
	machines := make([]clawMachine, len(blocks))
	for index, block := range blocks {
		block = strings.TrimRight(block, "\n")
		parsed := parseBlock(block)
		machines[index] = clawMachine{
			buttonA: parsed[buttonAKey],
			buttonB: parsed[buttonBKey],
			prize:   parsed[prizeKey],
		}
	}

	return machines
}

func parseBlock(block string) map[string]position {
	lines := strings.Split(block, "\n")
	positions := make(map[string]position, len(lines))
	for _, line := range lines {
		keyValue := strings.Split(line, ": ")
		key, value := keyValue[0], keyValue[1]
		x, y := parseCoordinates(value)
		positions[key] = position{
			x: x,
			y: y,
		}
	}

	return positions
}

func parseCoordinates(value string) (x, y int) {
	coordinates := strings.Split(value, ", ")
	return parseCoordinate(coordinates[0], "X", "Y", "="), parseCoordinate(coordinates[1], "X", "Y", "=")
}

func parseCoordinate(value string, remove ...string) int {
	for _, r := range remove {
		value = strings.ReplaceAll(value, r, "")
	}

	coordinate, _ := strconv.Atoi(value)

	return coordinate
}

type clawMachine struct {
	buttonA position
	buttonB position
	prize   position
}

func (c *clawMachine) solution() (solvable bool, a int, b int) {
	num := c.prize.y*c.buttonA.x - c.prize.x*c.buttonA.y
	den := c.buttonA.x*c.buttonB.y - c.buttonA.y*c.buttonB.x

	if num%den != 0 {
		return false, -1, -1
	}

	b = num / den
	a = (c.prize.x - c.buttonB.x*b) / c.buttonA.x
	return true, a, b
}

type position struct {
	x int
	y int
}
