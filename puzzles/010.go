package puzzles

import (
	"github.com/jasontconnell/euler/common"
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(10, Puzzle010{})
}

type Puzzle010 struct{}

func (p Puzzle010) Solve() puzzle.Answer {
	return puzzle.FromValue(p.sumPrimesBelow(2_000_000))
}

func (p Puzzle010) sumPrimesBelow(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		if common.IsPrime(i) {
			sum += i
		}
	}
	return sum
}
