package puzzles

import (
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(1, Puzzle001{})
}

type Puzzle001 struct{}

func (p Puzzle001) Solve() puzzle.Answer {
	return puzzle.FromValue(sum(0, 1000, 3, 5))
}

func sum(start, end int, multiples ...int) int {
	s := 0
	for i := start; i < end; i++ {
		for _, j := range multiples {
			if i%j == 0 {
				s += i
				break
			}
		}
	}
	return s
}
