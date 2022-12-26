package puzzles

import (
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(7, Puzzle007{})
}

type Puzzle007 struct{}

func (p Puzzle007) Solve() puzzle.Answer {
	return puzzle.FromValue(p.getNthPrime(1, 6))
}

func (p Puzzle007) getNthPrime(start, n int) int {
	return 0
}
