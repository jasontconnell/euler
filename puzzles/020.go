package puzzles

import (
	"github.com/jasontconnell/euler/common"
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(20, Puzzle020{})
}

type Puzzle020 struct{}

func (p Puzzle020) Solve() puzzle.Answer {
	return puzzle.FromValue(p.getFactDigitSum(100))
}

func (p Puzzle020) getFactDigitSum(n int) int {
	list := common.BigNumberCalc(1, n, func(idx, n int) int {
		return n * idx
	})

	return common.Sum(list)
}
