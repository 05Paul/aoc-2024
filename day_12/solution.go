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
	m := parse(content)
	total := 0
	for kind := range m.kinds {
		for index, region := range m.kindRegions(kind) {
			dbg.Printf(2, "%2d: Kind: %c, Area: %v, Sides: %v\n", index, kind, region.area, region.sides)
			total += region.area * region.sides
		}
	}

	return solution.New(total), nil
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
	sides     int
}

func (r *regionMap) kindRegions(kind rune) []region {
	regions := make([]region, 0)

	kinds := r.kinds[kind]
	kindsCopy := make([]position, len(kinds))
	copy(kindsCopy, kinds)
	for len(kindsCopy) > 0 {
		positions := r.positionsFrom(kind, kindsCopy[0], make(map[position]bool))
		for _, pos := range positions {
			if index := slices.Index(kindsCopy, pos); index != -1 {
				kindsCopy = append(kindsCopy[:index], kindsCopy[index+1:]...)
			}
		}

		regions = append(regions, r.regionFrom(kind, positions))
	}

	return regions
}

type fence struct {
	pos  position
	side int
}

type shapeSide struct {
	side int
	from int
	to   int
}

func (r *regionMap) regionFrom(kind rune, positions []position) region {
	perimiter := 0

	fences := make(map[fence]bool)
	for _, pos := range positions {
		current := 4
		for otherKind, fence := range r.neighbors(pos, false) {
			if !r.outOfBounds(fence.pos) && otherKind == kind {
				current -= 1
			} else {
				fence.pos = pos
				fences[fence] = true
			}
		}

		perimiter += current
	}

	dbg.Println(3, fences)
	sides := make([]shapeSide, 0)
	fenceSlice := keys(fences)
	dbg.Println(4, fences)
	for len(fences) > 0 {
		dbg.Printf(1, "Length before: %v\n", len(fences))
		s, removed := r.sideFrom(fenceSlice[0], fences)
		dbg.Printf(1, "Side: %v\nRemvoved: %v\n", s, removed)
		sides = append(sides, s)
		for _, r := range removed {
			delete(fences, r)
		}
		dbg.Printf(1, "Lenth after: %v\n", len(fences))
		fenceSlice = keys(fences)
	}

	dbg.Printf(4, "%c: %v\n", kind, len(sides))
	dbg.Printf(4, "%v\n", sides)

	return region{
		area:      len(positions),
		perimeter: perimiter,
		sides:     len(sides),
	}
}

func (r *regionMap) sideFrom(f fence, fences map[fence]bool) (shapeSide, []fence) {
	horizontal := f.side%2 != 0
	sideCoordinate := func(pos position) int {
		if horizontal {
			return pos.x
		} else {
			return pos.y
		}
	}

	sideMove := func(pos position, a bool) position {
		if horizontal && a {
			return pos.left()
		} else if horizontal {
			return pos.right()
		} else if a {
			return pos.up()
		} else {
			return pos.down()
		}
	}

	remove := []fence{f}
	side := shapeSide{
		side: f.side,
		from: sideCoordinate(f.pos),
		to:   sideCoordinate(f.pos),
	}

	a := f.pos
	for {
		temp := sideMove(a, true)
		if r.outOfBounds(temp) {
			break
		}

		tempFence := fence{
			pos:  temp,
			side: f.side,
		}
		if _, exists := fences[tempFence]; !exists {
			break
		}

		remove = append(remove, tempFence)
		a = temp
	}

	b := f.pos
	for {
		temp := sideMove(b, false)
		if r.outOfBounds(temp) {
			break
		}

		tempFence := fence{
			pos:  temp,
			side: f.side,
		}
		if _, exists := fences[tempFence]; !exists {
			break
		}

		remove = append(remove, tempFence)
		b = temp
	}

	return side, remove
}

func (r *regionMap) positionsFrom(kind rune, pos position, visited map[position]bool) []position {
	visited[pos] = true
	for otherKind, fence := range r.neighbors(pos, true) {
		if otherKind != kind {
			continue
		}

		position := fence.pos
		if _, exists := visited[position]; !exists {
			visited[position] = true
			for _, p := range r.positionsFrom(kind, position, visited) {
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

func (r *regionMap) outOfBounds(pos position) bool {
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

func (r *regionMap) neighbors(pos position, checkBounds bool) func(func(rune, fence) bool) {
	yieldDir := func(pos position, direction int, yield func(rune, fence) bool) bool {
		oob := r.outOfBounds(pos)
		if checkBounds && oob {
			return true
		}

		var kind rune
		if oob {
			kind = '-'
		} else {
			kind = r.regions[pos.y][pos.x]
		}

		return yield(kind, fence{
			pos:  pos,
			side: direction,
		})
	}

	return func(yield func(rune, fence) bool) {
		if !yieldDir(pos.left(), 0, yield) {
			return
		}

		if !yieldDir(pos.right(), 2, yield) {
			return
		}

		if !yieldDir(pos.up(), 1, yield) {
			return
		}

		if !yieldDir(pos.down(), 3, yield) {
			return
		}

	}
}

func keys[K comparable, V any](input map[K]V) []K {
	output := make([]K, 0, len(input))

	for key := range input {
		output = append(output, key)
	}

	return output
}
