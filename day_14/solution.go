package day14

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

func (d *day) SolvePart1(content string) (fmt.Stringer, error) {
	robots := parse(content)
	g := grid{
		height: 103,
		width:  101,
	}

	quadrants := make(map[int]int)
	for _, robot := range robots {
		g.simulate(&robot, 100)
		quad := g.quadrant(&robot)
		quadrants[quad] += 1
		dbg.Println(3, quad)
	}

	delete(quadrants, 0)
	product := 1
	for quad, count := range quadrants {
		product *= count
		dbg.Printf(4, "Q%v: %v\n", quad, count)
	}
	dbg.Println(3, quadrants)
	return solution.New(product), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func parse(content string) []robot {
	lines := strings.Split(strings.TrimRight(content, "\n"), "\n")
	robots := make([]robot, len(lines))
	for index, line := range lines {
		parts := strings.Split(line, " ")
		robots[index] = robot{
			pos:      parsePosition(parts[0]),
			velocity: parsePosition(parts[1]),
		}
	}

	return robots
}

func parsePosition(value string) position {
	parts := strings.Split(value, "=")
	coordinates := strings.Split(parts[1], ",")
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])
	return position{
		x: x,
		y: y,
	}
}

type position struct {
	x int
	y int
}

type robot struct {
	pos      position
	velocity position
}

type grid struct {
	height int
	width  int
}

func (g *grid) simulate(r *robot, seconds int) {
	r.pos.x = addWrapping(r.pos.x, r.velocity.x*seconds, g.width)
	r.pos.y = addWrapping(r.pos.y, r.velocity.y*seconds, g.height)
}

func (g *grid) quadrant(r *robot) int {
	xMid := g.width / 2
	yMid := g.height / 2

	if r.pos.x < xMid && r.pos.y > yMid {
		return 2
	}

	if r.pos.x > xMid && r.pos.y > yMid {
		return 1
	}

	if r.pos.x > xMid && r.pos.y < yMid {
		return 4
	}

	if r.pos.x < xMid && r.pos.y < yMid {
		return 3
	}

	return 0
}

func addWrapping(a int, b int, limit int) int {
	sum := a + b
	rem := sum % limit
	if rem < 0 {
		return limit + rem
	}

	return rem
}
