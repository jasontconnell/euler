package puzzles

import (
	"github.com/jasontconnell/euler/common"
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(23, Puzzle023{})
}

type Puzzle023 struct{}

func (p Puzzle023) Solve() puzzle.Answer {
	return puzzle.FromValue(p.sumNonAbundantNums(28124))
}

func (p Puzzle023) getAbundantsMap(max int) map[int]int {
	m := make(map[int]int)
	for i := 1; i < max; i++ {
		if common.IsPrime(i) {
			continue
		}
		d := common.ProperDivisors(i)
		s := common.Sum(d)

		if s > i {
			m[i] = i
		}
	}
	return m
}

func (p Puzzle023) sumNonAbundantNums(max int) int {
	m := p.getAbundantsMap(max)

	sum := 0
	for i := 1; i < max; i++ {
		canBeWritten := p.hasAbundantAddends(i, m)

		if !canBeWritten {
			sum += i
		}
	}
	return sum
}

func (p Puzzle023) hasAbundantAddends(n int, abundants map[int]int) bool {
	result := false
	for j := 1; j <= n/2; j++ {
		if abundants[j] != 0 && abundants[n-j] != 0 {
			result = true
			break
		}
	}
	return result
}
