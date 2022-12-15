package problems

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"alexi.ch/aoc/2022/lib"
)

type PointDay15 struct {
	x int64
	y int64
}

type PlaceDay15 struct {
	ptype rune
	dist  int64
	coord *PointDay15
}

type MapDay15 map[PointDay15]*PlaceDay15

type Day15 struct {
	s1        uint64
	s2        uint64
	tl        PointDay15
	br        PointDay15
	maxDist   int64
	testy1    int64
	maxCoord2 int64
	cave      MapDay15
}

func NewDay15() Day15 {
	return Day15{s1: 0, s2: 0, cave: make(MapDay15), tl: PointDay15{x: math.MaxInt64, y: math.MaxInt64}, br: PointDay15{x: math.MinInt64, y: math.MinInt64}, maxDist: 0}
}

func (d *Day15) Title() string {
	return "Day 15 - Beacon Exclusion Zone"
}

func (d *Day15) Setup() {
	// var lines = lib.ReadLines("data/15-test.txt")
	// d.testy1 = 10
	// d.maxCoord2 = 20

	var lines = lib.ReadLines("data/15-data.txt")
	d.testy1 = 2000000
	d.maxCoord2 = 4000000

	// matcher for "Sensor at x=20, y=1: closest beacon is at x=15, y=3"
	var matcher = regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	for _, line := range lines {
		var groups = matcher.FindStringSubmatch(line)
		if len(groups) == 5 {
			x1, err := strconv.ParseInt(groups[1], 10, 64)
			lib.Check(err)
			y1, err := strconv.ParseInt(groups[2], 10, 64)
			lib.Check(err)
			x2, err := strconv.ParseInt(groups[3], 10, 64)
			lib.Check(err)
			y2, err := strconv.ParseInt(groups[4], 10, 64)
			lib.Check(err)
			var p1 = PointDay15{x: x1, y: y1}
			var p2 = PointDay15{x: x2, y: y2}
			var dist = lib.AbsInt64(x1-x2) + lib.AbsInt64(y1-y2)
			d.cave[p1] = &PlaceDay15{ptype: 'S', dist: dist, coord: &p1}
			d.cave[p2] = &PlaceDay15{ptype: 'B', coord: &p2}

			if dist > d.maxDist {
				d.maxDist = dist
			}

			if x1 < d.tl.x {
				d.tl.x = x1
			}
			if x2 < d.tl.x {
				d.tl.x = x2
			}
			if x1 > d.br.x {
				d.br.x = x1
			}
			if x2 > d.br.x {
				d.br.x = x2
			}

			if y1 < d.tl.y {
				d.tl.y = y1
			}
			if y2 < d.tl.y {
				d.tl.y = y2
			}
			if y1 > d.br.y {
				d.br.y = y1
			}
			if y2 > d.br.y {
				d.br.y = y2
			}
		}
	}
	// d.printMap()
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day15) SolveProblem1() {
	// d.printMap()
	var count uint64 = 0
	var y = d.testy1
	var minX = d.tl.x - d.maxDist
	var maxX = d.br.x + d.maxDist

	for x := minX; x <= maxX; x++ {
		var p = PointDay15{x: x, y: y}
		if d.cave[p] == nil {
			if d.checkDeatchZone(&p) {
				count += 1
			}
		}
	}

	d.s1 = count
}

func (d *Day15) SolveProblem2() {

	// Idea:
	// The one single coordinate that is not covered by a sensor
	// must be surrounded (enclosed) by death zones of sensors:
	// so the single coordinate MUST be just outside the death zone "radius"
	// of a sensor.
	// So we check all sensor border coordinates only, within the
	// max square:
	for _, p := range d.cave {
		if p.ptype == 'S' {
			var deathzone, coord = d.checkSensorBorderDeathZone(p)
			if deathzone == false {
				// That has to be it:
				d.s2 = uint64(4000000*coord.x + coord.y)
				return
			}
		}
	}
}

func (d *Day15) Solution1() string {

	return fmt.Sprintf("%d", d.s1)
}

func (d *Day15) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day15) printMap() {
	fmt.Println()
	fmt.Printf("Map: x:%d:%d, y: %d:%d\n", d.tl.x, d.br.x, d.tl.y, d.br.y)
	for y := d.tl.y - 5; y <= d.br.y+5; y++ {
		for x := d.tl.x - 5; x <= d.br.x+5; x++ {
			var p = d.cave[PointDay15{x: x, y: y}]
			if p != nil {
				fmt.Printf("%c", p.ptype)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Day15) checkDeatchZone(coord *PointDay15) bool {
	for _, p := range d.cave {
		// if (p.ptype == 'S' || p.ptype == 'B') && p.coord == coord {
		// 	return false
		// }
		if p.ptype == 'S' {
			var dist = lib.AbsInt64(coord.x-p.coord.x) + lib.AbsInt64(coord.y-p.coord.y)
			if dist <= p.dist {
				return true
			}
		}
	}
	return false
}

/**
 * This method checks all coordinates around the sensor's death zone
 * (just +1 outside the zone).
 * If we find a coordinate that is NOT covered by any other sensor,
 * that is the single spot needed.
 *
 * Returns false and the coordinate if the coordinate is at a free spot (the one we want!)
 * Returns true (and a useless coordinate) if all coordinates around the sensor's death zone
 * are also death zones
 */
func (d *Day15) checkSensorBorderDeathZone(sensor *PlaceDay15) (bool, PointDay15) {
	var x int64 = 0
	var y int64 = 0
	var maxcoord int64 = d.maxCoord2

	var sensorDist = sensor.dist + 1
	for dt := int64(0); dt <= sensorDist; dt++ {
		// top, left
		x = sensor.coord.x - dt
		y = sensor.coord.y - (sensorDist - dt)
		p := PointDay15{x: x, y: y}
		if d.cave[p] != nil {
			return true, p
		}
		if x < 0 || y < 0 || x > maxcoord || y > maxcoord {
			return true, p
		}
		if d.checkDeatchZone(&p) == false {
			return false, p
		}

		// top, right
		x = sensor.coord.x - dt
		y = sensor.coord.y + (sensorDist - dt)
		p = PointDay15{x: x, y: y}
		if d.cave[p] != nil {
			return true, p
		}
		if x < 0 || y < 0 || x > maxcoord || y > maxcoord {
			return true, p
		}
		if d.checkDeatchZone(&p) == false {
			return false, p
		}

		// bottom, left
		x = sensor.coord.x + dt
		y = sensor.coord.y - (sensorDist - dt)
		p = PointDay15{x: x, y: y}
		if d.cave[p] != nil {
			return true, p
		}
		if x < 0 || y < 0 || x > maxcoord || y > maxcoord {
			return true, p
		}
		if d.checkDeatchZone(&p) == false {
			return false, p
		}

		// bottom, right
		x = sensor.coord.x + dt
		y = sensor.coord.y + (sensorDist - dt)
		p = PointDay15{x: x, y: y}
		if d.cave[p] != nil {
			return true, p
		}
		if x < 0 || y < 0 || x > maxcoord || y > maxcoord {
			return true, p
		}
		if d.checkDeatchZone(&p) == false {
			return false, p
		}
	}
	return true, PointDay15{x: 0, y: 0}
}
