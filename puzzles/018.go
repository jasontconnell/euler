package puzzles

import (
	"strconv"
	"strings"

	"github.com/jasontconnell/euler/common"
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(18, Puzzle018{})
}

var input18 string = `
              75
             95 64
            17 47 82
           18 35 87 10
          20 04 82 47 65
         19 01 23 75 03 34
        88 02 77 73 07 63 67
       99 65 04 28 06 16 70 92
      41 41 26 56 83 40 80 70 33
     41 48 72 33 47 32 37 16 94 29
    53 71 44 65 25 43 91 52 97 51 14
   70 11 33 28 77 73 17 78 39 68 17 57
  91 71 52 38 17 14 91 43 58 50 27 29 48
 63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

type Puzzle018 struct{}

type p18state struct {
	pt  common.Point2D
	sum int
}

func (p Puzzle018) Solve() puzzle.Answer {
	return puzzle.FromValue(p.getMaxPath(p.getPyramid(input18)))
}

func (p Puzzle018) getMaxPath(pyramid [][]int) int {
	max := 0
	// either move left (x=0) or right(x = 1) and down(y+1)

	start := common.Point2D{X: 0, Y: 0}
	queue := []p18state{{pt: start, sum: 0}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.pt.Y > len(pyramid)-1 {
			continue
		}

		row := pyramid[cur.pt.Y]
		if cur.pt.X > len(row)-1 {
			continue
		}

		cur.sum += pyramid[cur.pt.Y][cur.pt.X]

		if cur.pt.Y == len(pyramid)-1 {
			if cur.sum > max {
				max = cur.sum
			}
			continue
		}

		mvs := []common.Point2D{{X: cur.pt.X, Y: cur.pt.Y + 1}, {X: cur.pt.X + 1, Y: cur.pt.Y + 1}}
		for _, mv := range mvs {
			queue = append(queue, p18state{pt: mv, sum: cur.sum})
		}
	}

	return max
}

func (p Puzzle018) getPyramid(s string) [][]int {
	sp := strings.Split(strings.Trim(s, "\n "), "\n")
	pyramid := make([][]int, len(sp))
	for y, line := range sp {
		if len(strings.Trim(line, "\n ")) == 0 {
			continue
		}
		flds := strings.Fields(line)
		pyramid[y] = make([]int, len(flds))
		for x, f := range flds {
			n, _ := strconv.Atoi(f)
			pyramid[y][x] = n
		}
	}
	return pyramid
}
