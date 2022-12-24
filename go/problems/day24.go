package problems

import (
	"fmt"
	"sort"

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
	x int
	y int
}

type PlaceHashCombo struct {
	place Coord24
	hash  string
}

type Day24 struct {
	s1                  int
	s2                  int
	startmap            mapdata
	map_per_minute      map[int]*mapdata
	best_solution       int
	map_hash_per_minute map[int]string
	place_hashes        map[PlaceHashCombo]int
}

func NewDay24() Day24 {
	return Day24{
		s1:                  0,
		s2:                  0,
		startmap:            make(mapdata, 0),
		map_per_minute:      make(map[int]*mapdata),
		map_hash_per_minute: make(map[int]string),
		place_hashes:        make(map[PlaceHashCombo]int),
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

	// calc all maps and hashes per minute
	// var actmap = d.nextBlizzardMap(&d.startmap)
	d.map_per_minute[0] = &d.startmap
	d.map_hash_per_minute[0] = d.hashMap(&d.startmap)

	var minute = 1
	var seen_hashes = make([]string, 1)
	seen_hashes[0] = d.map_hash_per_minute[0]

	for i := 0; i < len(d.startmap)*len(d.startmap[0]); i++ {
		var actmap = d.nextBlizzardMap(d.map_per_minute[minute-1])
		var hash = d.hashMap(&actmap)

		if lib.Contains(seen_hashes, hash) {
			break
		}
		seen_hashes = append(seen_hashes, hash)
		d.map_per_minute[minute] = &actmap
		d.map_hash_per_minute[minute] = hash
		minute += 1
	}

	fmt.Printf("hashes: %d\n", len(d.map_hash_per_minute))
	d.printMap(&d.startmap)
}

func (d *Day24) SolveProblem1() {
	// var rounds = 0
	// var nextmap = d.startmap

	d.best_solution = 0

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
	// d.stored_maps[0] = &d.startmap

	fmt.Printf("Start at: %v\n", start)
	fmt.Printf("End at: :%v\n", end)

	// depth-First-Search with a recursive algorithm:
	// start with the start position in minute 0.
	//
	// check function: (takes actual minute)
	// check if I occupy an empty space. If not, return -1 (not a valid route)
	// if yes, collect all possible next steps.
	// if one of the next steps is the exit, return the number of steps (minutes) and end the process
	// call check function for each valid step with minute+1

	// 1st round, solution 1
	var sum = 0
	var minute = 1
	var s1 = 0
	// wait until I can move out of the start position:
	for {
		var tmpmap = d.map_per_minute[minute%len(d.map_hash_per_minute)]
		if len((*tmpmap)[start.y+1][start.x]) == 0 {
			break
		}
		minute += 1
	}
	s1 = d.checkMinute(minute, Coord24{x: start.x, y: start.y + 1}, &start, &end)
	sum += s1

	// 2nd round: back to the start:
	fmt.Println("2nd round")
	d.best_solution = 0
	d.place_hashes = make(map[PlaceHashCombo]int)
	var new_start = Coord24{x: end.x, y: end.y}
	var new_end = Coord24{x: start.x, y: start.y}
	var s2 = 0
	minute = s1 + 1
	// wait until I can move out of the start position:
	for {
		var tmpmap = d.map_per_minute[minute%len(d.map_hash_per_minute)]
		if len((*tmpmap)[new_start.y-1][new_start.x]) == 0 {
			break
		}
		fmt.Println("Wait for start...")
		minute += 1
	}
	s2 = d.checkMinute(minute, Coord24{x: new_start.x, y: new_start.y - 1}, &new_start, &new_end)
	sum += s2
	fmt.Printf("Back home: %d\n", s2)

	// 3rd round: to the end again!
	fmt.Println("3rd round")
	d.best_solution = 0
	d.place_hashes = make(map[PlaceHashCombo]int)
	var s3 = 0
	minute = s2 + 1
	// wait until I can move out of the start position:
	for {
		var tmpmap = d.map_per_minute[minute%len(d.map_hash_per_minute)]
		if len((*tmpmap)[start.y+1][start.x]) == 0 {
			break
		}
		fmt.Println("Wait for start...")
		minute += 1
	}
	s3 = d.checkMinute(minute, Coord24{x: start.x, y: start.y + 1}, &start, &end)
	sum += s3
	fmt.Printf("Back home: %d\n", s3)

	// for i := 1; i <= rounds; i++ {
	// 	nextmap = d.nextBlizzardMap(&nextmap)
	// 	fmt.Printf("Round: %d\n", i)
	// 	d.printMap(&nextmap)
	// }
	d.s1 = s1
	d.s2 = s3
}

func (d *Day24) SolveProblem2() {
	// d.s2 = 0
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

func (d *Day24) checkMinute(minute int, pos Coord24, start *Coord24, target *Coord24) int {
	if d.best_solution > 0 && d.best_solution <= minute {
		// there is already a better solution
		return -1
	}
	// invalid location (out of bounds):
	if pos.x < 0 || pos.y < 0 || pos.y >= len(d.startmap) || pos.x >= len(d.startmap[0]) {
		return -1
	}
	// invalid location (start pos):
	// if pos == *start {
	// 	return -2
	// }

	// fmt.Printf("Minute %d: Working on loc: %v\n", minute, pos)

	// get (cached) map for actual minute:
	var minute_mod = len(d.map_hash_per_minute)
	var actmap = d.map_per_minute[minute%minute_mod]

	// if we were here already, in a room with the same hash, we're in a loop: return
	var hash = d.map_hash_per_minute[minute%minute_mod]
	if d.place_hashes[PlaceHashCombo{place: pos, hash: hash}] > 0 && d.place_hashes[PlaceHashCombo{place: pos, hash: hash}] <= minute {
		return -1
	}
	d.place_hashes[PlaceHashCombo{place: pos, hash: hash}] = minute

	// am I in a blizzard, or in a wall?
	if len((*actmap)[pos.y][pos.x]) > 0 {
		// yes, so this is the wrong way
		return -1
	}

	// am I on the exit? Yay!
	if pos == *target {
		d.best_solution = minute
		return minute
	}

	// ok, clear grounds, go ahead and search forward
	// check each direction, and the "wait" direction recursively, if we find an exit:
	var valid_solutions = make([]int, 0)
	for _, dir := range dirs24 {
		var nextPos = Coord24{x: pos.x + dir.x, y: pos.y + dir.y}
		var res = d.checkMinute(minute+1, nextPos, start, target)
		if res > 0 {
			// return res
			// fmt.Printf("Found solution: %d\n", res)
			valid_solutions = append(valid_solutions, res)
		}
	}
	if len(valid_solutions) > 0 {
		sort.Ints(valid_solutions)
		return valid_solutions[0]
	}

	return -1
}

func (d *Day24) hashMap(m *mapdata) string {
	return fmt.Sprintf("%v", m)
}
