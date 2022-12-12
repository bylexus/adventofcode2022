package problems

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"alexi.ch/aoc/2022/lib"
)

type Monkey struct {
	items         []uint64
	inspect_times uint64
	op            Operator
	val_type      ValueType
	op_value      uint64
	div_test      uint64
	on_true       uint64
	on_false      uint64
}

type Operator string
type ValueType string

const (
	OpPlus Operator = "+"
	OpMult Operator = "*"
)

const (
	ValueTypeNumber = "number"
	ValueTypeOld    = "old"
)

type Day11 struct {
	s1 uint64
	s2 uint64

	monkeys1 []Monkey
	monkeys2 []Monkey
}

func NewDay11() Day11 {
	return Day11{s1: 0, s2: 0, monkeys1: make([]Monkey, 0), monkeys2: make([]Monkey, 0)}
}

func (d *Day11) Title() string {
	return "Day 11 - Monkey in the Middle"
}

func (d *Day11) Setup() {
	// var lines = lib.ReadLines("data/11-test.txt")
	var lines = lib.ReadLines("data/11-data.txt")

	var starting_items_re = regexp.MustCompile(`Starting items: (.*)`)
	var op_re = regexp.MustCompile(`Operation: new = old (.) (old|\d+)`)
	var test_re = regexp.MustCompile(`Test: divisible by (\d+)`)
	var true_re = regexp.MustCompile(`If true: throw to monkey (\d+)`)
	var false_re = regexp.MustCompile(`If false: throw to monkey (\d+)`)

	var line = 0
	for {
		if line >= len(lines) {
			break
		}

		// skip first line:
		line += 1

		// starting items:
		var si_group = starting_items_re.FindStringSubmatch(lines[line])
		var splitters = strings.Split(si_group[1], ",")
		var values = lib.Map(&splitters, func(item string) uint64 {
			var nr, err = strconv.ParseUint(strings.TrimSpace(item), 10, 64)
			lib.Check(err)
			return nr
		})
		fmt.Printf("parsed values: %v\n", values)
		var items1 = values
		var items2 = values
		line += 1

		// op:
		var op_group = op_re.FindStringSubmatch(lines[line])
		var op_str = strings.TrimSpace(op_group[2])
		var op Operator
		var val_type ValueType
		var op_value uint64 = 0
		var err error

		if op_group[1] == "*" {
			op = OpMult
		} else if op_group[1] == "+" {
			op = OpPlus
		} else {
			panic("Unknown operator")
		}

		if op_str == "old" {
			val_type = ValueTypeOld
		} else {
			val_type = ValueTypeNumber
			op_value, err = strconv.ParseUint(op_str, 10, 64)
			lib.Check(err)
		}
		line += 1

		// test:
		var test_group = test_re.FindStringSubmatch(lines[line])
		var div_by uint64 = 0
		div_by, err = strconv.ParseUint(test_group[1], 10, 64)
		lib.Check(err)
		line += 1

		// if true:
		var true_group = true_re.FindStringSubmatch(lines[line])
		var true_monkey uint64
		true_monkey, err = strconv.ParseUint(true_group[1], 10, 64)
		lib.Check(err)
		line += 1

		// if false:
		var false_group = false_re.FindStringSubmatch(lines[line])
		var false_monkey uint64
		false_monkey, err = strconv.ParseUint(false_group[1], 10, 64)
		lib.Check(err)
		line += 1

		// end, skip line
		line += 1

		var monkey1 = Monkey{
			items:         items1,
			inspect_times: 0,
			op:            op,
			op_value:      op_value,
			val_type:      val_type,
			div_test:      div_by,
			on_true:       true_monkey,
			on_false:      false_monkey,
		}
		var monkey2 = Monkey{
			items:         items2,
			inspect_times: 0,
			op:            op,
			op_value:      op_value,
			val_type:      val_type,
			div_test:      div_by,
			on_true:       true_monkey,
			on_false:      false_monkey,
		}
		d.monkeys1 = append(d.monkeys1, monkey1)
		d.monkeys2 = append(d.monkeys2, monkey2)
	}
	d.s1 = 0
	d.s2 = 0
	// fmt.Printf("%v\n", d.monkeys1)
}

func (d *Day11) SolveProblem1() {
	for round := 0; round < 20; round++ {
		for i := range d.monkeys1 {
			var monkey = &d.monkeys1[i]
			for _, item := range monkey.items {
				monkey.inspect_times += 1
				var new_item uint64 = 0
				if monkey.op == OpPlus {
					if monkey.val_type == ValueTypeNumber {
						new_item = item + monkey.op_value
					} else if monkey.val_type == ValueTypeOld {
						new_item = item + item
					}
				} else if monkey.op == OpMult {
					if monkey.val_type == ValueTypeNumber {
						new_item = item * monkey.op_value
					} else if monkey.val_type == ValueTypeOld {
						new_item = item * item
					}
				}
				new_item /= 3
				if new_item%monkey.div_test == 0 {
					d.monkeys1[monkey.on_true].items = append(d.monkeys1[monkey.on_true].items, new_item)
				} else {
					d.monkeys1[monkey.on_false].items = append(d.monkeys1[monkey.on_false].items, new_item)
				}
			}
			monkey.items = make([]uint64, 0)
		}
	}

	var inspections = lib.Map(&d.monkeys1, func(m Monkey) uint64 { return m.inspect_times })
	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})

	d.s1 = inspections[0] * inspections[1]
}

func (d *Day11) SolveProblem2() {
	var divisors = lib.Map(&d.monkeys2, func(m Monkey) uint64 {
		return m.div_test
	})
	var divisor uint64 = 1
	for _, d := range divisors {
		divisor *= d
	}
	for round := 0; round < 10000; round++ {
		for i := range d.monkeys2 {
			var monkey = &d.monkeys2[i]
			for _, item := range monkey.items {
				monkey.inspect_times += 1
				var new_item uint64 = 0
				if monkey.op == OpPlus {
					if monkey.val_type == ValueTypeNumber {
						new_item = item + monkey.op_value
					} else if monkey.val_type == ValueTypeOld {
						new_item = item + item
					}
				} else if monkey.op == OpMult {
					if monkey.val_type == ValueTypeNumber {
						new_item = item * monkey.op_value
					} else if monkey.val_type == ValueTypeOld {
						new_item = item * item
					}
				}
				new_item %= divisor
				if new_item%monkey.div_test == 0 {
					d.monkeys2[monkey.on_true].items = append(d.monkeys2[monkey.on_true].items, new_item)
				} else {
					d.monkeys2[monkey.on_false].items = append(d.monkeys2[monkey.on_false].items, new_item)
				}
			}
			monkey.items = make([]uint64, 0)
		}
	}

	var inspections = lib.Map(&d.monkeys2, func(m Monkey) uint64 { return m.inspect_times })
	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})

	d.s2 = inspections[0] * inspections[1]
}

func (d *Day11) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day11) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
