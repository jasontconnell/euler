package puzzles

import (
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(3, Puzzle003{})
}

type Puzzle003 struct{}

func (p Puzzle003) Solve() puzzle.Answer {
	return puzzle.FromValue(largestPrimeFactor(600_851_475_143))
}

func largestPrimeFactor(n int) int {
	max := 0
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			if isPrime(i) && i > max {
				max = i
			}
			if isPrime(n / i) {
				break
			}
		}
	}
	return max
}

func isPrime(n int) bool {
	if n == 1 {
		return true
	}
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func factors(n int) []int {
	list := []int{}

	return list
}
