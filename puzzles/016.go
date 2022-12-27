package puzzles

import (
	"github.com/jasontconnell/euler/common"
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(16, Puzzle016{})
}

type Puzzle016 struct{}

func (p Puzzle016) Solve() puzzle.Answer {
	return puzzle.FromValue(p.getPowerDigitSum(2, 1000))
}

func (p Puzzle016) getPowerDigitSum(b, exp int) int {
	list := common.BigNumberCalc(b, exp, func(idx int, n int) int {
		return n * 2
	})
	return common.Sum(list)
}
