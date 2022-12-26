package puzzles

import (
	"github.com/jasontconnell/euler/common"
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(3, Puzzle003{})
}

type Puzzle003 struct{}

func (p Puzzle003) Solve() puzzle.Answer {
	return puzzle.FromValue(p.largestPrimeFactor(600_851_475_143))
}

func (p Puzzle003) largestPrimeFactor(n int) int {
	max := 0
	for i := 2; i < common.Sqrt(n)+1; i++ {
		if n%i == 0 {
			if common.IsPrime(i) && i > max {
				max = i
			}
		}
	}
	return max
}
