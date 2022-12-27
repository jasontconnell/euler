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

	// digits := []int{b}
	// for i := 1; i < exp; i++ {
	// 	carry := 0
	// 	for j := len(digits) - 1; j >= 0; j-- {
	// 		x := digits[j]*2 + carry
	// 		digit := x % 10
	// 		carry = x / 10
	// 		digits[j] = digit
	// 	}
	// 	c := carry
	// 	for c > 0 {
	// 		digit := c % 10
	// 		digits = append([]int{digit}, digits...)
	// 		c /= 10
	// 	}
	// }

	// return common.Sum(digits)
}
