package day14

import (
	dbg "aoc/debug"
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
	robots := parse(content)
	g := grid{
		height: 103,
		width:  101,
	}

	steps := -1
	for frame := range 10000 {
		for index, robot := range robots {
			g.simulate(&robot, 1)
			robots[index] = robot
		}

		qrx, qry := g.quartileRange(robots)
		if qrx < 20 && qry < 20 {
			steps = frame + 1
			dbg.Printf(3, "(%v): x: %v, y: %v\n", frame, qrx, qry)
			g.print(robots)
			break
		}
		//time.Sleep(time.Millisecond * 250)
	}

	return solution.New(steps), nil
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

func (g *grid) print(robots []robot) {
	positionedRobots := make([][]*robot, g.height)
	for row := range positionedRobots {
		positionedRobots[row] = make([]*robot, g.width)
	}

	for _, r := range robots {
		positionedRobots[r.pos.y][r.pos.x] = &r
	}

	dbg.Println(1, strings.Repeat("-", g.width+2))

	for _, line := range positionedRobots {
		dbg.Print(1, "|")
		for _, r := range line {
			if r == nil {
				dbg.Print(1, " ")
			} else {
				dbg.Print(1, "■")
			}
		}
		dbg.Println(1, "|")
	}
	dbg.Println(1, strings.Repeat("-", g.width+2))
}

func (g *grid) quartileRange(robots []robot) (x int, y int) {
	quartileR := func(values []int) int {
		length := len(values)
		return values[length/4*3] - values[length/4]
	}
	xValues := make([]int, len(robots))
	yValues := make([]int, len(robots))

	for index, r := range robots {
		xValues[index] = r.pos.x
		yValues[index] = r.pos.y
	}

	slices.Sort(xValues)
	slices.Sort(yValues)

	return quartileR(xValues), quartileR(yValues)
}

func addWrapping(a int, b int, limit int) int {
	sum := a + b
	rem := sum % limit
	if rem < 0 {
		return limit + rem
	}

	return rem
}
