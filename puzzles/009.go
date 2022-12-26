package puzzles

import (
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(9, Puzzle009{})
}

type Puzzle009 struct{}

func (p Puzzle009) Solve() puzzle.Answer {
	return puzzle.FromValue(p.findSpecialPythagorean(1000))
}

func (p Puzzle009) findSpecialPythagorean(n int) int {
	max := n
	ans := 0
	found := false
	for i := 1; i <= max && !found; i++ {
		for j := i + 1; j <= max && !found; j++ {
			for k := j + 1; k <= max && !found; k++ {
				if i*i+j*j == k*k && i+j+k == n {
					ans = i * j * k
					found = true
				}
			}
		}
	}
	return ans
}
