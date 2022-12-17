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
	s1          int64
	s2          int64
	jet         string
	jet_pos     int64
	pattern_idx int64
	patterns    []Stone
	cave        map[Point17]rune
	width       int64
	height      int64
	// end_pos_hashes map[string]int64
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

		// end_pos_hashes: make(map[string]int64),
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
	 */

	var stone_start = Point17{
		x: 0, y: 0,
	}

	// y pos: this is the BOTTOM of the rock
	var stone_pos = Point17{x: 0, y: 0}

	// var max_stones int64 = 1
	var max_stones int64 = 2022
	// var max_stones int64 = 100000
	// var max_stones int64 = 1_000_000_000_000

	var moved_fall bool = true
	d.height = 0

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
	}

	d.s1 = d.height
}

func (d *Day17) SolveProblem2() {
	d.s2 = 0
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
