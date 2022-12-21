package problems

import (
	"fmt"
	"strconv"

	"alexi.ch/aoc/2022/lib"
)

type Number20 struct {
	nr       int64
	orig_pos int64
	prev     *Number20
	next     *Number20
}

func (n Number20) String() string {
	return fmt.Sprintf("%d", n.nr)
}

type Day20 struct {
	s1       int64
	s2       int64
	numbers  []*Number20
	numbers2 []*Number20
}

func NewDay20() Day20 {
	return Day20{s1: 0, s2: 0, numbers: make([]*Number20, 0)}
}

func (d *Day20) Title() string {
	return "Day 20 - Grove Positioning System"
}

func (d *Day20) Setup() {
	// var lines = lib.ReadLines("data/20-test.txt")
	var lines = lib.ReadLines("data/20-data.txt")
	var prev1 *Number20 = nil
	var prev2 *Number20 = nil
	for i, line := range lines {
		if len(line) > 0 {
			nr, err := strconv.ParseInt(line, 10, 64)
			lib.Check(err)
			var nr1 = Number20{nr: nr, orig_pos: int64(i)}
			var nr2 = Number20{nr: nr * 811589153, orig_pos: int64(i)}
			if prev1 != nil {
				prev1.next = &nr1
				nr1.prev = prev1
			}
			if prev2 != nil {
				prev2.next = &nr2
				nr2.prev = prev2
			}

			prev1 = &nr1
			prev2 = &nr2

			d.numbers = append(d.numbers, &nr1)
			d.numbers2 = append(d.numbers2, &nr2)
		}
	}
	// form ring:
	d.numbers[0].prev = d.numbers[len(d.numbers)-1]
	d.numbers[len(d.numbers)-1].next = d.numbers[0]

	d.numbers2[0].prev = d.numbers2[len(d.numbers2)-1]
	d.numbers2[len(d.numbers2)-1].next = d.numbers2[0]

	// fmt.Println(d.ringStr(d.numbers[0]))
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day20) SolveProblem1() {

	/**
	 * Ring buffer problem:
	 *
	 * The actual solution takes a real ring buffer, formed with
	 * a double linked list.
	 *
	 * I just move the one element n mod list length, instead of the full length.
	 * This way I can save a loooot of movements.
	 */

	d.moveAllNumbersOnce(&d.numbers)
	var length int64 = int64(len(d.numbers))

	// fmt.Printf("after move: %v\n\n", d.numbers)
	var idx = d.findIndexOfNr(0, &d.numbers)
	var act = d.numbers[idx]
	for i := int64(0); i < 1000%length; i++ {
		act = act.next
	}
	var nr_1000 = act.nr
	for i := int64(0); i < 1000%length; i++ {
		act = act.next
	}
	var nr_2000 = act.nr
	for i := int64(0); i < 1000%length; i++ {
		act = act.next
	}
	var nr_3000 = act.nr
	// fmt.Printf("1000: %d, 2000: %d, 3000: %d\n", nr_1000, nr_2000, nr_3000)

	d.s1 = nr_1000 + nr_2000 + nr_3000
}

func (d *Day20) SolveProblem2() {
	/**
	 * Same procedure as in 1, but with larger numbers and more rotations.
	 * Because part 1 is already very optimized, this one is a breeze!
	 */
	var length = int64(len(d.numbers2))
	// fmt.Printf("Length: %d\n", length)

	for r := 0; r < 10; r++ {
		d.moveAllNumbersOnce(&d.numbers2)
	}

	// fmt.Printf("after move: %v\n\n", d.numbers)
	var idx = d.findIndexOfNr(0, &d.numbers2)
	var act = d.numbers2[idx]
	for i := int64(0); i < 1000%length; i++ {
		act = act.next
	}
	var nr_1000 = act.nr
	for i := int64(0); i < 1000%length; i++ {
		act = act.next
	}
	var nr_2000 = act.nr
	for i := int64(0); i < 1000%length; i++ {
		act = act.next
	}
	var nr_3000 = act.nr
	// fmt.Printf("1000: %d, 2000: %d, 3000: %d\n", nr_1000, nr_2000, nr_3000)

	d.s2 = nr_1000 + nr_2000 + nr_3000
}

func (d *Day20) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day20) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

/**
 * Ring buffer rotation function
 * Rotates all numbers in order once
 */
func (d *Day20) moveAllNumbersOnce(numbers *[]*Number20) {
	var length = int64(len(*numbers))

	for i := int64(0); i < length; i++ {
		// find number by searching the orig_pos:
		var start_idx = d.findIndexOfOrigPos(i, numbers)
		var element = (*numbers)[start_idx]

		if element.nr == 0 {
			continue
		}

		// remove actual element from list:
		element.prev.next = element.next
		element.next.prev = element.prev

		// calc insert pos:
		var act = element

		if element.nr > 0 {
			act = act.prev
			for s := int64(0); s < (element.nr % (length - 1)); s++ {
				act = act.next
			}
			// insert element again:
			element.prev = act
			element.next = act.next
			act.next.prev = element
			act.next = element
		} else {
			act = act.next
			for s := int64(0); s > (element.nr % (length - 1)); s-- {
				act = act.prev
			}
			// insert element again:
			element.next = act
			element.prev = act.prev
			act.prev.next = element
			act.prev = element
		}
	}
}

func (d *Day20) findIndexOfOrigPos(pos int64, numbers *[]*Number20) int64 {
	for i := int64(0); i < int64(len(*numbers)); i++ {
		if (*numbers)[i].orig_pos == pos {
			return i
		}
	}
	panic("nr not found at original pos - cannot be!")
}

func (d *Day20) findIndexOfNr(nr int64, numbers *[]*Number20) int64 {
	for i := int64(0); i < int64(len(*numbers)); i++ {
		if (*numbers)[i].nr == nr {
			return i
		}
	}
	panic("nr not found at original pos - cannot be!")
}

func (d *Day20) ringStr(start *Number20) string {
	var out = ""
	var act = start
	for {
		out += fmt.Sprintf("%v ", act.nr)
		act = act.next
		if act == start {
			break
		}
	}
	return out
}
