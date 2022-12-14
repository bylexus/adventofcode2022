package problems

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"alexi.ch/aoc/2022/lib"
)

type PointDay14 struct {
	x int64
	y int64
}

type PlaceDay14 struct {
	material rune
	coord    *PointDay14
}

type Day14 struct {
	s1       int64
	s2       int64
	cave     map[PointDay14]*PlaceDay14
	caveMinX int64
	caveMaxX int64
	caveMinY int64
	caveMaxY int64
	start    PointDay14
}

func NewDay14() Day14 {
	return Day14{
		s1:       0,
		s2:       0,
		cave:     make(map[PointDay14]*PlaceDay14),
		caveMaxX: 0,
		caveMaxY: 0,
		caveMinX: math.MaxInt64,
		caveMinY: 0,
		start:    PointDay14{x: 500, y: 0},
	}
}

func (d *Day14) Title() string {
	return "Day 14 - Regolith Reservoir"
}

func (d *Day14) Setup() {
	// var lines = lib.ReadLines("data/14-test.txt")
	var lines = lib.ReadLines("data/14-data.txt")
	var coord_matcher = regexp.MustCompile(`((\d+),(\d+))`)
	for _, line := range lines {
		// parse strings to coords
		var groups = coord_matcher.FindAllStringSubmatch(line, -1)
		var coords = make([]PointDay14, 0)
		for _, group := range groups {
			// fmt.Printf("group: %#v\n", group)
			if len(group) == 4 {
				x, err := strconv.ParseInt(group[2], 10, 64)
				lib.Check(err)
				y, err := strconv.ParseInt(group[3], 10, 64)
				lib.Check(err)
				coords = append(coords, PointDay14{x: x, y: y})
			}
		}

		// draw cave
		var start = coords[0]
		for _, coord := range coords[1:] {
			var minY = lib.Min(start.y, coord.y)
			var maxY = lib.Max(start.y, coord.y)
			var minX = lib.Min(start.x, coord.x)
			var maxX = lib.Max(start.x, coord.x)
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					if x > d.caveMaxX {
						d.caveMaxX = x
					}
					if x < d.caveMinX {
						d.caveMinX = x
					}
					if y > d.caveMaxY {
						d.caveMaxY = y
					}
					var coord = PointDay14{x: x, y: y}
					d.cave[coord] = &PlaceDay14{
						material: '#',
						coord:    &coord,
					}
				}
			}
			start = coord
		}
	}
	// d.printCave()
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day14) SolveProblem1() {
	var counter int64 = 0
	for {
		// d.printCave()
		if d.drizzle(d.start) == true {
			counter += 1
		} else {
			break
		}
	}

	d.s1 = counter
}

func (d *Day14) SolveProblem2() {
	d.resetCave()
	var counter int64 = 0
	for {
		if d.drizzle2(d.start) == true {
			counter += 1
		} else {
			break
		}
	}

	// d.printCave()
	d.s2 = counter + 1
}

func (d *Day14) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day14) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

/**
 * let a single sand rock drizzle down,
 * until it is placed on the ground (return true)
 * or it falls into the abyss (return false)
 */
func (d *Day14) drizzle(start PointDay14) bool {
	var target PointDay14

	for {
		// check if we reached the bottom of the cave:
		if start.y > d.caveMaxY {
			return false
		}
		// can it be placed directly below?
		target = PointDay14{x: start.x, y: start.y + 1}
		if d.getCaveMaterial(target) == ' ' {
			start = target
			continue
		}
		// can it be placed to the left?
		target = PointDay14{x: start.x - 1, y: start.y + 1}
		if d.getCaveMaterial(target) == ' ' {
			start = target
			continue
		}
		// can it be placed to the right?
		target = PointDay14{x: start.x + 1, y: start.y + 1}
		if d.getCaveMaterial(target) == ' ' {
			start = target
			continue
		}
		d.cave[start] = &PlaceDay14{material: 'o', coord: &start}
		return true
	}
}

/**
 * let a single sand rock drizzle down,
 * until it is placed on the ground (return true)
 * or it reaches the entrypoint (start):
 * an infinite plane at yMax+2 defines the pit's ground
 */
func (d *Day14) drizzle2(start PointDay14) bool {
	var target PointDay14

	for {
		// can it be placed directly below?
		target = PointDay14{x: start.x, y: start.y + 1}
		if d.getCaveMaterial(target) == ' ' {
			start = target
			continue
		}
		// can it be placed to the left?
		target = PointDay14{x: start.x - 1, y: start.y + 1}
		if d.getCaveMaterial(target) == ' ' {
			start = target
			continue
		}
		// can it be placed to the right?
		target = PointDay14{x: start.x + 1, y: start.y + 1}
		if d.getCaveMaterial(target) == ' ' {
			start = target
			continue
		}

		// check if we reached the start point:
		if start == d.start {
			return false
		}

		// else, drop it!
		d.cave[start] = &PlaceDay14{material: 'o', coord: &start}
		return true
	}
}

func (d *Day14) printCave() {
	fmt.Println()
	for y := d.caveMinY; y <= d.caveMaxY; y++ {
		for x := d.caveMinX; x <= d.caveMaxX; x++ {
			var place = d.cave[PointDay14{x: int64(x), y: int64(y)}]
			if place != nil {
				fmt.Printf("%c", place.material)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Day14) resetCave() {
	for k, v := range d.cave {
		if v.material != '#' {
			delete(d.cave, k)
		}
	}
}

func (d *Day14) getCaveMaterial(coord PointDay14) rune {
	// endless horizontal bottom rocks:
	if coord.y == d.caveMaxY+2 {
		return '#'
	}
	place := d.cave[coord]
	if place == nil {
		return ' '
	}
	return place.material
}
