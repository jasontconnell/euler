package puzzles

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jasontconnell/euler/common"
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(22, Puzzle022{})
}

type Puzzle022 struct{}

func (p Puzzle022) Solve() puzzle.Answer {
	input, err := common.ReadFile("./input/022.in")
	if err != nil {
		fmt.Println(err)
		return puzzle.FromValue("error reading file " + err.Error())
	}
	return puzzle.FromValue(p.sumNameScores(input))
}

func (p Puzzle022) sumNameScores(str string) int {
	names := strings.Split(str, ",")
	for i, name := range names {
		n := name[1 : len(name)-1]
		names[i] = n
	}

	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})

	sum := 0
	for i, name := range names {
		sum += (i + 1) * p.getNameScore(name)
	}

	return sum
}

func (p Puzzle022) getNameScore(str string) int {
	sc := 0
	for _, c := range str {
		sc += int(c) - 64
	}
	return sc
}
