package problems

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"alexi.ch/aoc/2022/lib"
)

type Day01 struct {
	numbers []uint64
	sums    []uint64
	s1      uint64
	s2      uint64
}

func NewDay01() Day01 {
	return Day01{s1: 0, s2: 0, numbers: make([]uint64, 0), sums: make([]uint64, 0)}
}

func (d *Day01) Title() string {
	return "Day 01 - Calorie Counting"
}

func (d *Day01) Setup() {
	// var lines = lib.ReadLines("data/01-test.txt")
	var lines = lib.ReadLines("data/01-data.txt")
	for _, line := range lines {
		nr, err := strconv.ParseUint(strings.TrimSpace(line), 10, 64)
		if err != nil {
			nr = 0
		}
		d.numbers = append(d.numbers, nr)
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day01) SolveProblem1() {
	var sums = make([]uint64, 0)
	var localSum uint64 = 0

	for _, nr := range d.numbers {
		if nr > 0 {
			localSum += nr
		} else {
			sums = append(sums, localSum)
			localSum = 0
		}
	}
	max, err := lib.FindMax(sums)
	d.sums = sums
	if err == nil {
		d.s1 = *max
	}
}

func (d *Day01) SolveProblem2() {
	// sums already calculated in solve1
	sort.Slice(d.sums, func(i int, j int) bool {
		return d.sums[i] > d.sums[j]
	})
	d.s2 = lib.Sum(d.sums[:3])
}

func (d *Day01) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day01) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
