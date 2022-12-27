package puzzles

import (
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(19, Puzzle019{})
}

type Puzzle019 struct{}

var days map[int]int = map[int]int{
	1: 31, 2: 28, 3: 31, 4: 30, 5: 31, 6: 30, 7: 31, 8: 31, 9: 30, 10: 31, 11: 30, 12: 31,
}

const (
	monday int = iota
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

func (p Puzzle019) Solve() puzzle.Answer {
	return puzzle.FromValue(p.countSundays(1901, 1, 1, 2000, 12, 31))
}

func (p Puzzle019) countSundays(sy, sm, sd, ey, em, ed int) int {
	refy := 1900
	refd := 1
	refm := 1

	cury, curd, curm := refy, refd, refm

	counting := false
	dow := 0 // ref date starts on monday
	done := false
	sundays := 0

	for !done {
		counting = counting || cury == sy && curm == sm && curd == sd
		if counting && dow == sunday && curd == 1 {
			sundays++
		}
		done = cury == ey && curm == em && curd == ed

		curd++
		dtotal := days[curm]

		if p.isLeap(cury) && curm == 2 {
			dtotal = 29
		}

		if curd > dtotal {
			curm = (curm + 1) % 13
			if curm == 0 {
				curm = 1
			}
			curd = 1
		}

		if curm == 1 && curd == 1 {
			cury++
		}

		dow = (dow + 1) % 7

	}
	return sundays
}

func (p Puzzle019) daysInYear(y int) int {
	d := 0
	for _, m := range days {
		d += m
	}
	if p.isLeap(y) {
		d++
	}
	return d
}

func (p Puzzle019) isLeap(y int) bool {
	if y%100 == 0 {
		return y%400 == 0
	}
	return y%4 == 0
}
