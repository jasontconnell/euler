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
	if len(os.Args) == 1 {
		log.Fatal("need puzzle number")
	}
	snum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("invalid puzzle number", os.Args[1])
	}

	start := time.Now()
	ans := puzzle.Solve(snum)
	fmt.Println("Answer:", ans)
	fmt.Println("Time", time.Since(start))
}
