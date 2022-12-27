package puzzle

import (
	"fmt"
	"sort"
)

type Registry struct {
	puzzles map[int]Solver
}

type Puzzle struct {
	Num    int
	Solver Solver
}

var registry *Registry

func newRegistry() *Registry {
	r := new(Registry)
	r.puzzles = make(map[int]Solver)
	return r
}

func ensureRegistry() {
	if registry == nil {
		registry = newRegistry()
	}
}

func Register(n int, s Solver) {
	ensureRegistry()
	registry.puzzles[n] = s
}

func Solve(n int) Answer {
	ensureRegistry()
	s, ok := registry.puzzles[n]
	if !ok {
		return FromValue(fmt.Sprintf("No solution for puzzle %d", n))
	}
	return s.Solve()
}

func AllPuzzles() []Puzzle {
	ensureRegistry()
	list := []int{}
	for k := range registry.puzzles {
		list = append(list, k)
	}
	sort.Ints(list)
	solvers := []Puzzle{}
	for _, n := range list {
		solvers = append(solvers, Puzzle{Num: n, Solver: registry.puzzles[n]})
	}
	return solvers
}
