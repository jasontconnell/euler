### Project Euler ###

This is my project to solve puzzles on [Project Euler](projecteuler.net)

Init a new code file for a puzzle with `.\init.ps1 [num]`

For instance, to start a new file for puzzle 37, from the root of the repo, type `.\init.ps1 37`

This will create a file `037.go` which will contain the following (as of this writing)

```
package puzzles

import (
	"github.com/jasontconnell/euler/puzzle"
)

func init(){
	puzzle.Register(37, Puzzle037{})
}

type Puzzle037 struct {}

func (p Puzzle037) Solve() puzzle.Answer {
	return puzzle.FromValue(0)
}
```

Each puzzle registers itself with the main program, and the main program can then just call
```
puzzle.Solve(37)
```

You can compile the project with `go build` but I typically just run `main.go`. Pass in the puzzle number to solve like `go run main.go 37` or if compiled, `.\euler.exe 37`.

There are currently 822 puzzles on Project Euler so the 00n naming convention is necessary.

The puzzle returns an `Answer` object which is usually just initialized with `puzzle.FromAnswer([the answer])` and prints out the answer and the time it took to solve it.

Feel free to learn from this but the overall purpose of Project Euler is to learn how to solve the problems, not to just input a correct answer for each one :)