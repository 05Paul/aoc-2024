package day12

import (
	dbg "aoc/debug"
	"aoc/solution"
	"fmt"
	"slices"
	"strings"
)

func New() solution.Solver {
	return &day{}
}

type day struct{}

func (d *day) SolvePart1(content string) (fmt.Stringer, error) {
	m := parse(content)
	total := 0
	for kind := range m.kinds {
		for index, region := range m.kindRegions(kind) {
			dbg.Printf(2, "%2d: Kind: %c, Area: %v, Perimeter: %v\n", index, kind, region.area, region.perimeter)
			total += region.area * region.perimeter
		}
	}

	return solution.New(total), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {

	return nil, fmt.Errorf("")
}

func parse(content string) regionMap {
	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]

	grid := make([][]rune, len(lines))
	kinds := make(map[rune][]position)
	for y, line := range lines {
		grid[y] = make([]rune, len(line))
		for x, character := range line {
			grid[y][x] = character
			kinds[character] = append(kinds[character], position{
				x: x,
				y: y,
			})
		}
	}

	return regionMap{
		regions: grid,
		kinds:   kinds,
	}
}

type position struct {
	x int
	y int
}

func (p *position) up() position {
	return position{
		x: p.x,
		y: p.y - 1,
	}
}

func (p *position) down() position {
	return position{
		x: p.x,
		y: p.y + 1,
	}
}

func (p *position) left() position {
	return position{
		x: p.x - 1,
		y: p.y,
	}
}

func (p *position) right() position {
	return position{
		x: p.x + 1,
		y: p.y,
	}
}

type regionMap struct {
	regions [][]rune
	kinds   map[rune][]position
}

type region struct {
	area      int
	perimeter int
}

func (r *regionMap) kindRegions(kind rune) []region {
	regions := make([]region, 0)

	kinds := r.kinds[kind]
	kindsCopy := make([]position, len(kinds))
	copy(kindsCopy, kinds)
	for len(kindsCopy) > 0 {
		positions := r.regionFrom(kind, kindsCopy[0], make(map[position]bool))
		for _, pos := range positions {
			if index := slices.Index(kindsCopy, pos); index != -1 {
				kindsCopy = append(kindsCopy[:index], kindsCopy[index+1:]...)
			}
		}

		regions = append(regions, region{
			area:      len(positions),
			perimeter: r.perimeter(kind, positions),
		})
	}

	return regions
}

func (r *regionMap) perimeter(kind rune, positions []position) int {
	total := 0
	for _, pos := range positions {
		current := 4
		for otherKind := range r.neighbors(pos) {
			if otherKind == kind {
				current -= 1
			}
		}
		total += current
	}

	return total
}

func (r *regionMap) regionFrom(kind rune, pos position, visited map[position]bool) []position {
	visited[pos] = true
	for otherKind, position := range r.neighbors(pos) {
		if otherKind != kind {
			continue
		}

		if _, exists := visited[position]; !exists {
			visited[position] = true
			for _, p := range r.regionFrom(kind, position, visited) {
				visited[p] = true
			}
		}
	}

	positions := make([]position, 0, len(visited))
	for p := range visited {
		positions = append(positions, p)
	}

	return positions
}

func (r *regionMap) neighbors(pos position) func(func(rune, position) bool) {
	outOfBounds := func(pos position) bool {
		if pos.x < 0 {
			return true
		}

		if pos.y < 0 {
			return true
		}

		if pos.x >= len(r.regions[0]) {
			return true
		}

		if pos.y >= len(r.regions) {
			return true
		}

		return false
	}

	yieldDir := func(pos position, yield func(rune, position) bool) bool {
		if outOfBounds(pos) {
			return true
		}

		return yield(r.regions[pos.y][pos.x], pos)
	}

	return func(yield func(rune, position) bool) {
		if !yieldDir(pos.left(), yield) {
			return
		}

		if !yieldDir(pos.right(), yield) {
			return
		}

		if !yieldDir(pos.up(), yield) {
			return
		}

		if !yieldDir(pos.down(), yield) {
			return
		}

	}
}
