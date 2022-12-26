package puzzle

type Registry struct {
	puzzles map[int]Solver
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
	s, ok := registry.puzzles[n]
	if !ok {
		return Answer{}
	}
	return s.Solve()
}
