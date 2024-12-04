package solution

import "fmt"

type Solver interface {
	SolvePart1(puzzle string) (fmt.Stringer, error)
	SolvePart2(puzzle string) (fmt.Stringer, error)
}

type Solution[T any] struct {
	value T
}

func Solve[T any](value T) *Solution[T] {
	return &Solution[T]{
		value: value,
	}
}

func (s *Solution[any]) String() string {
	return fmt.Sprintf("%v", s.value)
}
