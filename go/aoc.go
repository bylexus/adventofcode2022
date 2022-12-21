package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"alexi.ch/aoc/2022/problems"
)

func main() {
	tannenbaum()
	var problem_map = map[string](func() problems.Problem){
		"01": func() problems.Problem { p := problems.NewDay01(); return &p },
		"11": func() problems.Problem { p := problems.NewDay11(); return &p },
		"12": func() problems.Problem { p := problems.NewDay12(); return &p },
		"13": func() problems.Problem { p := problems.NewDay13(); return &p },
		"14": func() problems.Problem { p := problems.NewDay14(); return &p },
		"15": func() problems.Problem { p := problems.NewDay15(); return &p },
		"16": func() problems.Problem { p := problems.NewDay16(); return &p },
		"17": func() problems.Problem { p := problems.NewDay17(); return &p },
		"18": func() problems.Problem { p := problems.NewDay18(); return &p },
		"19": func() problems.Problem { p := problems.NewDay19(); return &p },
		"20": func() problems.Problem { p := problems.NewDay20(); return &p },
		"21": func() problems.Problem { p := problems.NewDay21(); return &p },
		"22": func() problems.Problem { p := problems.NewDay22(); return &p },
	}

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

	// Run solving all problems  in parallel:
	var start = time.Now()
	var wg sync.WaitGroup
	for _, p := range to_solve {
		wg.Add(1)
		go func(probKey string) {
			defer wg.Done()
			var prob = problem_map[probKey]
			if prob != nil {
				problems.Solve(prob())
			} else {
				panic("Problem not found")
			}
		}(p)
	}
	wg.Wait()
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
