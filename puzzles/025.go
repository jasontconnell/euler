package puzzles

import (
	"github.com/jasontconnell/euler/common"
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(25, Puzzle025{})
}

type Puzzle025 struct{}

func (p Puzzle025) Solve() puzzle.Answer {
	return puzzle.FromValue(p.getFibWithNDigits(1000))
}

func (p Puzzle025) getNthFib(n int) int {
	fibN := 0
	common.BigNumberFib(func(idx int, digits []int) bool {
		cont := len(digits) < n
		if !cont {
			fibN = idx
		}
		return cont
	})

	return fibN
}

func (p Puzzle025) getFibWithNDigits(n int) int {
	fibN := 0
	common.BigNumberFib(func(idx int, digits []int) bool {
		cont := len(digits) < n
		if !cont {
			fibN = idx
		}
		return cont
	})

	return fibN
}
