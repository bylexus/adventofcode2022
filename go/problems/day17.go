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

	var stone_start = Point17{
		x: 0, y: 0,
	}

	// y pos: this is the BOTTOM of the rock
	var stone_pos = Point17{x: 0, y: 0}

	var max_stones = 2022
	// var max_stones = 1_000_000_000_000

	var moved_fall bool = true
	var moved_jet bool = false
	d.height = 0

	for s := 1; s <= max_stones; s++ {
		// create stone: calc start position
		// which stone?
		var stone = d.patterns[d.pattern_idx]
		stone_start = Point17{x: 2, y: d.height + 3}
		// fmt.Printf("start pos of stone: %d, height: %d\n", stone_start.y, d.height)

		stone_pos = stone_start
		moved_fall = true

		// move it, until it reaches something:
		for {
			if moved_fall == false {
				// fmt.Printf("Ground collision at x: %d, y: %d\n", stone_pos.x, stone_pos.y)
				break
			}

			// jet movement:
			// fmt.Printf("before jet: stone x: %d, y: %d\n", stone_pos.x, stone_pos.y)
			stone_pos, moved_jet = d.jet_movement(stone_pos, &stone)
			// fmt.Printf("after jet: stone x: %d, y: %d\n", stone_pos.x, stone_pos.y)
			if moved_jet == false {
				// fmt.Println("  no jet movement happened")
			}

			// fall down:
			stone_pos, moved_fall = d.fall_down_movement(stone_pos, &stone)
			if moved_fall == false {
				// fmt.Printf("  no stone movement happened: %d, %d\n", stone_pos.x, stone_pos.y)
			} else {
				// fmt.Printf("  stone falled one to x: %d, y: %d\n", stone_pos.x, stone_pos.y)
			}

			// // if the next fall would cause a collission, stop:
			// if d.collides(&stone, Point17{x: stone_pos.x, y: stone_pos.y - 1}) {
			// 	fmt.Printf("Ground collision at x: %d, y: %d\n", stone_pos.x, stone_pos.y)
			// 	break
			// }
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

	// TODO: Colission detection
	// fmt.Printf("Jet consideration: %d: %c\n", d.jet_pos, dir)
	if dir == '<' && startPos.x > 0 {
		// fmt.Printf("  Jet movement: <\n")
		endPos.x -= 1
		if d.collides(stone, endPos) {
			// fmt.Printf("  Jet collieds, no movement: <\n")
			endPos.x += 1
		} else {
			moved = true
		}
	} else if dir == '>' && (startPos.x+int64(len((*stone)[0]))) < d.width {
		// fmt.Printf("  Jet movement: >\n")
		endPos.x += 1
		if d.collides(stone, endPos) {
			// fmt.Printf("  Jet collieds, no movement: >\n")
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

	// TODO: colission detection
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
				// fmt.Printf("Collission: %c with %c\n", (*stone)[y][x], d.cave[Point17{y: pos.y + (h - 1 - y), x: pos.x + x}])
				return true
			}
		}
	}

	return false
}

func (d *Day17) printCave() {
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
