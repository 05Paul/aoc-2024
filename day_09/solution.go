package day09

import (
	"aoc/solution"
	"fmt"
	"slices"
	"strconv"
)

func New() solution.Solver {
	return &day{}
}

type day struct{}

func (d *day) SolvePart1(content string) (fmt.Stringer, error) {
	disk := parse(content)
	compacted := disk.compact(true)
	return solution.New(compacted.checksum()), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	disk := parse(content)
	compacted := disk.compact2()
	return solution.New(compacted.checksum()), nil
}

func (d *disk) print() {
	for _, block := range d.blocks {
		for range block.length {
			if block.id == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(block.id)
			}
		}
	}
	fmt.Println()
}

func (d *disk) print2() {
	for _, block := range d.blocks {
		fmt.Printf("%v: %v\n", block.id, block.length)
	}
}

func parse(content string) disk {
	blocks := make([]block, 0)
	last := 0
	id := 0
	for index, character := range content {
		value, err := strconv.Atoi(string(character))
		if err != nil {
			continue
		}

		if index%2 == 0 {
			blocks = append(blocks, block{
				id:     id,
				from:   last,
				length: value,
			})
			id += 1
		} else {
			blocks = append(blocks, block{
				id:     -1,
				from:   last,
				length: value,
			})
		}

		last += value
	}

	return disk{
		blocks: blocks,
	}
}

func repeat(value int, count int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for range count {
			if !yield(value) {
				return
			}
		}
	}
}

func rangeFrom(from int, length int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for offset := range length {
			if !yield(from + offset) {
				return
			}
		}
	}
}

func extend(values []int, iter func(func(int) bool)) []int {
	for value := range iter {
		values = append(values, value)
	}

	return values
}

type block struct {
	id     int
	from   int
	length int
}

type disk struct {
	blocks []block
}

func (d *disk) copied() *disk {
	copied := make([]block, len(d.blocks))
	copy(copied, d.blocks)

	return &disk{
		blocks: copied,
	}
}

func (d *disk) compact(allowFragmentation bool) *disk {
	compacted := d.copied()
	freeIndex := 1
	blockIndex := len(compacted.blocks) - 1
	id := compacted.blocks[blockIndex].id
	// compacted.print()
	for {
		if id < 0 {
			break
		}

		//	fmt.Println("------")
		//  fmt.Printf("  ID: %v, Current ID: %v\n", id, compacted.blocks[blockIndex].id)
		//	fmt.Printf("  BlockIndex: %v\n", blockIndex)
		//	fmt.Printf("  FreeIndex: %v, ID: %v\n", freeIndex, compacted.blocks[freeIndex].id)
		if compacted.blocks[freeIndex].id != -1 {
			freeIndex += 1
			continue
		}

		if compacted.blocks[blockIndex].id == -1 {
			blockIndex -= 1
			continue
		}

		if id != compacted.blocks[blockIndex].id {
			blockIndex -= 1
			continue
		}

		if freeIndex >= blockIndex {
			id -= 1
			freeIndex = 0
			continue
		}

		oldLen := len(compacted.blocks)
		completlyMoved, emptyRemaining, err := compacted.move(blockIndex, freeIndex, allowFragmentation)
		if err != nil {
			freeIndex += 1
			continue
		}

		newLen := len(compacted.blocks)
		diff := oldLen - newLen
		blockIndex -= diff

		if completlyMoved {
			id -= 1
		}

		_ = emptyRemaining
	}

	compacted.join()
	return compacted
}

func (d *disk) compact2() *disk {
	compacted := d.copied()
	block := len(compacted.blocks) - 1
	highestId := compacted.blocks[block].id

	for sub := range highestId + 1 {
		compacted.compactBlock(highestId - sub)
	}

	return compacted
}

func (d *disk) compactBlock(id int) bool {
	index := len(d.blocks) - 1
	for {
		if index < 0 {
			return false
		}

		if d.blocks[index].id == id {
			break
		}

		index -= 1
	}

	from := d.blocks[index]
	for to, blk := range d.blocks {
		if to >= index {
			break
		}

		if blk.id != -1 {
			continue
		}

		if blk.length == from.length {
			d.blocks[to].id = from.id
			d.blocks[index].id = -1
			return true
		}

		if blk.length > from.length {
			d.blocks[index].id = -1
			d.blocks[to].id = from.id
			d.blocks[to].length = from.length
			d.blocks = slices.Insert(d.blocks, to+1, block{
				id:     -1,
				from:   blk.from + from.length,
				length: blk.length - from.length,
			})
			return true
		}

	}

	return false
}

func (d *disk) join() {
	joined := make([]block, 0)
	current := d.blocks[0]
	for _, block := range d.blocks[1:] {
		if block.id != current.id {
			joined = append(joined, current)
			current = block
			continue
		}

		current.length += block.length
	}

	joined = append(joined, current)

	d.blocks = joined
}

func (d *disk) move(fromIndex int, toIndex int, allowFragmentation bool) (completlyMoved bool, emptyRemaining bool, err error) {
	to := d.blocks[toIndex]
	from := d.blocks[fromIndex]
	if to.id != -1 {
		return false, true, fmt.Errorf("Not empty")
	}

	if from.length == to.length {
		d.blocks[fromIndex].id = -1
		d.blocks[toIndex].id = from.id
		return true, false, nil
	}

	if from.length > to.length && !allowFragmentation {
		return false, true, fmt.Errorf("Cannot fragment")
	}

	if from.length > to.length {
		d.blocks[toIndex].id = from.id
		d.blocks[fromIndex].length = from.length - to.length
		d.blocks = slices.Insert(d.blocks, fromIndex+1, block{
			id:     -1,
			from:   from.from + to.length,
			length: to.length,
		})
		return false, false, nil
	}

	d.blocks[fromIndex].id = -1
	d.blocks = slices.Insert(d.blocks, toIndex, block{
		id:     from.id,
		from:   to.from,
		length: from.length,
	})

	d.blocks[toIndex+1].from = to.from + from.length
	d.blocks[toIndex+1].length = to.length - from.length

	return true, true, nil
}

func (d *disk) checksum() int {
	sum := 0
	index := 0
	for _, block := range d.blocks {
		if block.id == -1 {
			index += block.length
			continue
		}

		for range block.length {
			sum += index * block.id
			index += 1
		}
	}

	return sum
}
