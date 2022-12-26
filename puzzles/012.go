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
		d := p.divisors(s)
		if d > min {
			result = s
			break
		}
		s += cur
		cur++
	}
	return result
}

func (p Puzzle012) divisors(n int) int {
	count := 0
	if common.IsPrime(n) {
		return 2
	}

	for i := common.Sqrt(n); i >= 2; i-- {
		if n%i == 0 {
			count += 2 // i + i*?
		}
	}
	// count 1 and n
	return count + 2
}
