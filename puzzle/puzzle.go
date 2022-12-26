package puzzle

import "fmt"

type Solver interface {
	Solve() Answer
}

type Answer struct {
	value  string
	Solved bool
}

func FromValue(value interface{}) Answer {
	a := Answer{}
	a.SetValue(value)
	return a
}

func (a *Answer) SetValue(value interface{}) {
	a.value = fmt.Sprintf("%v", value)
}

func (a Answer) String() string {
	return a.value
}
