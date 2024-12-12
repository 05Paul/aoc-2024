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
	total := stoneCount(stones, 25)

	return solution.New(total), nil
}

func (d *day) SolvePart2(content string) (fmt.Stringer, error) {
	stones := parse(content)
	total := stoneCount(stones, 75)

	return solution.New(total), nil
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

func stoneCount(stones []int, blinks int) int {
	cache := make(map[int][]int)
	stoneMap := make(map[int]int)
	for _, stone := range stones {
		stoneMap[stone] += 1
	}
	dbg.Printf(1, "Stones: %v\n", stoneMap)
	total := change(stoneMap, blinks, cache)
	return total
}

func change(stones map[int]int, maxLevel int, cache map[int][]int) int {
	cache[0] = []int{1}

	for range maxLevel {
		next := make(map[int]int)
		for stone, count := range stones {
			if into, exists := cache[stone]; exists {
				for _, i := range into {
					next[i] += count
				}
				continue
			}

			length := len(strconv.Itoa(stone))
			if length%2 == 0 {
				split := int(math.Pow10(length / 2))
				left := stone / split
				right := stone % split

				dbg.Printf(2, "Split(%v) -> %v, %v\n", stone, left, right)

				cache[stone] = []int{left, right}
				next[left] += count
				next[right] += count
				continue
			}

			product := stone * 2024
			cache[stone] = []int{product}
			next[product] += count
			dbg.Printf(2, "%v * 2024 -> %v\n", stone, product)
		}
		stones = next
	}

	total := 0
	for _, count := range stones {
		total += count
	}

	return total
}
