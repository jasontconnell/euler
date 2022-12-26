package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jasontconnell/euler/puzzle"
	_ "github.com/jasontconnell/euler/puzzles"
)

func main() {
	pn := flag.Int("num", 0, "puzzle number")
	flag.Parse()

	start := time.Now()
	ans := puzzle.Solve(*pn)
	fmt.Println("Answer:", ans)
	fmt.Println("Time", time.Since(start))
}
