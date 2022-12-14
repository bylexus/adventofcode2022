package problems

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"alexi.ch/aoc/2022/lib"
)

type Entry struct {
	value     uint
	content   []Entry
	is_list   bool
	is_marker bool
}

func (d Entry) String() string {
	if d.is_list {
		var content = strings.Join(lib.Map(&d.content, func(e Entry) string {
			return e.String()
		}), ",")
		return fmt.Sprintf("[%s]", content)
	} else {
		return fmt.Sprintf("%d", d.value)
	}
}

type Day13 struct {
	s1      uint
	s2      uint
	entries []Entry
}

func NewDay13() Day13 {
	return Day13{s1: 0, s2: 0, entries: make([]Entry, 0)}
}

func (d *Day13) Title() string {
	return "Day 13 - Distress Signal"
}

func (d *Day13) Setup() {
	// var lines = lib.ReadLines("data/13-test.txt")
	var lines = lib.ReadLines("data/13-data.txt")
	for i := 0; i < len(lines); i += 3 {
		var left = d.parse(lines[i])[0]
		var right = d.parse(lines[i+1])[0]
		d.entries = append(d.entries, left)
		d.entries = append(d.entries, right)
		// fmt.Printf("%v\n", left)
		// fmt.Printf("%v\n", right)
	}
}

/**
 * compares the two entries, returns:
 * 1 if the entries are in order
 * 0 if no decision could be made
 * -1 if the entries are not in order
 */
func (d *Day13) compare(left *Entry, right *Entry) int {
	// both are numbers: {
	if !left.is_list && !right.is_list {
		if left.value < right.value {
			return 1
		}
		if left.value > right.value {
			return -1
		}
		return 0
	}

	// one list, one integer
	if left.is_list && !right.is_list {
		var list = Entry{
			is_list: true,
			content: make([]Entry, 0),
		}
		list.content = append(list.content, *right)
		return d.compare(left, &list)
	}
	// one list, one integer
	if !left.is_list && right.is_list {
		var list = Entry{
			is_list: true,
			content: make([]Entry, 0),
		}
		list.content = append(list.content, *left)
		return d.compare(&list, right)
	}

	// both are lists:
	if left.is_list && right.is_list {
		for i := 0; i < len(left.content); i++ {
			// not in order if right side has no comparison element:
			if len(right.content) < i+1 {
				return -1
			}
			var cmp = d.compare(&left.content[i], &right.content[i])
			if cmp == 1 {
				return 1
			}
			if cmp == -1 {
				return -1
			}
		}
		if len(left.content) < len(right.content) {
			return 1
		}
		return 0
	}

	panic("should not reach this line")
}

func (d *Day13) SolveProblem1() {
	var sum uint = 0
	for i := 0; i < len(d.entries); i += 2 {
		var left = &d.entries[i]
		var right = &d.entries[i+1]
		if d.compare(left, right) == 1 {
			sum += uint(i/2 + 1)
		}
	}
	d.s1 = sum
}

func (d *Day13) SolveProblem2() {
	var m1 = d.parse("[[2]]")[0]
	var m2 = d.parse("[[6]]")[0]
	m1.is_marker = true
	m2.is_marker = true

	var all = d.entries
	all = append(all, m1)
	all = append(all, m2)

	sort.Slice(all, func(i, j int) bool {
		return d.compare(&all[i], &all[j]) == 1
	})

	var p = 1
	for i, m := range all {
		if m.is_marker {
			p = p * (i + 1)
		}
	}
	d.s2 = uint(p)
}

func (d *Day13) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day13) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day13) parse(line string) []Entry {
	// Modes:
	// 0: start
	// 1: list parsing
	// 2: number parsing
	var mode = 0
	var list_start = 0
	var brace_count = 0
	var result = make([]Entry, 0)
	var nr_str = ""

	for i, c := range line {
		if mode == 0 {
			if c == '[' {
				mode = 1
				list_start = i + 1
				brace_count = 1
				continue
			}
			if c == ',' {
				continue
			}
			if c >= '0' && c <= '9' {
				mode = 2
				nr_str = fmt.Sprintf("%c", c)
				continue
			}
			panic(fmt.Sprintf("Unknown value in line: %c\n", c))
		}
		if mode == 1 {
			if c == '[' {
				brace_count += 1
			}
			if c == ']' {
				brace_count -= 1
				if brace_count == 0 {
					// full list read, extract content
					var sub = d.parse(line[list_start:i])
					var entry = Entry{
						is_list: true,
						content: sub,
					}
					result = append(result, entry)
					mode = 0
				}
			}
		}
		if mode == 2 {
			if c >= '0' && c <= '9' {
				nr_str = nr_str + fmt.Sprintf("%c", c)
			} else {
				nr, err := strconv.ParseUint(nr_str, 10, 32)
				lib.Check(err)
				var entry = Entry{
					is_list: false,
					value:   uint(nr),
				}
				result = append(result, entry)
				mode = 0
			}
		}
	}

	// take last number, if we were in the process of parsing one:
	if mode == 2 {
		nr, err := strconv.ParseUint(nr_str, 10, 32)
		lib.Check(err)
		var entry = Entry{
			is_list: false,
			value:   uint(nr),
		}
		result = append(result, entry)
	}

	return result
}
