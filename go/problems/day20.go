package problems

import (
	"fmt"
	"strconv"

	"alexi.ch/aoc/2022/lib"
)

type Number20 struct {
	nr       int64
	orig_pos int64
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
	for i, line := range lines {
		if len(line) > 0 {
			nr, err := strconv.ParseInt(line, 10, 64)
			lib.Check(err)
			d.numbers = append(d.numbers, &Number20{nr: nr, orig_pos: int64(i)})
			d.numbers2 = append(d.numbers2, &Number20{nr: nr * 811589153, orig_pos: int64(i)})
		}
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day20) SolveProblem1() {

	/**
	 * Ring buffer problem:
	 * I take a simple array as ring buffer.
	 * Moving an element is done the following way:
	 * - remove the element from the array
	 * - calculate the new insert position. This can easily be done using mod() as ring wrapper function
	 * - insert removed element at the new position
	 *
	 * This is done by slicing / re-creating the array. Maybe I will redo it using a double-linked list.
	 */

	var length = int64(len(d.numbers))
	// fmt.Printf("Length: %d\n", length)
	var insert_idx int64 = 0

	for i := int64(0); i < length; i++ {
		// find number by searching the orig_pos:
		var start_idx = d.findIndexOfOrigPos(i, &d.numbers)
		var element = d.numbers[start_idx]

		// fmt.Printf("before move: %v\n", d.numbers)
		// fmt.Printf("moving Nr: %d\n", element.nr)

		if element.nr == 0 {
			continue
		}

		// remove actual element from list:
		var tmp = make([]*Number20, 0)
		tmp = append(tmp, d.numbers[:start_idx]...)
		tmp = append(tmp, d.numbers[start_idx+1:]...)
		// fmt.Printf("after remove: %v\n", tmp)

		// calc insert pos:
		if element.nr > 0 {
			insert_idx = (start_idx + element.nr - 1) % (length - 1)
			// insert AFTER insert_idx:
			var tmp2 = make([]*Number20, 0)
			tmp2 = append(tmp2, tmp[:insert_idx+1]...)
			tmp2 = append(tmp2, element)
			tmp2 = append(tmp2, tmp[insert_idx+1:]...)
			tmp = tmp2
		} else {
			// --> (len + (start + nr) % len) % 6
			var l = length - 1
			insert_idx = (l + (start_idx+element.nr-1)%l) % l
			// insert AFTER insert_idx:
			var tmp2 = make([]*Number20, 0)
			tmp2 = append(tmp2, tmp[:insert_idx+1]...)
			tmp2 = append(tmp2, element)
			tmp2 = append(tmp2, tmp[insert_idx+1:]...)
			tmp = tmp2
		}

		d.numbers = tmp
		// fmt.Printf("after move: %v\n\n", d.numbers)
	}

	// fmt.Printf("after move: %v\n\n", d.numbers)
	var zero_idx = d.findIndexOfNr(0, &d.numbers)
	var nr_1000 = d.numbers[(zero_idx+1000)%int64(len(d.numbers))].nr
	var nr_2000 = d.numbers[(zero_idx+2000)%int64(len(d.numbers))].nr
	var nr_3000 = d.numbers[(zero_idx+3000)%int64(len(d.numbers))].nr
	fmt.Printf("1000: %d, 2000: %d, 3000: %d\n", nr_1000, nr_2000, nr_3000)

	d.s1 = nr_1000 + nr_2000 + nr_3000
}

func (d *Day20) SolveProblem2() {
	/**
	 * Same procedure as in 1, but with larger numbers and more rotations.
	 * Because part 1 is already very optimized, this one is a breeze!
	 */
	var length = int64(len(d.numbers2))
	// fmt.Printf("Length: %d\n", length)
	var insert_idx int64 = 0

	for r := 0; r < 10; r++ {
		for i := int64(0); i < length; i++ {
			// find number by searching the orig_pos:
			var start_idx = d.findIndexOfOrigPos(i, &d.numbers2)
			var element = d.numbers2[start_idx]

			// fmt.Printf("before move: %v\n", d.numbers2)
			// fmt.Printf("moving Nr: %d\n", element.nr)

			if element.nr == 0 {
				continue
			}

			// remove actual element from list:
			var tmp = make([]*Number20, 0)
			tmp = append(tmp, d.numbers2[:start_idx]...)
			tmp = append(tmp, d.numbers2[start_idx+1:]...)
			// fmt.Printf("after remove: %v\n", tmp)

			// calc insert pos:
			if element.nr > 0 {
				insert_idx = (start_idx + element.nr - 1) % (length - 1)
				// insert AFTER insert_idx:
				var tmp2 = make([]*Number20, 0)
				tmp2 = append(tmp2, tmp[:insert_idx+1]...)
				tmp2 = append(tmp2, element)
				tmp2 = append(tmp2, tmp[insert_idx+1:]...)
				tmp = tmp2
			} else {
				// --> (len + (start + nr) % len) % 6
				var l = length - 1
				insert_idx = (l + (start_idx+element.nr-1)%l) % l
				// insert AFTER insert_idx:
				var tmp2 = make([]*Number20, 0)
				tmp2 = append(tmp2, tmp[:insert_idx+1]...)
				tmp2 = append(tmp2, element)
				tmp2 = append(tmp2, tmp[insert_idx+1:]...)
				tmp = tmp2
			}

			d.numbers2 = tmp
			// fmt.Printf("after move: %v\n\n", d.numbers)
		}
	}

	// fmt.Printf("after move: %v\n\n", d.numbers2)
	var zero_idx = d.findIndexOfNr(0, &d.numbers2)
	var nr_1000 = d.numbers2[(zero_idx+1000)%length].nr
	var nr_2000 = d.numbers2[(zero_idx+2000)%length].nr
	var nr_3000 = d.numbers2[(zero_idx+3000)%length].nr
	fmt.Printf("1000: %d, 2000: %d, 3000: %d\n", nr_1000, nr_2000, nr_3000)

	d.s2 = nr_1000 + nr_2000 + nr_3000
}

func (d *Day20) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day20) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
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
