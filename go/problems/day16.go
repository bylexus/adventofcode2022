package problems

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"alexi.ch/aoc/2022/lib"
)

type Room16 struct {
	name         string
	visited      bool
	flow         int64
	paths        []string
	enter_minute int64
	parent       *Room16
	total_flow   int64
}

func (r *Room16) String() string {
	return fmt.Sprintf("Room: %s, Flow: %d, Leads to: %s",
		r.name,
		r.flow,
		strings.Join(r.paths, ", "),
	)
}

type Day16 struct {
	s1   int64
	s2   int64
	cave map[string]*Room16
}

func NewDay16() Day16 {
	return Day16{s1: 0, s2: 0, cave: make(map[string]*Room16)}
}

func (d *Day16) Title() string {
	return "Day 16 - Proboscidea Volcanium"
}

func (d *Day16) Setup() {
	var lines = lib.ReadLines("data/16-test.txt")
	// var lines = lib.ReadLines("data/16-data.txt")

	// matcher for:
	// Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
	var matcher = regexp.MustCompile(`Valve (\w+) has flow rate=(-?\d+); tunnels? leads? to valves? (.*)`)
	for _, line := range lines {
		var group = matcher.FindStringSubmatch(line)
		if len(group) == 4 {
			var name = strings.TrimSpace(group[1])
			flow, err := strconv.ParseInt(strings.TrimSpace(group[2]), 10, 64)
			lib.Check(err)
			var paths = strings.Split(strings.TrimSpace(group[3]), ", ")
			d.cave[name] = &Room16{
				name:    name,
				visited: false,
				flow:    flow,
				paths:   paths,
			}
		}
	}
	fmt.Printf("%#v\n", d.cave)
	for _, c := range d.cave {
		fmt.Printf("%s\n", c)
	}
}

func (d *Day16) SolveProblem1() {
	// TODO: No solution yet...
	return
	/*
		idea:
		  I need to create a tree of possible paths: From each node, there are
		  many possible sub-tree paths to walk.
		  Each path need to be walked until 30minues are over.
		  Each node knows how many "points" it will generate (= how many valve minutes to count).
		  The path with the most points wins.

		  Problem is that we cannot calculate the whole tree: we must examine each node after processing,
		  and only keep the sub-tree with the most points until proceeding.
	*/
	// we start with a new node AA, with an empty path
	var tree_of_possibilities = Room16{
		name:         "AA",
		paths:        make([]string, 0),
		enter_minute: 1,
		parent:       nil,
	}
	d.buildSubtree(&tree_of_possibilities)

	d.s1 = tree_of_possibilities.total_flow
}

func (d *Day16) SolveProblem2() {
	d.s2 = 0
}

func (d *Day16) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day16) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day16) buildSubtree(room *Room16) {
	// if we reach minute 30, we're done.
	if room.enter_minute >= 30 {
		d.printChain(room)
		return
	}
	// don't process node if it is visited too much (more than it has entrances)
	var parent_count = d.countSameParent(room)
	// if parent_count > len(d.cave[room.name].paths) {
	// 	return
	// }
	// fmt.Printf("Enter minute: %d, Room: %s\n", room.enter_minute, room.name)
	var next_minute = room.enter_minute + 1

	// do we possibly need to open a valve?
	// this is the case if:
	// - our own flow value is > 0
	// - or room was not already visited in the tree above
	if room.flow > 0 && parent_count == 0 {
		room.total_flow = room.flow * (30 - room.enter_minute)
		// opening a valve costs one minute
		next_minute += 1
	}

	// process all possible paths, keep the one that outputs the most:
	var possible_paths = d.cave[room.name].paths
	var biggest_room *Room16 = nil
	for _, path := range possible_paths {

		var new_room = Room16{
			name:         path,
			flow:         d.cave[path].flow,
			parent:       room,
			enter_minute: next_minute,
			total_flow:   0,
		}
		d.buildSubtree(&new_room)
		if biggest_room == nil || new_room.total_flow > biggest_room.total_flow {
			biggest_room = &new_room
		}
	}

	// now we have the biggest room: add it to my total flow:
	if biggest_room != nil {
		room.total_flow += biggest_room.total_flow
	}
}

func (d *Day16) countSameParent(room *Room16) int {
	if room == nil {
		return 0
	}
	var parent = room.parent
	var count = 0
	for {
		if parent == nil {
			return count
		}
		if parent.name == room.name {
			count += 1
		}
		parent = parent.parent
	}
}

func (d *Day16) printChain(room *Room16) {
	for ; room != nil; room = room.parent {
		fmt.Printf("%s <- ", room.name)
	}
	fmt.Println()
}
