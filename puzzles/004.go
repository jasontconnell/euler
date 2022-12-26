package puzzles

import (
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(4, Puzzle004{})
}

type Puzzle004 struct{}

func (p Puzzle004) Solve() puzzle.Answer {
	return puzzle.FromValue(p.largestPalendrome(100, 999))
}

func (p Puzzle004) largestPalendrome(start, end int) int {
	result := 0
	for i := end; i >= start; i-- {
		for j := end; j >= start; j-- {
			if p.isPalendrome(i*j) && i*j > result {
				result = i * j
			}
		}
	}
	return result
}

func (p Puzzle004) isPalendrome(n int) bool {
	digits := []int{}
	x := n
	for x > 0 {
		d := x % 10
		digits = append(digits, d)
		x /= 10
	}
	pd := true
	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)-1-i] {
			pd = false
			break
		}
	}
	return pd
}
