package problems

import (
	"fmt"
	"time"
)

type Problem interface {
	Title() string
	Setup()
	SolveProblem1()
	SolveProblem2()
	Solution1() string
	Solution2() string
}

func Solve(p Problem) {
	var start = time.Now()
	p.Setup()
	var setupTime = time.Now().Sub(start)

	var start1 = time.Now()
	p.SolveProblem1()
	var solve1Time = time.Now().Sub(start1)

	var start2 = time.Now()
	p.SolveProblem2()
	var end = time.Now()
	var solve2Time = end.Sub(start2)

	fmt.Printf("\n\n***%s***\n", p.Title())
	fmt.Printf("  Setup time: %s\n", setupTime)
	fmt.Printf("  \x1B[1;97mSolution 1: %s\n\x1B[0m", p.Solution1())
	fmt.Printf("        took: %s\n", solve1Time)
	fmt.Printf("  \x1B[1;97mSolution 2: %s\n\x1B[0m", p.Solution2())
	fmt.Printf("        took: %s\n", solve2Time)
	fmt.Printf("         all: %s\n", end.Sub(start))
}
