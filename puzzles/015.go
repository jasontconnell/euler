package puzzles

import (
	"github.com/jasontconnell/euler/common"
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(15, Puzzle015{})
}

type Puzzle015 struct{}

func (p Puzzle015) Solve() puzzle.Answer {
	return puzzle.FromValue(p.findPaths(20))
}

func (p Puzzle015) findPaths(n int) int {
	g := make(map[common.Point2D]bool)
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			pt := common.Point2D{X: x, Y: y}
			g[pt] = true
		}
	}

	start := common.Point2D{X: 0, Y: 0}
	goal := common.Point2D{X: n, Y: n}
	visit := make(map[common.Point2D]int)

	// break it down into solvable areas
	for y := n - 1; y >= 0; y-- {
		for x := n - 1; x >= 0; x-- {
			pt := common.Point2D{X: x, Y: y}
			num := p.traverse(pt, goal, g, visit)
			visit[pt] = num
		}
	}

	return visit[start]
}

func (p Puzzle015) traverse(start, goal common.Point2D, g map[common.Point2D]bool, visit map[common.Point2D]int) int {
	q := []common.Point2D{start}
	num := 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if cur.X == goal.X || cur.Y == goal.Y {
			// no other possible variations left
			// once you get perpendicular to the goal
			num++
			continue
		}

		if c, ok := visit[cur]; ok {
			num += c
			continue
		}

		if cur == goal {
			num++
			continue
		}

		mvs := []common.Point2D{{X: cur.X + 1, Y: cur.Y}, {X: cur.X, Y: cur.Y + 1}}
		for _, mv := range mvs {
			if mv.X > goal.X || mv.X < 0 || mv.Y > goal.Y || mv.Y < 0 {
				continue
			}
			q = append(q, mv)
		}
	}
	return num
}
