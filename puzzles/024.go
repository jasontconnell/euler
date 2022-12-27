package puzzles

import (
	"fmt"
	"math"
	"sort"

	"github.com/jasontconnell/euler/common"
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(24, Puzzle024{})
}

type Puzzle024 struct{}

func (p Puzzle024) Solve() puzzle.Answer {
	return puzzle.FromValue(p.getNthPerm(1_000_000, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func (p Puzzle024) getNthPerm(n int, nums []int) string {
	perms := common.Permutate(nums)
	sort.Slice(perms, func(i, j int) bool {
		si, sj := 0, 0
		for xsi, d := range perms[i] {
			si += int(math.Pow10((len(perms[i]) - xsi))) * d
		}

		for xsj, d := range perms[j] {
			sj += int(math.Pow10((len(perms[j]) - xsj))) * d
		}
		return si < sj
	})

	s := ""
	for _, d := range perms[n-1] {
		s += fmt.Sprintf("%d", d)
	}
	return s
}
