package day09

import (
	"aoc/solution"
	"fmt"
	"strconv"
)

func New() solution.Solver {
	return &day{}
}

type day struct{}

func (d *day) SolvePart1(content string) (fmt.Stringer, error) {
	blocks, free := parse(content)
	defragmented := defragment(blocks, free)
	return solution.New(checksum(defragmented)), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func printDisk(blocks []int) {
	for _, value := range blocks {
		if value == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(value)
		}
	}
	fmt.Println()
}

func parse(content string) ([]int, []int) {
	blocks := make([]int, 0)
	free := make([]int, 0)
	fileId := 0
	for index, character := range content {
		value, err := strconv.Atoi(string(character))
		if err != nil {
			continue
		}

		if index%2 != 0 {
			free = extend(free, rangeFrom(len(blocks), value))
			blocks = extend(blocks, repeat(-1, value))
		} else {
			blocks = extend(blocks, repeat(fileId, value))
			fileId += 1
		}
	}

	return blocks, free
}

func defragment(blocks []int, free []int) []int {
	defragmented := make([]int, len(blocks))
	copy(defragmented, blocks)
	freeIndex := 0
	for index := range blocks {
		index = len(blocks) - index - 1
		if freeIndex >= len(free) || index <= free[freeIndex] {
			break
		}

		id := blocks[index]

		if id == -1 {
			continue
		}

		defragmented[free[freeIndex]] = id
		defragmented[index] = -1
		freeIndex += 1
	}

	return defragmented
}

func checksum(blocks []int) int {
	sum := 0
	for index, id := range blocks {
		if id == -1 {
			break
		}
		sum += index * id
	}

	return sum
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
