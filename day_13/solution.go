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
		solved, a, b := machine.solution(100)

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
	return nil, fmt.Errorf("Not yet implemented")
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

func (c *clawMachine) solution(maxPlays int) (solvable bool, a int, b int) {
	maxDiv := min(max(c.maxXDiv(), c.maxYDiv()), maxPlays)
	for diff := range maxDiv + 1 {
		b = maxDiv - diff

		remX := c.prize.x - b*c.buttonB.x
		if remX < 0 && remX%c.buttonA.x != 0 {
			continue
		}

		remY := c.prize.y - b*c.buttonB.y
		if remY < 0 && remY%c.buttonA.y != 0 {
			continue
		}

		a = remX / c.buttonA.x
		if a > maxPlays {
			continue
		}

		solX := a*c.buttonA.x + b*c.buttonB.x
		solY := a*c.buttonA.y + b*c.buttonB.y
		if solX == c.prize.x && solY == c.prize.y {
			dbg.Printf(3, "Prize(x=%v, y=%v), Solution(x=%v, y=%v)\n", c.prize.x, c.prize.y, solX, solY)
			return true, a, b
		}
	}

	return false, -1, -1
}

func (c *clawMachine) maxXDiv() int {
	return c.prize.x / min(c.buttonA.x, c.buttonB.x)
}

func (c *clawMachine) maxYDiv() int {
	return c.prize.y / min(c.buttonA.y, c.buttonB.y)
}

type position struct {
	x int
	y int
}
