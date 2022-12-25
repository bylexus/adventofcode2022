package problems

import (
	"fmt"
	"math"

	"alexi.ch/aoc/2022/lib"
)

type snafu []int
type decimal int64

func (s snafu) toDec() decimal {
	var nr decimal = 0
	for i := 0; i < len(s); i++ {
		var pow = len(s) - i - 1
		var digit = s[i]
		if digit > 0 {
			nr += decimal(digit) * decimal(math.Pow(5, float64(pow)))
		} else if digit < 0 {
			nr -= decimal(-1*digit) * decimal(math.Pow(5, float64(pow)))
		}
	}
	return nr
}

func (s snafu) String() string {
	var res = ""
	for _, i := range s {
		switch i {
		case 2:
			res += "2"
		case 1:
			res += "1"
		case 0:
			res += "0"
		case -1:
			res += "-"
		case -2:
			res += "="
		}
	}
	return res
}

func (d decimal) toSnafu() snafu {
	var res = make(snafu, 0)

	for {
		if d == 0 {
			break
		}
		var remainder = d % 5
		if remainder > 2 {
			d += remainder
			if remainder == 3 {
				res = append(res, -2)
			} else {
				res = append(res, -1)
			}
		} else {
			res = append(res, int(remainder))
		}
		d = d / 5
	}

	// reverse:
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

type Day25 struct {
	s1            string
	s2            decimal
	snafu_numbers []snafu
}

func NewDay25() Day25 {
	return Day25{s1: "", s2: 0, snafu_numbers: make([]snafu, 0)}
}

func (d *Day25) Title() string {
	return "Day 25 - Full of Hot Air"
}

func (d *Day25) Setup() {
	// var lines = lib.ReadLines("data/25-test.txt")
	var lines = lib.ReadLines("data/25-data.txt")
	for _, line := range lines {
		if len(line) > 0 {
			var snafu = make(snafu, len(line))
			for i, c := range line {
				switch c {
				case '2':
					snafu[i] = 2
				case '1':
					snafu[i] = 1
				case '0':
					snafu[i] = 0
				case '-':
					snafu[i] = -1
				case '=':
					snafu[i] = -2
				}
			}
			d.snafu_numbers = append(d.snafu_numbers, snafu)
		}
	}
	// fmt.Printf("%v\n", d.snafu_numbers)
	// for _, snafu := range d.snafu_numbers {
	// 	fmt.Printf("%v: %d\n", snafu, snafu.toDec())
	// }
}

func (d *Day25) SolveProblem1() {
	var sum decimal = 0
	for _, snafu := range d.snafu_numbers {
		sum += snafu.toDec()
		// fmt.Printf("%s: %d ==> %s\n", snafu, snafu.toDec(), snafu.toDec().toSnafu())
	}
	d.s1 = sum.toSnafu().String()
}

func (d *Day25) SolveProblem2() {
	d.s2 = 0
}

func (d *Day25) Solution1() string {
	return fmt.Sprintf("%s", d.s1)
}

func (d *Day25) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
