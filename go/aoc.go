package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"alexi.ch/aoc/2022/problems"
)

func main() {
	tannenbaum()
	var problem_map = make(map[string](problems.Problem))

	var day01 = problems.NewDay01()
	problem_map["01"] = &day01
	var day11 = problems.NewDay11()
	problem_map["11"] = &day11
	var day12 = problems.NewDay12()
	problem_map["12"] = &day12
	var day13 = problems.NewDay13()
	problem_map["13"] = &day13
	var day14 = problems.NewDay14()
	problem_map["14"] = &day14
	var day15 = problems.NewDay15()
	problem_map["15"] = &day15

	var to_solve = make([]string, 0)
	for _, arg := range os.Args[1:] {
		to_solve = append(to_solve, arg)
	}

	if len(to_solve) == 0 {
		var keys = make([]string, 0)
		for key := range problem_map {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		to_solve = keys
	}

	var start = time.Now()
	for _, p := range to_solve {
		var prob = problem_map[p]
		if prob != nil {
			problems.Solve(prob)
		} else {
			panic("Problem not found")
		}
	}
	var duration = time.Now().Sub(start)
	fmt.Printf("\n\nFull runtime: %s\n\n", duration)
}

func tannenbaum() {
	var t = strings.Join([]string{
		"\x1B[1;97m",
		"Advent of Code 2022",
		"--------------------",
		"",
		"        \x1B[1;93m*   *",
		"         \\ /",
		"         AoC",
		"         -\x1B[1;91m*\x1B[1;93m-",
		"          \x1B[1;37m|\x1B[0;32m",
		"          *",
		"         /*\\",
		"        /\x1B[1;94m*\x1B[0;32m*\x1B[1;93m*\x1B[0;32m\\",
		"       /\x1B[1;91m*\x1B[0;32m***\x1B[1;94m*\x1B[0;32m\\",
		"      /**\x1B[1;93m*\x1B[0;32m****\\",
		"     /**\x1B[1;94m*\x1B[0;32m***\x1B[1;91m*\x1B[0;32m**\\",
		"    /********\x1B[1;93m*\x1B[0;32m**\\",
		"   /**\x1B[1;91m*\x1B[0;32m*****\x1B[1;94m*\x1B[0;32m****\\",
		"  /**\x1B[1;94m*\x1B[0;32m*\x1B[1;93m*\x1B[0;32m**********\\",
		" /**\x1B[1;94m*\x1B[0;32m*****\x1B[1;93m*\x1B[0;32m**\x1B[1;91m*\x1B[0;32m****\x1B[1;93m*\x1B[0;32m\\",
		"          #",
		"          #",
		"       \x1B[1;97m2-0-2-2",
		"       #######",
		"\x1B[0m",
	}, "\n")
	fmt.Printf(t)
}
