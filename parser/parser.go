package parser

import (
	"strconv"
)

func New[In any, Out any](combinator func(captures []In) Operation[Out], captures ...SubCapture) Capture[Out] {
	return &Parser[In, Out]{
		sub:          captures,
		currentIndex: 0,
		outputs:      make([]In, 0),
		combinator:   combinator,
	}
}

func Multiply(operands ...int) Operation[int] {
	return &multiplication[int]{
		operands: operands,
		onEmpty:  nil,
		fold: func(accumulator int, operand int) int {
			return accumulator * operand
		},
	}
}

func NotEmpty(operands ...any) Operation[bool] {
	return &validate[any]{
		values: operands,
		check: func(values []any) bool {
			return len(values) > 0
		},
	}
}

func CaptureString(value string) SubCapture {
	return &stringCapture{
		value:        value,
		runes:        []rune(value),
		currentIndex: 0,
	}
}

func CaptureBetween(from string, to string) SubCapture {
	return &betweenCapture{
		from: &stringCapture{
			value:        from,
			runes:        []rune(from),
			currentIndex: 0,
		},
		to: &stringCapture{
			value:        to,
			runes:        []rune(to),
			currentIndex: 0,
		},
		fromDone: false,
	}
}

func CaptureInt(minDigits int, maxDigits int) SubCapture {
	return &intCapture{
		minDigits:  minDigits,
		maxDigits:  maxDigits,
		digitCount: 0,
		value:      0,
	}
}

type Capture[N any] interface {
	Parse(character rune) (content Operation[N], complete bool)
	Reset()
}

type SubCapture interface {
	SubParse(character rune) (content any, complete bool, reset bool, captured bool)
	Reset()
}

type Operation[T any] interface {
	Apply() *T
}

type Parser[In any, Out any] struct {
	sub          []SubCapture
	currentIndex int
	outputs      []In
	combinator   func(captures []In) Operation[Out]
}

func (s *Parser[In, Out]) Parse(character rune) (content Operation[Out], complete bool) {
	subOut, complete, reset, captured := s.sub[s.currentIndex].SubParse(character)

	if reset {
		s.Reset()

		if !captured {
			return s.Parse(character)
		}

		return nil, false
	}

	if !complete {
		return nil, false
	}

	if subOutValue, ok := subOut.(In); ok {
		s.outputs = append(s.outputs, subOutValue)
	}

	s.currentIndex += 1

	if s.currentIndex != len(s.sub) {
		s.sub[s.currentIndex].Reset()
		if !captured {
			return s.Parse(character)
		}
		return nil, false
	}

	out := s.combinator(s.outputs)
	s.Reset()

	return out, true
}

func (s *Parser[In, Out]) Reset() {
	s.currentIndex = 0
	s.outputs = make([]In, 0)
	for _, su := range s.sub {
		su.Reset()
	}
}

type stringCapture struct {
	value        string
	runes        []rune
	currentIndex int
}

func (s *stringCapture) SubParse(character rune) (content any, complete bool, reset bool, captured bool) {
	if s.currentIndex < len(s.runes) && s.runes[s.currentIndex] == character {
		s.currentIndex += 1
	} else {
		return nil, false, true, true
	}

	complete = s.currentIndex == len(s.runes)

	return s.value, complete, false, true
}

func (s *stringCapture) Reset() {
	s.currentIndex = 0
}

type betweenCapture struct {
	from     *stringCapture
	to       *stringCapture
	fromDone bool
}

func (b *betweenCapture) SubParse(character rune) (content any, complete bool, reset bool, captured bool) {
	if b.fromDone {
		complete, reset = b.capture(character, b.to)
		if reset {
			b.to.Reset()
		}

		if complete {
			return nil, true, false, true
		}
	} else {
		complete, reset = b.capture(character, b.from)
		if reset {
			b.Reset()
			return nil, true, false, false
		}

		if complete {
			b.fromDone = true
		}
	}

	return nil, false, false, true
}

func (b *betweenCapture) capture(character rune, capture *stringCapture) (complete bool, reset bool) {
	_, complete, reset, _ = capture.SubParse(character)
	return complete, reset
}

func (b *betweenCapture) Reset() {
	b.from.Reset()
	b.to.Reset()
	b.fromDone = false
}

type intCapture struct {
	minDigits  int
	maxDigits  int
	digitCount int
	value      int
}

func (i *intCapture) SubParse(character rune) (content any, complete bool, reset bool, captured bool) {
	value, err := strconv.Atoi(string(character))
	if err != nil {
		if i.digitCount >= i.minDigits && i.digitCount <= i.maxDigits {
			return i.value, true, false, false
		}
		return nil, false, true, false
	}

	i.digitCount += 1
	if i.digitCount > i.maxDigits {
		return nil, false, true, true
	}

	i.value *= 10
	i.value += value

	return nil, false, false, true
}

func (i *intCapture) Reset() {
	i.value = 0
	i.digitCount = 0
}

type multiplication[T any] struct {
	operands []T
	onEmpty  *T
	fold     func(accumulator T, operand T) T
}

func (m *multiplication[T]) Apply() *T {
	if len(m.operands) < 0 {
		return m.onEmpty
	}

	accumulator := m.operands[0]
	for _, operand := range m.operands[1:] {
		accumulator = m.fold(accumulator, operand)
	}

	return &accumulator
}

type validate[T any] struct {
	values []T
	check  func(values []T) bool
}

func (v *validate[T]) Apply() *bool {
	valid := v.check(v.values)
	return &valid
}
