package problems

import (
	"fmt"
	"regexp"
	"strconv"

	"alexi.ch/aoc/2022/lib"
)

type Price struct {
	material string
	price    int64
}

type Prices map[string][]Price
type Blueprint struct {
	prices     Prices
	robots     map[string]int64
	new_robots map[string]int64
	material   map[string]int64
}

type Day19 struct {
	s1         uint64
	s2         uint64
	blueprints []Blueprint
}

func NewDay19() Day19 {
	return Day19{s1: 0, s2: 0, blueprints: make([]Blueprint, 0)}
}

func (d *Day19) Title() string {
	return "Day 19 - Not Enough Minerals"
}

func (d *Day19) Setup() {
	var lines = lib.ReadLines("data/19-test.txt")
	// var lines = lib.ReadLines("data/19-data.txt")

	// matcher for:
	// Blueprint 2: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 5 clay. Each geode robot costs 3 ore and 7 obsidian.
	var matcher = regexp.MustCompile(`(Each (\w+) robot costs (\d+) (\w+)( and (\d+) (\w+))?.)+`)
	for _, line := range lines {
		var blueprint = Blueprint{
			prices:     make(Prices),
			robots:     make(map[string]int64),
			new_robots: make(map[string]int64),
			material:   make(map[string]int64),
		}
		// each blueprint start with 1 ore robot:
		blueprint.robots["ore"] = 1

		var groups = matcher.FindAllStringSubmatch(line, -1)
		for _, g := range groups {
			// fmt.Printf("%#v\n", g)
			var robot = g[2]

			// price 1
			amt1, err := strconv.ParseInt(g[3], 10, 64)
			lib.Check(err)
			var price1 = Price{
				material: g[4],
				price:    amt1,
			}
			blueprint.prices[robot] = append(blueprint.prices[robot], price1)
			if len(g[6]) > 0 {
				// price 2
				amt2, err := strconv.ParseInt(g[6], 10, 64)
				lib.Check(err)
				var price2 = Price{
					material: g[7],
					price:    amt2,
				}
				blueprint.prices[robot] = append(blueprint.prices[robot], price2)
			}

		}
		d.blueprints = append(d.blueprints, blueprint)
	}
	// fmt.Printf("%#v\n", d.blueprints)
}

func (d *Day19) SolveProblem1() {
	// keep track of what I have and what I need
	// The "geode" defines the requirement.
	// Put the materials I need for a geode in a list.
	// can I build one? If not,
	// check what we need to build the missing elements.
	// put that in the list, too.
	// as soon as we CAN build something from our list,
	// build it.
	var maxMinutes = 24

	for i, b := range d.blueprints {
		for minute := 0; minute < maxMinutes; minute++ {

			// can I build a robot? Note that we can only build 1 per minute
			for wantRobot, p := range b.prices {
				var enough = true
				// do I have enough material?
				for _, price := range p {
					if b.material[price.material] < price.price {
						enough = false
						break
					}
				}
				if enough {
					for _, price := range p {
						b.material[price.material] -= price.price
					}
					b.new_robots[wantRobot] += 1
					// done, cannot build one more this round:
					break
				}
			}

			// collect materials:
			for material, nr := range b.robots {
				b.material[material] += nr
			}

			// did we create new robots before ? add them:
			for r, nr := range b.new_robots {
				b.robots[r] += nr
				b.new_robots[r] = 0
			}
			fmt.Printf("Blueprint %d: After Minute %d: Robots: %v, material: %v\n", i+1, minute+1, b.robots, b.material)
		}
	}

	d.s1 = 0
}

func (d *Day19) SolveProblem2() {
	d.s2 = 0
}

func (d *Day19) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day19) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
