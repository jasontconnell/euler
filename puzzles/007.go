package puzzles

import (
	"github.com/jasontconnell/euler/common"
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(7, Puzzle007{})
}

type Puzzle007 struct{}

func (p Puzzle007) Solve() puzzle.Answer {
	return puzzle.FromValue(p.getNthPrime(1, 10_001))
}

func (p Puzzle007) getNthPrime(start, n int) int {
	idx := 0
	cur := start
	for idx != n {
		cur++
		if common.IsPrime(cur) {
			idx++
		}
	}
	return cur
}
