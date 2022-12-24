package problems

import (
	"fmt"

	"alexi.ch/aoc/2022/lib"
)

var dirs24 = []Coord24{
	{x: 1, y: 0},
	{x: 0, y: 1},
	{x: -1, y: 0},
	{x: 0, y: -1},
	// wait direction: don't move
	{x: 0, y: 0},
}

type mapentry []rune
type mapdata [][]mapentry

type Coord24 struct {
	x      int
	y      int
	minute int
}

type Day24 struct {
	s1          int
	s2          int
	startmap    mapdata
	stored_maps map[int]*mapdata
}

func NewDay24() Day24 {
	return Day24{
		s1:          0,
		s2:          0,
		startmap:    make(mapdata, 0),
		stored_maps: make(map[int]*mapdata),
	}
}

func (d *Day24) Title() string {
	return "Day 24 - Blizzard Basin"
}

func (d *Day24) Setup() {
	// var lines = lib.ReadLines("data/24-test.txt")
	var lines = lib.ReadLines("data/24-data.txt")
	for _, line := range lines {
		if len(line) > 0 {
			var entries = make([]mapentry, 0)
			for _, r := range line {
				var entry = make([]rune, 0)
				if r != '.' {
					entry = append(entry, r)
				}
				entries = append(entries, entry)
			}
			d.startmap = append(d.startmap, entries)
		}
	}
	d.printMap(&d.startmap)
}

func (d *Day24) SolveProblem1() {
	// var rounds = 0
	// var nextmap = d.startmap

	var start = Coord24{x: 0, y: 0}
	var end = Coord24{x: 0, y: len(d.startmap) - 1}

	// find start point:
	for x := 0; x < len(d.startmap[start.y]); x++ {
		if len(d.startmap[start.y][x]) == 0 {
			start.x = x
		}
	}
	// find end point:
	for x := 0; x < len(d.startmap[end.y]); x++ {
		if len(d.startmap[end.y][x]) == 0 {
			end.x = x
		}
	}

	// store first map:
	d.stored_maps[0] = &d.startmap

	fmt.Printf("Start at: %v\n", start)
	fmt.Printf("End at: :%v\n", end)

	// breath first search with a queue:
	// fill the queue with the 1st position in minute 1 (down)
	//
	// as long as the queue is not empty:
	// take the front item
	// check: is it at the end position? ok, got it, done!
	//
	// check if it occupies an empty space in that minute. If not, skip (not a valid location)
	// if yes, collect all possible next steps, put them at the end of the queue with minute + 1

	var queue = make([]Coord24, 1)
	queue[0] = Coord24{x: start.x, y: start.y + 1, minute: 1}
	for {
		if len(queue) == 0 {
			break
		}
		var act = &queue[0]
		queue = queue[1:]
		// found the first one that reaches the exit:
		if act.x == end.x && act.y == end.y {
			d.s1 = act.minute
			break
		}
		// fmt.Printf("queue len: %d\n", len(queue))
		queue = d.appendNextStepsToQueue(act, queue, &start, &end)
	}

	// for i := 1; i <= rounds; i++ {
	// 	nextmap = d.nextBlizzardMap(&nextmap)
	// 	fmt.Printf("Round: %d\n", i)
	// 	d.printMap(&nextmap)
	// }
	// d.s1 = 0
}

func (d *Day24) SolveProblem2() {
	d.s2 = 0
}

func (d *Day24) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day24) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day24) printMap(data *mapdata) {
	for y := 0; y < len(*data); y++ {
		for x := 0; x < len((*data)[y]); x++ {
			var entry = (*data)[y][x]
			if len(entry) == 0 {
				fmt.Print(".")
			} else if len(entry) == 1 {
				fmt.Printf("%c", entry[0])
			} else {
				fmt.Printf("%d", len(entry))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Day24) nextBlizzardMap(act *mapdata) mapdata {
	var next = make(mapdata, 0)

	// create new, empty map:
	for _, line := range *act {
		var entries = make([]mapentry, 0)
		for range line {
			var entry = make([]rune, 0)
			entries = append(entries, entry)
		}
		next = append(next, entries)
	}

	// calculate new map by considering each location from the act map:
	for y, line := range *act {
		for x, entry := range line {
			for _, obj := range entry {
				switch obj {
				case '#':
					// just a wall: put it in the same place as before:
					next[y][x] = append(next[y][x], '#')
				case '<':
					// left-ward wind
					var x1 = x - 1
					if len((*act)[y][x1]) > 0 && (*act)[y][x1][0] == '#' {
						x1 = len((*act)[y]) - 2
					}
					next[y][x1] = append(next[y][x1], obj)
				case '>':
					// right-ward wind
					var x1 = x + 1
					if len((*act)[y][x1]) > 0 && (*act)[y][x1][0] == '#' {
						x1 = 1
					}
					next[y][x1] = append(next[y][x1], obj)
				case '^':
					// up-ward wind
					var y1 = y - 1
					if len((*act)[y1][x]) > 0 && (*act)[y1][x][0] == '#' {
						y1 = len(*act) - 2
					}
					next[y1][x] = append(next[y1][x], obj)
				case 'v':
					// down-ward wind
					var y1 = y + 1
					if len((*act)[y1][x]) > 0 && (*act)[y1][x][0] == '#' {
						y1 = 1
					}
					next[y1][x] = append(next[y1][x], obj)
				}
			}
		}
	}

	return next
}

func (d *Day24) appendNextStepsToQueue(pos *Coord24, queue []Coord24, start *Coord24, target *Coord24) []Coord24 {
	// pre-cache map for actual minute:
	var actmap = d.stored_maps[pos.minute]
	if actmap == nil {
		var newmap = d.nextBlizzardMap(d.stored_maps[pos.minute-1])
		actmap = &newmap
		d.stored_maps[pos.minute] = actmap
	}

	// pre-cache map for NEXT minute:
	var nextmap = d.stored_maps[pos.minute+1]
	if nextmap == nil {
		var newmap = d.nextBlizzardMap(d.stored_maps[pos.minute])
		nextmap = &newmap
		d.stored_maps[pos.minute+1] = nextmap
	}

	var actMinute = pos.minute

	// invalid location (out of bounds):
	if pos.x < 0 || pos.y < 0 || pos.y >= len(d.startmap) || pos.x >= len(d.startmap[0]) {
		return queue
	}
	// invalid location (start pos):
	if pos.x == start.x && pos.y == start.y {
		return queue
	}
	// fmt.Printf("Minute %d: Working on loc: %v\n", minute, pos)

	// am I in a blizzard, or in a wall?
	if len((*actmap)[pos.y][pos.x]) > 0 {
		// yes, so this is the wrong way
		return queue
	}

	// // am I on the exit? Yay!
	// if *pos == *target {
	// 	pos.minute = actMinute
	// 	return queue
	// }

	// ok, clear grounds, go ahead and fill the next locations
	// check each direction, and the "wait" direction recursively, if we find an exit:
	for _, dir := range dirs24 {
		var loc = Coord24{
			x:      pos.x + dir.x,
			y:      pos.y + dir.y,
			minute: actMinute + 1,
		}
		// skip start:
		if loc.x == start.x && loc.y == start.y {
			continue
		}
		// invalid location (out of bounds):
		if loc.x < 0 || loc.y < 0 || loc.y >= len(d.startmap) || loc.x >= len(d.startmap[0]) {
			continue
		}
		// skip wall:
		if len(d.startmap[loc.y][loc.x]) > 0 && d.startmap[loc.y][loc.x][0] == '#' {
			continue
		}
		// skip blizzard location in NEXT map:
		if len((*nextmap)[loc.y][loc.x]) > 0 {
			continue
		}

		queue = append(queue, loc)
	}
	return queue
	// var valid_solutions = make([]int, 0)
	// if len(valid_solutions) > 0 {
	// 	sort.Ints(valid_solutions)
	// 	return valid_solutions[0]
	// }

}
