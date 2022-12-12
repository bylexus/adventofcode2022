# Advent of Code 2022

Eagerly, we're all awaiting [Advent of Code, Edition 2022!](https://adventofcode.com/2022/). Finally, it's here.
What a magical time of the year!

My AoC 2022 solutions in GO, after putting Rust away.

Alex

## Run problems

All problems can be run by its day index, e.g:

```
$ go run 01
```

or all together:

```
$ go run
```

## Define Problems

1) Create a struct in the `problems` package that implements the `Problem` interface, e.g.:

```go
package problems

import (
	"fmt"
	"alexi.ch/aoc/2022/lib"
)

type DayXX struct {
	s1 uint64
	s2 uint64
}

func NewDayXX() DayXX {
	return DayXX{s1: 0, s2: 0}
}

func (d *DayXX) Title() string {
	return "Day XX - Title comes here"
}

func (d *DayXX) Setup() {
	// var lines = lib.ReadLines("data/01-test.txt")
	var lines = lib.ReadLines("data/01-data.txt")
	for _, line := range lines {
		line = line
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *DayXX) SolveProblem1() {
	d.s1 = 0
}

func (d *DayXX) SolveProblem2() {
	d.s2 = 0
}

func (d *DayXX) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *DayXX) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
```

2) import and instantiate the struct in the main program `aoc.go`:

```go
// aoc.go
var dayXX = problems.NewDayXX()
problem_map["XX"] = &dayXX
```

