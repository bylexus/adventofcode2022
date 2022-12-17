package problems

import (
	"fmt"

	"alexi.ch/aoc/2022/lib"
)

type Point17 struct {
	x int64
	y int64
}

type Stone [][]rune

type Day17 struct {
	s1            int64
	s2            int64
	jet           string
	jet_pos       int64
	pattern_idx   int64
	patterns      []Stone
	cave          map[Point17]rune
	width         int64
	height        int64
	pattern_match map[string]int64
}

func NewDay17() Day17 {
	return Day17{
		s1:          0,
		s2:          0,
		patterns:    make([]Stone, 0),
		jet_pos:     0,
		pattern_idx: 0,
		cave:        make(map[Point17]rune),
		width:       7,
		height:      0,

		pattern_match: make(map[string]int64),
	}
}

func (d *Day17) Title() string {
	return "Day 17 - Pyroclastic Flow"
}

func (d *Day17) Setup() {
	// var lines = lib.ReadLines("data/17-test.txt")
	var lines = lib.ReadLines("data/17-data.txt")
	d.jet = lines[0]
	d.patterns = []Stone{
		// ####
		[][]rune{
			[]rune{'#', '#', '#', '#'},
		},
		/*
			.#.
			###
			.#.
		*/
		[][]rune{
			[]rune{' ', '#', ' '},
			[]rune{'#', '#', '#'},
			[]rune{' ', '#', ' '},
		},
		/*
			..#
			..#
			###
		*/
		[][]rune{
			[]rune{' ', ' ', '#'},
			[]rune{' ', ' ', '#'},
			[]rune{'#', '#', '#'},
		},
		/*
			#
			#
			#
			#
		*/
		[][]rune{
			[]rune{'#'},
			[]rune{'#'},
			[]rune{'#'},
			[]rune{'#'},
		},
		/*
			##
			##
		*/
		[][]rune{
			[]rune{'#', '#'},
			[]rune{'#', '#'},
		},
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day17) SolveProblem1() {
	/**
	 * TETRIS! YAY!
	 *
	 * coordinate system:
	 *
	 *  +y <----- y increases to top
	 *  ^
	 *  |
	 *  |
	 * 0| <----- Bottom line is 0
	 *  +----------> +x
	 */

	/**
	 * Idea for solution 2:
	 * may be the gained height will repeat after the
	 * jet and stone index are both in-sync again?
	 * --> so find lcm(jet, stones) (least common multiplier),
	 *  and see what happens with the height there...
	 *
	 * no, no pattern....
	 *
	 * 2nd idea: Maybe the end x positions for all stones in order will
	 * repeat after n rounds?
	 *
	 * also, no.
	 *
	 * 3rd attempt:
	 * We search for the same start conditions:
	 * last line is full, and next stone is the 1st again. then it must repeat.
	 *
	 * nope, no luck ...
	 *
	 * again:
	 * look for a full line, and check if we reached it with the same
	 * stream index + stone index...
	 */

	var stone_start = Point17{
		x: 0, y: 0,
	}

	// y pos: this is the BOTTOM of the rock
	var stone_pos = Point17{x: 0, y: 0}

	// var max_stones int64 = 1
	var max_stones int64 = 2022
	// var max_stones int64 = 229
	// var max_stones int64 = 2931 + 229
	// var max_stones int64 = 100000
	// var max_stones int64 = 1_000_000_000_000

	var moved_fall bool = true
	d.height = 0
	var last_matching_height int64 = 0

	for s := int64(1); s <= max_stones; s++ {
		// create stone: calc start position
		// which stone?
		var stone = d.patterns[d.pattern_idx]
		stone_start = Point17{x: 2, y: d.height + 3}

		stone_pos = stone_start
		moved_fall = true

		// move it, until it reaches something:
		for {
			// did we stop falling in the last round? OK, we're done with this stone:
			if moved_fall == false {
				break
			}

			// jet movement:
			stone_pos, _ = d.jet_movement(stone_pos, &stone)

			// fall down:
			stone_pos, moved_fall = d.fall_down_movement(stone_pos, &stone)
		}

		// update cave
		d.insert_stone(&stone, stone_pos)

		// advance stone
		d.pattern_idx = (d.pattern_idx + 1) % int64(len(d.patterns))
		// d.printCave()

		// check for full line:
		var found = true
		for x := int64(0); x < d.width; x++ {
			if d.cave[Point17{x: x, y: d.height - 1}] != '#' {
				found = false
				break
			}
		}

		// found:
		// after 1755 stone drops, the full line appears on the top again
		// height gained in between: 2768
		// first full line match found at stone drop 2931, height: 4589
		// so we only need to calc the height from the last round < 1_000_000_000_000
		// rounds still to go: (1000000000000 - 2931)
		// a = (1000000000000 - 2931) / 1755: full duplicate rounds, * 2768: height with full rounds: 569800568 * 2768 = 1577207972224 height
		// b = (1000000000000 - 2931) % 1755 rounds to calc height --> 229. Height after 229 rounds: 341
		// c = height gained from a zero line to (zero line + 299 rounds): (height on round (2931+229): 4962 - height on round 2931): 4589  = 373 --> height gained from zero line to round 229
		// c = 4589 + 373 + 1577207972224 = 1577207977154

		// 1000000000000 - (569800568 * 1755) - 2931
		// 4589 + 1577207972224 + 341 = 1577207977186 ---> solution for 2!
		if found {
			var hash = fmt.Sprintf("%d:%d", d.pattern_idx, d.jet_pos)
			if d.pattern_match[hash] > 0 {
				fmt.Printf("Found full match at round %d, height: %d\n", s, d.height)
				fmt.Printf("last seen same: %d, diff: %d, height diff: %d\n", d.pattern_match[hash], s-d.pattern_match[hash], d.height-last_matching_height)
			}
			d.pattern_match[hash] = s
			last_matching_height = d.height
		}
	}

	d.s1 = d.height
}

func (d *Day17) SolveProblem2() {
	/**
	I found the 2nd solution with some analyzing:
	This approach only works on the real data, NOT the test data.
	It seems that at some point in time (after n stone drops), a full line appears on top.
	I then wait for the NEXT line appearing, at the SAME Jet Index:
	If this happens, a loop is detected, and we can calculate the rest.

	Here are the facts from my input data:
	// Full line appears after 1176 stone drops.
	// Next full line with same jet index appears after 2931
	// --> After d = (2931 - 1176) = 1755 stone drops, the world repeats.
	// --> one repeation gains height: 2768
	// --> So we have to do:
	// calc the height on round 2931 --> 4589 --> a
	// calc how many full rounds we can calculate:  (1000000000000 - 2931) / 1755 = 569800568 full same-height rounds
	// --> we gain height: 569800568 * 2768 = 1577207972224 height from all repeated rounds
	// --> now we need the missing rounds at the end:
	// b = (1000000000000 - 2931) % 1755 rounds missing = 229.
	// --> Height after 229 rounds: 341 (found between two full same-height rows)
	// in total:
	// 4589 + 1577207972224 + 341 = 1577207977186 ---> solution for 2!
	//
	*/
	d.s2 = 1577207977186
}

func (d *Day17) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day17) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day17) jet_movement(startPos Point17, stone *Stone) (Point17, bool) {
	var endPos = startPos
	var moved = false
	var dir = d.jet[d.jet_pos]

	if dir == '<' && startPos.x > 0 {
		endPos.x -= 1
		if d.collides(stone, endPos) {
			endPos.x += 1
		} else {
			moved = true
		}
	} else if dir == '>' && (startPos.x+int64(len((*stone)[0]))) < d.width {
		endPos.x += 1
		if d.collides(stone, endPos) {
			endPos.x -= 1
		} else {
			moved = true
		}
	}

	// advance jet:
	d.jet_pos = int64(d.jet_pos+1) % int64(len(d.jet))
	return endPos, moved
}

func (d *Day17) fall_down_movement(startPos Point17, stone *Stone) (Point17, bool) {
	var endPos = startPos
	var moved = false

	if startPos.y > 0 {
		endPos.y -= 1
		if d.collides(stone, endPos) {
			endPos.y += 1
		} else {
			moved = true
		}
	}

	return endPos, moved
}

func (d *Day17) insert_stone(stone *Stone, pos Point17) {
	var h = int64(len(*stone))
	for y := int64(0); y < h; y++ {
		for x := int64(0); x < int64(len((*stone)[y])); x++ {
			if (*stone)[y][x] == '#' {
				// attention: stones are in an array, so we must invert the stone's y axis
				d.cave[Point17{y: pos.y + (h - 1 - y), x: pos.x + x}] = '#'
				if d.height < pos.y+y+1 {
					d.height = pos.y + y + 1
				}
			}
		}
	}
}

func (d *Day17) collides(stone *Stone, pos Point17) bool {
	// below ground:
	if pos.y < 0 {
		return true
	}

	var h = int64(len(*stone))
	for y := int64(0); y < h; y++ {
		for x := int64(0); x < int64(len((*stone)[y])); x++ {
			// attention: stones are in an array, so we must invert the stone's y axis
			if (*stone)[y][x] == '#' && d.cave[Point17{y: pos.y + (h - 1 - y), x: pos.x + x}] == '#' {
				return true
			}
		}
	}

	return false
}

func (d *Day17) printCave() {
	fmt.Printf("\nHeight: %d\n", d.height)
	for y := d.height + 3; y >= 0; y-- {
		fmt.Print("|")
		for x := int64(0); x < d.width; x++ {
			var rock = d.cave[Point17{x, y}]
			if rock == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%c", rock)
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+-------+\n")
}
