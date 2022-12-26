package puzzles

import (
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(6, Puzzle006{})
}

type Puzzle006 struct{}

func (p Puzzle006) Solve() puzzle.Answer {
	s := sumSquares(1, 100)
	ss := sumRange(1, 100)
	return puzzle.FromValue(ss*ss - s)
}

func sumRange(start, end int) int {
	n := 0
	for i := start; i <= end; i++ {
		n += i
	}
	return n
}

func sumSquares(start, end int) int {
	n := 0
	for i := start; i <= end; i++ {
		n += i * i
	}
	return n
}
