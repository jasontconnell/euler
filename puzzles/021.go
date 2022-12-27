package puzzles

import (
	"github.com/jasontconnell/euler/common"
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(21, Puzzle021{})
}

type Puzzle021 struct{}

func (p Puzzle021) Solve() puzzle.Answer {
	return puzzle.FromValue(p.sumAmicableNums(10000))
}

func (p Puzzle021) sumAmicableNums(max int) int {
	visit := make(map[int]bool)
	sum := 0
	for i := 1; i < max; i++ {
		if _, ok := visit[i]; ok {
			continue
		}
		if common.IsPrime(i) {
			continue
		}
		diva := common.ProperDivisors(i)
		suma := common.Sum(diva)

		if common.IsPrime(suma) {
			continue
		}

		divb := common.ProperDivisors(suma)
		sumb := common.Sum(divb)

		if sumb == i && suma != sumb {
			sum += suma + sumb
			visit[suma] = true
			visit[sumb] = true
		}
	}
	return sum
}
