package puzzles

import (
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(2, Puzzle002{})
}

type Puzzle002 struct{}

func (p Puzzle002) Solve() puzzle.Answer {
	return puzzle.FromValue(sumEvenFibs(4_000_000))
}

func sumEvenFibs(high int) int {
	prev := 1
	cur := 1
	sum := 0
	for cur < high {
		tmp := cur
		cur = cur + prev
		prev = tmp

		if cur%2 == 0 {
			sum += cur
		}
	}
	return sum
}
