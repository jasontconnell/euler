package puzzles

import (
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(14, Puzzle014{})
}

type Puzzle014 struct{}

func (p Puzzle014) Solve() puzzle.Answer {
	return puzzle.FromValue(p.getLongestCollatzSequence(1_000_000))
}

func (p Puzzle014) getLongestCollatzSequence(maxstart int) int {
	m := make(map[int]int)
	max := 0
	num := 0

	for i := 1; i <= maxstart; i++ {
		c := p.getCollatzSequenceLength(i, m)
		if c > max {
			max = c
			num = i
		}
	}
	return num
}

func (p Puzzle014) getCollatzSequenceLength(n int, m map[int]int) int {
	x := n
	seq := 1 // include n
	for x != 1 {
		if c, ok := m[x]; ok {
			seq = seq + c
			m[n] = seq
			break
		}

		switch x % 2 {
		case 0:
			x /= 2
		case 1:
			x = x*3 + 1
		}
		seq++
	}
	m[n] = seq
	return seq
}
