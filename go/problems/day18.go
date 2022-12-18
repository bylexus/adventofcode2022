package problems

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"

	"alexi.ch/aoc/2022/lib"
)

type Point18 struct {
	x int64
	y int64
	z int64
}

type material uint

const (
	M_Rock  material = 1
	M_Water material = 2
)

type Cube map[Point18]material

type Day18 struct {
	s1       uint64
	s2       uint64
	cube     Cube
	minCoord Point18
	maxCoord Point18
}

func NewDay18() Day18 {
	return Day18{
		s1:   0,
		s2:   0,
		cube: make(Cube, 0),
		minCoord: Point18{
			x: math.MaxInt64,
			y: math.MaxInt64,
			z: math.MaxInt64,
		},
		maxCoord: Point18{
			x: 0,
			y: 0,
			z: 0,
		},
	}
}

func (d *Day18) Title() string {
	return "Day 18 - Boiling Boulders"
}

func (d *Day18) Setup() {
	// var lines = lib.ReadLines("data/18-test.txt")
	var lines = lib.ReadLines("data/18-data.txt")
	for _, line := range lines {
		var parts = strings.Split(line, ",")
		if len(parts) == 3 {
			x, err := strconv.ParseInt(parts[0], 10, 64)
			lib.Check(err)
			y, err := strconv.ParseInt(parts[1], 10, 64)
			lib.Check(err)
			z, err := strconv.ParseInt(parts[2], 10, 64)
			lib.Check(err)
			d.cube[Point18{x: x, y: y, z: z}] = M_Rock

			// ------------- povray box output ------------------
			// fmt.Printf("box { <%d, %d, %d>, <%d, %d, %d> scale 0.999 }\n", x, y, z, x+1, y+1, z+1)
			// ------------- end povray box output --------------

			if x < d.minCoord.x {
				d.minCoord.x = x
			}
			if y < d.minCoord.y {
				d.minCoord.y = y
			}
			if z < d.minCoord.z {
				d.minCoord.z = z
			}
			if x > d.maxCoord.x {
				d.maxCoord.x = x
			}
			if y > d.maxCoord.y {
				d.maxCoord.y = y
			}
			if z > d.maxCoord.z {
				d.maxCoord.z = z
			}
		}
	}
	// fmt.Printf("Cube dimension: min: %v, max: %v\n", d.minCoord, d.maxCoord)
	// fmt.Printf("%v\n", d.cube)
}

func (d *Day18) SolveProblem1() {
	var surfaces uint64 = 0
	var dirs = []Point18{
		{x: 1, y: 0, z: 0},
		{x: -1, y: 0, z: 0},
		{x: 0, y: 1, z: 0},
		{x: 0, y: -1, z: 0},
		{x: 0, y: 0, z: 1},
		{x: 0, y: 0, z: -1},
	}
	for p := range d.cube {
		for _, dt := range dirs {
			if d.cube[Point18{
				x: p.x + dt.x,
				y: p.y + dt.y,
				z: p.z + dt.z,
			}] == 0 {
				surfaces += 1
			}
		}
	}
	d.s1 = surfaces
}

func (d *Day18) SolveProblem2() {
	// flood-fill the cube!
	// we need to flood-fill the outside volumina, to cover all surfaces in water.
	// I'm using a breath-first algo with a stack for this.

	// starting at one coordinate outside the cube to fill water
	var start = &Point18{
		x: d.minCoord.x - 1,
		y: d.minCoord.y - 1,
		z: d.minCoord.z - 1,
	}
	// flood fill:
	d.fill(start)

	// count surfaces that touches water:
	var surfaces uint64 = 0
	var dirs = []Point18{
		{x: 1, y: 0, z: 0},
		{x: -1, y: 0, z: 0},
		{x: 0, y: 1, z: 0},
		{x: 0, y: -1, z: 0},
		{x: 0, y: 0, z: 1},
		{x: 0, y: 0, z: -1},
	}
	for p := range d.cube {
		for _, dt := range dirs {
			if d.cube[p] == M_Rock && d.cube[Point18{
				x: p.x + dt.x,
				y: p.y + dt.y,
				z: p.z + dt.z,
			}] == M_Water {
				surfaces += 1
			}
		}
	}

	d.s2 = surfaces
}

func (d *Day18) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day18) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day18) fill(p *Point18) {
	var dirs = []Point18{
		{x: 1, y: 0, z: 0},
		{x: -1, y: 0, z: 0},
		{x: 0, y: 1, z: 0},
		{x: 0, y: -1, z: 0},
		{x: 0, y: 0, z: 1},
		{x: 0, y: 0, z: -1},
	}

	var todo = list.New()
	todo.PushBack(p)
	for act_i := todo.Front(); act_i != nil; act_i = act_i.Next() {
		var act = act_i.Value.(*Point18)
		todo.Remove(act_i)
		if d.cube[*act] == 0 {
			d.cube[*act] = M_Water
			for _, dt := range dirs {
				var nextPoint = Point18{
					x: act.x + dt.x,
					y: act.y + dt.y,
					z: act.z + dt.z,
				}
				// limit to outer hull:
				if nextPoint.x >= d.minCoord.x-1 && nextPoint.y >= d.minCoord.y-1 && nextPoint.z >= d.minCoord.z-1 &&
					nextPoint.x <= d.maxCoord.x+1 && nextPoint.y <= d.maxCoord.y+1 && nextPoint.z <= d.maxCoord.z+1 {
					if d.cube[nextPoint] == 0 {
						todo.PushBack(&nextPoint)
					}
				}
			}
		}
	}
}
