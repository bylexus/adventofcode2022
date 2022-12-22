package problems

import (
	"fmt"
	"regexp"
	"strconv"

	"alexi.ch/aoc/2022/lib"
)

type instr_type int

const (
	I_NR    instr_type = 1
	I_INSTR instr_type = 2
)

type Instruction struct {
	itype instr_type
	nr    int
	dir   rune
}

type Point22 struct {
	x int
	y int
}

type Day22 struct {
	s1           int
	s2           uint64
	mmap         map[Point22]rune
	max_coord    Point22
	instructions []Instruction
}

func NewDay22() Day22 {
	return Day22{s1: 0, s2: 0, mmap: make(map[Point22]rune), instructions: make([]Instruction, 0)}
}

func (d *Day22) Title() string {
	return "Day 22 - Monkey Map"
}

func (d *Day22) Setup() {
	// var lines = lib.ReadLines("data/22-test.txt")
	var lines = lib.ReadLines("data/22-data.txt")
	var idx = 0
	d.max_coord = Point22{x: 0, y: 0}

	// read the maze lines, until the empy line,
	// count also max line width:
	for y, line := range lines {
		if len(line) > 0 {
			for x, c := range line {
				var pt = Point22{
					x: x, y: y,
				}

				d.mmap[pt] = c

				if x > d.max_coord.x {
					d.max_coord.x = x
				}
			}
			idx = y
		} else {
			break
		}
	}
	d.max_coord.y = idx

	// read the instructions
	var matcher = regexp.MustCompile(`(\d+|\w)`)
	var instructions = lines[idx+2]
	var groups = matcher.FindAllStringSubmatch(instructions, -1)
	for _, g := range groups {
		nr, err := strconv.ParseInt(g[0], 10, 32)
		if err != nil {
			d.instructions = append(d.instructions, Instruction{itype: I_INSTR, dir: rune(g[0][0])})
		} else {
			d.instructions = append(d.instructions, Instruction{itype: I_NR, nr: int(nr)})
		}
	}

	// fmt.Printf("%#v\n", d.instructions)
	// fmt.Printf("%#v\n", d.mmap)
	// fmt.Printf("%#v\n", d.max_coord)
	// d.printMap()
}

func (d *Day22) SolveProblem1() {
	// find start pos: 1st pos on top line that is available:
	var actPos = Point22{x: 0, y: 0}
	var dirVec = Point22{x: 1, y: 0}
	for x := 0; x < d.max_coord.x; x++ {
		if d.mmap[Point22{x: x, y: 0}] == '.' {
			actPos.x = x
			break
		}
	}
	fmt.Printf("Start at: x:%d, y:%d\n", actPos.x, actPos.y)
	for _, instr := range d.instructions {
		if instr.itype == I_NR {
			actPos = d.walk(actPos, instr.nr, dirVec)
		} else if instr.itype == I_INSTR {
			dirVec = d.turn(dirVec, instr.dir)
		}
		// d.mmap[actPos] = '@'
		// d.printMap()
	}
	// d.printMap()
	var dirNr = 0
	if dirVec == (Point22{x: 1, y: 0}) {
		dirNr = 0
	}
	if dirVec == (Point22{x: 0, y: 1}) {
		dirNr = 1
	}
	if dirVec == (Point22{x: -1, y: 0}) {
		dirNr = 2
	}
	if dirVec == (Point22{x: 0, y: -1}) {
		dirNr = 3
	}
	d.s1 = 1000*(actPos.y+1) + 4*(actPos.x+1) + dirNr
}

func (d *Day22) SolveProblem2() {
	d.s2 = 0
}

func (d *Day22) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day22) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day22) printMap() {
	for y := 0; y < d.max_coord.y; y++ {
		for x := 0; x < d.max_coord.x; x++ {
			fmt.Printf("%c", d.mmap[Point22{x: x, y: y}])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Day22) walk(actPos Point22, nr int, dirVec Point22) Point22 {
	for i := 0; i < nr; i++ {
		// mark current dir on map, just for testing:
		//------------------------------
		if dirVec.x == 1 {
			d.mmap[actPos] = '>'
		}
		if dirVec.x == -1 {
			d.mmap[actPos] = '<'
		}
		if dirVec.y == 1 {
			d.mmap[actPos] = 'v'
		}
		if dirVec.y == -1 {
			d.mmap[actPos] = '^'
		}
		//------------------------------
		var nextPos = d.calcNextPos(actPos, dirVec)
		var floor = d.mmap[nextPos]

		if floor == ' ' || floor == 0 {
			fmt.Printf("act pos: x: %d, y: %d, dir: %v\n", actPos.x, actPos.y, dirVec)
			fmt.Printf("next pos: x: %d, y: %d, dir: %v\n", nextPos.x, nextPos.y, dirVec)
			fmt.Printf("steps: %d\n", nr)
			fmt.Printf("floor: %d, %c\n", floor, floor)
			panic("Oops! walked off the map! Should not happen!")
		}
		if floor == '#' {
			// hit wall, stop
			break
		} else {
			// ok, go ahead:
			actPos = nextPos
		}
	}
	return actPos
}

func (d *Day22) calcNextPos(actPos Point22, dirVec Point22) Point22 {
	var nextPos = Point22{x: actPos.x + dirVec.x, y: actPos.y + dirVec.y}

	if d.mmap[nextPos] == 0 || d.mmap[nextPos] == ' ' {
		// map ends, wrap around: turn direction and walk back until the other border:
		dirVec.x *= -1
		dirVec.y *= -1
		actPos = nextPos
		for {
			nextPos = Point22{x: actPos.x + dirVec.x, y: actPos.y + dirVec.y}
			if d.mmap[nextPos] == 0 || d.mmap[nextPos] == ' ' {
				return actPos
			}
			actPos = nextPos
		}
	} else {
		// next pos still on the map
		return nextPos
	}
}

func (d *Day22) turn(dirVec Point22, dir rune) Point22 {
	if dir == 'L' {
		// (x,y) --> '(y, -x)
		return Point22{x: dirVec.y, y: dirVec.x * -1}
	}
	if dir == 'R' {
		// (x, y) ---> '(-y, x)
		return Point22{x: dirVec.y * -1, y: dirVec.x}
	}
	panic("Unknown direction")
}
