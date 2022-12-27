package puzzles

import (
	"github.com/jasontconnell/euler/common"
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(12, Puzzle012{})
}

type Puzzle012 struct{}

func (p Puzzle012) Solve() puzzle.Answer {
	return puzzle.FromValue(p.maxDivisibleTriangleNum(500))
}

func (p Puzzle012) maxDivisibleTriangleNum(min int) int {
	s := 1
	cur := 2
	result := 0

	for {
		d := common.Divisors(s, false)
		if len(d) > min {
			result = s
			break
		}
		s += cur
		cur++
	}
	return result
}
