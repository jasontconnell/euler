package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jasontconnell/euler/puzzle"
	_ "github.com/jasontconnell/euler/puzzles"
)

func main() {
	start := time.Now()
	if len(os.Args) == 1 {
		all := puzzle.AllPuzzles()
		for _, puzzle := range all {
			fmt.Println("Puzzle", puzzle.Num)
			fmt.Println("Answer:", puzzle.Solver.Solve())
		}
	} else {
		snum, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("invalid puzzle number", os.Args[1])
		}

		ans := puzzle.Solve(snum)
		fmt.Println("Answer:", ans)
	}
	fmt.Println("Time", time.Since(start))
}
