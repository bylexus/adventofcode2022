package problems

import (
	"fmt"
	"math"

	"alexi.ch/aoc/2022/lib"
)

var dirs23 = []Point23{
	// N
	{x: 0, y: -1},
	// NE
	{x: 1, y: -1},
	// E
	{x: 1, y: 0},
	// SE
	{x: 1, y: 1},
	// S
	{x: 0, y: 1},
	// SW
	{x: -1, y: 1},
	// W
	{x: -1, y: 0},
	// NW
	{x: -1, y: -1},
}

var dirOrder23 = []Point23{
	{x: 0, y: -1},
	{x: 0, y: 1},
	{x: -1, y: 0},
	{x: 1, y: 0},
}

type Point23 struct {
	x int
	y int
}

type Elve23 struct {
	proposedLocation *Point23
}

type ElvesMap map[Point23]*Elve23

type Day23 struct {
	s1          int
	s2          int
	elvesmap    ElvesMap
	proposalMap map[Point23]int
	tl          Point23
	br          Point23
}

func NewDay23() Day23 {
	return Day23{
		s1:          0,
		s2:          0,
		elvesmap:    make(ElvesMap),
		proposalMap: make(map[Point23]int),
		tl:          Point23{x: math.MaxInt, y: math.MaxInt},
		br:          Point23{x: math.MinInt, y: math.MinInt},
	}
}

func (d *Day23) Title() string {
	return "Day 23 - Unstable Diffusion"
}

func (d *Day23) Setup() {
	// var lines = lib.ReadLines("data/23-test.txt")
	// var lines = lib.ReadLines("data/23-test2.txt")
	var lines = lib.ReadLines("data/23-data.txt")
	for y, line := range lines {
		if len(line) > 0 {
			for x, r := range line {
				if r == '#' {
					var elve = Elve23{
						// location: Point23{x: x, y: y},
					}
					d.elvesmap[Point23{x: x, y: y}] = &elve
					d.updateMapSize(Point23{x: x, y: y})
				}
			}
		}
	}
	// d.printMap()
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day23) SolveProblem1() {
	var rounds = 10
	for r := 1; r <= rounds; r++ {
		d.processRound()
		// fmt.Printf("Round: %d\n", r)
		// d.printMap()
	}

	// calculate min bounding box:
	d.tl = Point23{x: math.MaxInt, y: math.MaxInt}
	d.br = Point23{x: math.MinInt, y: math.MinInt}
	for point := range d.elvesmap {
		d.updateMapSize(point)
	}
	// fmt.Printf("Final map:\n")
	// d.printMap()

	// count empty grounds:
	var counter = 0
	for y := d.tl.y; y <= d.br.y; y++ {
		for x := d.tl.x; x <= d.br.x; x++ {
			if d.elvesmap[Point23{x: x, y: y}] == nil {
				counter++
			}
		}
	}
	d.s1 = counter
}

func (d *Day23) SolveProblem2() {
	// we already have done 10 rounds in solution 1, so start with that:
	var rounds = 10
	for {
		rounds++
		if d.processRound() == true {
			break
		}
		// d.printMap()
		// time.Sleep(80 * time.Millisecond)
	}

	d.s2 = rounds
}

func (d *Day23) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day23) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day23) printMap() {
	fmt.Printf("Map size: x: %d-%d, y: %d-%d\n", d.tl.x, d.br.x, d.tl.y, d.br.y)
	for y := d.tl.y; y <= d.br.y; y++ {
		for x := d.tl.x; x <= d.br.x; x++ {
			var elve = d.elvesmap[Point23{x: x, y: y}]
			if elve != nil {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Day23) isFreeAround(point Point23) bool {
	for _, dir := range dirs23 {
		var checkPoint = Point23{x: point.x + dir.x, y: point.y + dir.y}
		if d.elvesmap[checkPoint] != nil {
			return false
		}
	}
	return true
}

func (d *Day23) isFreeAhead(point Point23, dir Point23) bool {
	for i := -1; i <= 1; i++ {
		var lookTo Point23
		if dir.x == 0 {
			lookTo = Point23{x: point.x + i, y: point.y + dir.y}
		} else {
			lookTo = Point23{x: point.x + dir.x, y: point.y + i}
		}
		if d.elvesmap[lookTo] != nil {
			return false
		}
	}
	return true
}

func (d *Day23) updateMapSize(point Point23) {
	if point.x < d.tl.x {
		d.tl.x = point.x
	}
	if point.y < d.tl.y {
		d.tl.y = point.y
	}
	if point.x > d.br.x {
		d.br.x = point.x
	}
	if point.y > d.br.y {
		d.br.y = point.y
	}
}

func (d *Day23) processRound() bool {
	// reset proposals:
	d.proposalMap = make(map[Point23]int)
	var noElveMoved = true

	// first half: Elfes proposals
	for point, elve := range d.elvesmap {
		if d.isFreeAround(point) {
			// if no one is around, stay here
			continue
		} else {
			// else, propose a new location:
			elve.proposedLocation = nil
			for _, dir := range dirOrder23 {
				if d.isFreeAhead(point, dir) {
					var proposalPoint = Point23{x: point.x + dir.x, y: point.y + dir.y}
					d.proposalMap[proposalPoint] += 1
					elve.proposedLocation = &proposalPoint
					break
				}
			}
		}
	}

	// 2nd half: check if the elve's proposals ara acceptable (only move if an elve is the only one proposing a new dir)
	for point, elve := range d.elvesmap {
		if elve.proposedLocation != nil && d.proposalMap[*elve.proposedLocation] == 1 {
			// ok, our elve is the only one that wants to move there - so move!
			delete(d.elvesmap, point)
			d.elvesmap[*elve.proposedLocation] = elve
			d.updateMapSize(*elve.proposedLocation)
			noElveMoved = false
		}
		elve.proposedLocation = nil
	}

	// re-order direction proposal array:
	var first = dirOrder23[0]
	dirOrder23 = append(dirOrder23, first)
	dirOrder23 = dirOrder23[1:]

	return noElveMoved
}
