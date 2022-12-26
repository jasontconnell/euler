package puzzles

import (
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(5, Puzzle005{})
}

type Puzzle005 struct{}

func (p Puzzle005) Solve() puzzle.Answer {
	return puzzle.FromValue(p.smallestPossibleDivisibleByRange(1, 20))
}

func (p Puzzle005) smallestPossibleDivisibleByRange(start, end int) int {
	n := 1
	for i := start; i <= end; i++ {
		n *= i
	}

	result := 0
	for i := end; i <= n; i++ {
		c := true
		for j := start; j <= end; j++ {
			if i%j != 0 {
				c = false
				break
			}
		}
		if c {
			result = i
			break
		}
	}
	return result
}
