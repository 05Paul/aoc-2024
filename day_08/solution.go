package day08

import (
	"aoc/solution"
	"fmt"
	"strings"
)

func New() solution.Solver {
	return &day{}
}

type day struct{}

func (d *day) SolvePart1(content string) (fmt.Stringer, error) {
	positions, grid := parse(content)
	iterator := iter(positions)
	antinodes := make(map[position]bool)
	for combination := range combinations(iterator) {
		for _, antinode := range combination.antinodes() {
			if grid.inbounds(antinode) {
				antinodes[antinode] = true
			}
		}
	}

	return solution.New(len(antinodes)), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func parse(content string) (map[rune][]position, grid) {
	antenna := make(map[rune][]position)

	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]

	height := len(lines) - 1
	width := 0

	for y, line := range lines {
		width = len(line) - 1
		for x, character := range line {
			if character == '.' {
				continue
			}

			var value []position
			var exists bool
			if value, exists = antenna[character]; !exists {
				value = []position{}
			}

			position := position{
				x: x,
				y: height - y,
			}
			value = append(value, position)
			antenna[character] = value
		}
	}

	return antenna, zeroed(width, height)
}

type grid struct {
	minX int
	maxX int
	minY int
	maxY int
}

func (g *grid) inbounds(pos position) bool {
	return pos.x >= g.minX && pos.x <= g.maxX && pos.y >= g.minY && pos.y <= g.maxY
}

func zeroed(maxX int, maxY int) grid {
	return grid{
		minX: 0,
		maxX: maxX,
		minY: 0,
		maxY: maxY,
	}
}

type position struct {
	x int
	y int
}

type iterate struct {
	value            map[rune][]position
	keys             []rune
	keyIndex         int
	valueIndex       int
	combinationIndex int
}

func iter(value map[rune][]position) iterate {
	return iterate{
		value:            value,
		keys:             keys(value),
		keyIndex:         0,
		valueIndex:       0,
		combinationIndex: 1,
	}
}

type combination struct {
	first  position
	second position
}

type alignment = int

const (
	horizontal      alignment = 0
	vertical        alignment = 1
	diagonalIncline alignment = 2
	diagonalDecline alignment = 3
)

func (c *combination) inline() (bool, alignment) {
	if c.first.y == c.second.y {
		return true, horizontal
	}

	if c.first.x == c.second.x {
		return true, vertical
	}

	xDiff := c.xDiff()
	yDiff := c.yDiff()
	inline := yDiff%xDiff == 0
	slope := yDiff / xDiff

	if inline && slope < 0 {
		if c.first.x > c.second.x {
			return true, diagonalIncline
		}

		return true, diagonalDecline
	}

	if inline && slope > 0 {
		if c.first.x > c.second.x {
			return true, diagonalDecline
		}

		return true, diagonalIncline
	}

	return false, -1
}

func (c *combination) antinodes() []position {
	xDist, yDist := c.distance()
	return []position{
		{
			x: c.first.x - xDist,
			y: c.first.y - yDist,
		},
		{
			x: c.second.x + xDist,
			y: c.second.y + yDist,
		},
	}
}

func (c *combination) distance() (int, int) {
	return c.xDiff(), c.yDiff()
}

func (c *combination) xDiff() int {
	return c.second.x - c.first.x
}

func (c *combination) yDiff() int {
	return c.second.y - c.first.y
}

func (i *iterate) next() *combination {
	if i.keyIndex >= len(i.keys) {
		return nil
	}

	if i.valueIndex >= len(i.value[i.keys[i.keyIndex]]) {
		i.keyIndex += 1
		i.valueIndex = 0
		i.combinationIndex = i.valueIndex + 1
		return i.next()
	}

	if i.combinationIndex >= len(i.value[i.keys[i.keyIndex]]) {
		i.valueIndex += 1
		i.combinationIndex = i.valueIndex + 1
		return i.next()
	}

	first := i.value[i.keys[i.keyIndex]][i.valueIndex]
	second := i.value[i.keys[i.keyIndex]][i.combinationIndex]

	i.combinationIndex += 1

	return &combination{
		first:  first,
		second: second,
	}
}

func combinations(iter iterate) func(func(combination) bool) {
	return func(yield func(combination) bool) {
		value := iter.next()
		for value != nil {
			if !yield(*value) {
				return
			}
			value = iter.next()
		}
	}
}

func keys[K comparable, V any](value map[K]V) []K {
	keys := make([]K, len(value))
	index := 0
	for key := range value {
		keys[index] = key
		index += 1
	}

	return keys
}
