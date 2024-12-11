package day11

import (
	dbg "aoc/debug"
	"aoc/solution"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func New() solution.Solver {
	return &day{}
}

type day struct{}

func (d *day) SolvePart1(content string) (fmt.Stringer, error) {
	stones := parse(content)
	total := 0
	for _, stone := range stones {
		cur := change(stone, 0, 25)
		dbg.Printf("\n%v", strings.Repeat("-", 10))
		total += cur
	}
	dbg.Println()

	return solution.New(total), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func parse(content string) []int {
	nums := strings.Split(content, " ")
	numbers := make([]int, len(nums))

	for index, num := range nums {
		num = strings.ReplaceAll(num, "\n", "")
		if number, err := strconv.Atoi(num); err == nil {
			numbers[index] = number
		}
	}

	return numbers
}

func change(number int, level int, maxLevel int) int {
	dbg.Printf("\n%v%v -> ", strings.Repeat("|", level), number)
	if level == maxLevel {
		dbg.Print("Done")
		return 1
	}

	if number == 0 {
		dbg.Print("1")
		return change(1, level+1, maxLevel)
	}

	length := len(strconv.Itoa(number))
	if length%2 == 0 {
		split := int(math.Pow10(length / 2))
		dbg.Printf("Split(%v)", split)
		return change(number/split, level+1, maxLevel) + change(number%split, level+1, maxLevel)
	}

	dbg.Print("* 2024")
	return change(number*2024, level+1, maxLevel)
}
