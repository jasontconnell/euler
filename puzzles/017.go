package puzzles

import (
	"github.com/jasontconnell/euler/puzzle"
)

func init() {
	puzzle.Register(17, Puzzle017{})
}

type Puzzle017 struct{}

// maps numbers to the number of letters in the word
// like 7 is "seven" or 5 letters
var letters map[int]int = map[int]int{
	0: 0, // it won't map, we don't say the 0 in 302 and 50 and 10
	1: 3, 10: 3, 100: 3 + 7,
	2: 3, 20: 6, 200: 3 + 7,
	3: 5, 30: 6, 300: 5 + 7,
	4: 4, 40: 5, 400: 4 + 7,
	5: 4, 50: 5, 500: 4 + 7,
	6: 3, 60: 5, 600: 3 + 7,
	7: 5, 70: 7, 700: 5 + 7,
	8: 5, 80: 6, 800: 5 + 7,
	9: 4, 90: 6, 900: 4 + 7,

	11: 6, 12: 6, 13: 8, 14: 8, 15: 7, 16: 7, 17: 9, 18: 8, 19: 8,
	1000: 11,
}

// used for debugging
var words map[int]string = map[int]string{
	0: "", // it won't map, we don't say the 0 in 302 and 50 and 10
	1: "one", 10: "ten", 100: "onehundred",
	2: "two", 20: "twenty", 200: "twohundred",
	3: "three", 30: "thirty", 300: "threehundred",
	4: "four", 40: "forty", 400: "fourhundred",
	5: "five", 50: "fifty", 500: "fivehundred",
	6: "six", 60: "sixty", 600: "sixhundred",
	7: "seven", 70: "seventy", 700: "sevenhundred",
	8: "eight", 80: "eighty", 800: "eighthundred",
	9: "nine", 90: "ninety", 900: "ninehundred",

	11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen",
	16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen",
	1000: "onethousand",
}

func (p Puzzle017) Solve() puzzle.Answer {
	return puzzle.FromValue(p.getLetterCounts(1, 1000))
}

func (p Puzzle017) getLetterCounts(start, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		sum += p.getCountForN(i)
	}
	return sum
}

func (p Puzzle017) getCountForN(i int) int {
	if ns, ok := letters[i]; ok {
		return ns
	}

	cur := i
	s := 10
	if i > 9 {
		s = 100
	} else if i > 99 {
		s = 1000
	}

	local := 0
	done := false
	for s > 0 && !done {
		clean := cur
		if s != 0 {
			clean = clean - (cur % s)
		}

		if ns, ok := letters[cur]; ok {
			local += ns
			done = true
		} else if clean%100 == 0 && clean >= 100 {
			local += letters[clean]
			if cur%100 != 0 {
				local += 3
			}
		} else if clean%10 == 0 {
			d := clean % 10
			local += letters[clean] + letters[d]
		}

		cur %= s
		s /= 10
	}
	return local
}

func (p Puzzle017) asWords(i int) string {
	cur := i
	s := 10
	if i > 9 {
		s = 100
	} else if i > 99 {
		s = 1000
	}

	local := ""
	done := false
	for s > 0 && !done {
		clean := cur
		if s != 0 {
			clean = clean - (cur % s)
		}

		if ns, ok := words[cur]; ok {
			local += ns
			done = true
		} else if clean%100 == 0 && clean >= 100 {
			local += words[clean]
			if cur%100 != 0 {
				local += "and"
			}
		} else if clean%10 == 0 {
			d := clean % 10
			local += words[clean] + words[d]
		}

		cur %= s
		s /= 10
	}
	return local
}
