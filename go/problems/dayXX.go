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
	return "Day 01 - Calorie Counting"
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
