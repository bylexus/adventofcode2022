package problems

import (
	"fmt"
)

type Day22 struct {
	s1 uint64
	s2 uint64
}

func NewDay22() Day22 {
	return Day22{s1: 0, s2: 0}
}

func (d *Day22) Title() string {
	return "Day 22 - xxx"
}

func (d *Day22) Setup() {
	// var lines = lib.ReadLines("data/22-test.txt")
	// var lines = lib.ReadLines("data/22-data.txt")
	// for _, line := range lines {
	// 	line = line
	// }
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day22) SolveProblem1() {
	d.s1 = 0
}

func (d *Day22) SolveProblem2() {
	d.s2 = 0
}

func (d *Day22) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day22) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
