package problems

import (
	"fmt"
	"regexp"
	"strconv"

	"alexi.ch/aoc/2022/lib"
)

type Expression21 interface {
	value() (int64, bool)
	needsMonkey(string) bool
}

type Operation21 struct {
	op      rune
	left    string
	right   string
	monkeys *Monkeys21
}

func (this *Operation21) value() (int64, bool) {
	// Calculate the value of this operation:
	// Make sure the value is valid, by
	// checking the validity of sub-operations,
	// and to test division
	var left = (*this.monkeys)[this.left]
	var right = (*this.monkeys)[this.right]

	lval, ok := left.value()
	if ok == false {
		return 0, false
	}
	rval, ok := right.value()
	if ok == false {
		return 0, false
	}
	switch this.op {
	case '+':
		return lval + rval, true
	case '-':
		return lval - rval, true
	case '*':
		return lval * rval, true
	case '/':
		// only valid if division is possible without remainder:
		if lval%rval != 0 {
			return 0, false
		}
		return lval / rval, true
	default:
		panic("Unknown operation!")
	}
}

func (this *Operation21) needsMonkey(name string) bool {
	if this.left == name {
		return true
	}
	if this.right == name {
		return true
	}
	var left = (*this.monkeys)[this.left]
	var right = (*this.monkeys)[this.right]
	return left.needsMonkey(name) || right.needsMonkey(name)
}

type Value21 struct {
	nr int64
}

func (this *Value21) value() (int64, bool) {
	return this.nr, true
}

func (this *Value21) needsMonkey(name string) bool {
	return false
}

type HumnValue21 struct {
	nr int64
}

func (this *HumnValue21) value() (int64, bool) {
	return this.nr, true
}

func (this *HumnValue21) needsMonkey(name string) bool {
	return false
}
func (this *HumnValue21) setValue(value int64) {
	this.nr = value
}

type Monkeys21 map[string]Expression21

type Day21 struct {
	s1       int64
	s2       int64
	monkeys  Monkeys21
	monkeys2 Monkeys21
}

func NewDay21() Day21 {
	return Day21{s1: 0, s2: 0, monkeys: make(Monkeys21), monkeys2: make(Monkeys21)}
}

func (d *Day21) Title() string {
	return "Day 21 - Monkey Math"
}

func (d *Day21) Setup() {
	// var lines = lib.ReadLines("data/21-test.txt")
	var lines = lib.ReadLines("data/21-data.txt")

	// matches: root: pppw + sjmn
	var opMatcher = regexp.MustCompile(`(\w+): (\w+) ([+*/-]) (\w+)`)
	var nrMatcher = regexp.MustCompile(`(\w+): (\d+)`)

	for _, line := range lines {
		var opGroup = opMatcher.FindStringSubmatch(line)
		var nrGroup = nrMatcher.FindStringSubmatch(line)
		if len(opGroup) == 5 {
			var name = opGroup[1]
			var op1 = Operation21{
				op:      rune(opGroup[3][0]),
				left:    opGroup[2],
				right:   opGroup[4],
				monkeys: &d.monkeys,
			}
			var op2 = Operation21{
				op:      rune(opGroup[3][0]),
				left:    opGroup[2],
				right:   opGroup[4],
				monkeys: &d.monkeys2,
			}
			d.monkeys[name] = &op1
			d.monkeys2[name] = &op2

		} else if len(nrGroup) == 3 {
			var name = nrGroup[1]
			nr, err := strconv.ParseInt(nrGroup[2], 10, 64)
			lib.Check(err)
			if name == "humn" {
				var val1 = HumnValue21{
					nr: int64(nr),
				}
				var val2 = HumnValue21{
					nr: int64(nr),
				}
				d.monkeys[name] = &val1
				d.monkeys2[name] = &val2

			} else {
				var val1 = Value21{
					nr: int64(nr),
				}
				var val2 = Value21{
					nr: int64(nr),
				}
				d.monkeys[name] = &val1
				d.monkeys2[name] = &val2
			}
		}
	}
	// fmt.Printf("%#v\n", d.monkeys)
}

func (d *Day21) SolveProblem1() {
	var rootMonkey = d.monkeys["root"]

	res, ok := rootMonkey.value()
	if ok != true {
		panic("Oops, something went completely wrong ...")
	}
	d.s1 = res
}

func (d *Day21) SolveProblem2() {
	var rootMonkey = d.monkeys2["root"].(*Operation21)

	// test: in which branch does "humn" appear?
	var leftMonkey = d.monkeys2[rootMonkey.left]
	var rightMonkey = d.monkeys2[rootMonkey.right]
	var leftNeedsHumn = leftMonkey.needsMonkey("humn")
	var rightNeedsNumn = rightMonkey.needsMonkey("humn")

	// calc the side that can be calculated:
	var res int64 = 0
	var ok bool
	var humnMonkey = d.monkeys2["humn"].(*HumnValue21)
	var fixedMonkey Expression21
	var guessMonkey Expression21
	if leftNeedsHumn == true {
		guessMonkey = leftMonkey
		fixedMonkey = rightMonkey
	} else if rightNeedsNumn == true {
		guessMonkey = rightMonkey
		fixedMonkey = leftMonkey
	} else {
		panic("both monkeys need humn, not possible")
	}
	res, ok = fixedMonkey.value()
	if ok != true {
		panic("Oops, this should not happen: this side should not return a false result.")
	}

	// Use a Bisect to find the correct solution:
	// after some fiddling, I noticed that
	// the solution DEcreases, while the input for humn INcreases.
	// So we start guessing between a boundary, and use a
	// bisect (binary search) to find the correct solution:

	// NOTE that this ONLY works with MY real data! It does not even work with the
	// test data.....
	var upperBound int64 = 1000000000000000
	var lowerBound int64 = 0
	var humnValue int64 = 0

	for {
		humnMonkey.setValue(humnValue)
		guess, ok := guessMonkey.value()
		if guess > res {
			if ok != true {
				humnValue += 1
				continue
			}
			lowerBound = humnValue
		} else if guess < res {
			if ok != true {
				humnValue -= 1
				continue
			}
			upperBound = humnValue
		} else if guess == res {
			break
		}
		if upperBound == lowerBound {
			fmt.Println("no solution found")
			break
		}
		humnValue = (upperBound + lowerBound) / 2
	}

	d.s2 = humnValue
}

func (d *Day21) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day21) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
