package problems

import (
	"container/list"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sort"

	"alexi.ch/aoc/2022/lib"
)

type Heatmap map[Point]*Place

type Place struct {
	coord    Point
	height   rune
	visited  bool
	distance uint64
}

type Point struct {
	x uint64
	y uint64
}

type Day12 struct {
	s1      uint64
	s2      uint64
	start   Point
	end     Point
	maxX    uint64
	maxY    uint64
	heatmap Heatmap
}

func NewDay12() Day12 {
	return Day12{s1: 0, s2: 0, heatmap: make(Heatmap)}
}

func (d *Day12) Title() string {
	return "Day 12 - Hill Climbing Algorithm"
}

func (d *Day12) Setup() {
	// var lines = lib.ReadLines("data/12-test.txt")
	var lines = lib.ReadLines("data/12-data.txt")
	for y, line := range lines {
		for x, chr := range line {
			var point = Point{x: uint64(x), y: uint64(y)}
			var place = Place{coord: Point{x: uint64(x), y: uint64(y)}, visited: false, distance: 0}
			if uint64(x) > d.maxX {
				d.maxX = uint64(x)
			}
			if uint64(y) > d.maxY {
				d.maxY = uint64(y)
			}
			if chr == 'S' {
				place.height = 'a'
				place.distance = 0
				d.start = Point{x: uint64(x), y: uint64(y)}
			} else if chr == 'E' {
				place.height = 'z'
				d.end = Point{x: uint64(x), y: uint64(y)}
			} else {
				place.height = chr
			}
			d.heatmap[point] = &place
		}
	}
	// d.printMap()
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day12) SolveProblem1() {
	// implement a djikstra here
	var todo = list.New()
	var start_point = d.start
	var end_point = d.end
	var start = d.heatmap[start_point]
	var end = d.heatmap[end_point]

	todo.PushBack(start)
	for todo.Len() > 0 {
		var act = todo.Front()
		todo.Remove(act)
		var act_place = act.Value.(*Place)
		d.djikstra(act_place, end, todo)
	}
	/** Uncomment the line below to generate a map of the solution as PNG image: **/
	// d.drawImage("day12-map-1.png")
	d.s1 = end.distance
}

func (d *Day12) SolveProblem2() {

	var shortest_paths = make([]uint64, 0)
	var end_point = d.end
	var end = d.heatmap[end_point]

	for _, entry := range d.heatmap {
		d.resetMap()
		if entry.height > 'a' {
			continue
		}

		// implement a djikstra here
		var todo = list.New()

		todo.PushBack(entry)
		for todo.Len() > 0 {
			var act = todo.Front()
			todo.Remove(act)
			var act_place = act.Value.(*Place)
			d.djikstra(act_place, end, todo)
		}
		if end.distance > 0 {
			shortest_paths = append(shortest_paths, end.distance)
		}
	}
	// fmt.Printf("%v\n", shortest_paths)

	sort.Slice(shortest_paths, func(i, j int) bool {
		return shortest_paths[i] < shortest_paths[j]
	})

	d.s2 = shortest_paths[0]
}

func (d *Day12) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day12) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func (d *Day12) printMap() {
	var y uint64 = 0
	var x uint64 = 0
	fmt.Println()
	fmt.Println()
	for y = 0; y < d.maxY; y++ {
		for x = 0; x < d.maxX; x++ {
			var loc = d.heatmap[Point{x: x, y: y}]
			if x == d.start.x && y == d.start.y {
				fmt.Print("S")
			} else if x == d.end.x && y == d.end.y {
				fmt.Print("E")
			} else {
				fmt.Printf("%c", loc.height)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Day12) djikstra(act *Place, end *Place, todo *list.List) {
	// just to make sure:
	if act.visited == true {
		return
	}
	// fmt.Printf("Working on: %v\n", act.coord)

	// mark act as visited
	act.visited = true

	// update distances to allowed neighbours:
	var neighbours = list.New()
	var left = d.heatmap[Point{x: act.coord.x - 1, y: act.coord.y}]
	if left != nil && left.height <= act.height+1 {
		neighbours.PushBack(left)
	}
	var right = d.heatmap[Point{x: act.coord.x + 1, y: act.coord.y}]
	if right != nil && right.height <= act.height+1 {
		neighbours.PushBack(right)
	}
	var up = d.heatmap[Point{x: act.coord.x, y: act.coord.y - 1}]
	if up != nil && up.height <= act.height+1 {
		neighbours.PushBack(up)
	}
	var down = d.heatmap[Point{x: act.coord.x, y: act.coord.y + 1}]
	if down != nil && down.height <= act.height+1 {
		neighbours.PushBack(down)
	}

	var neighbour = neighbours.Front()
	for neighbour != nil {
		var n = neighbour.Value.(*Place)
		neighbours.Remove(neighbour)

		// can we reach the next location with a shorter path? update:
		if n.distance == 0 || n.distance > act.distance+1 {
			n.distance = act.distance + 1
		}

		// do we need to add it to the todo list?
		if !n.visited && n != end {
			todo.PushBack(n)
		}

		neighbour = neighbours.Front()
	}
}

func (d *Day12) resetMap() {
	for _, place := range d.heatmap {
		place.visited = false
		place.distance = 0
	}
}

func (d *Day12) drawImage(out_path string) {
	var rect = image.Rect(0, 0, int(d.maxX), int(d.maxY))
	var img = image.NewRGBA(rect)

	// draw heat map:
	for _, place := range d.heatmap {
		var col_val = uint8(float32(place.height-'a') / 25.0 * 255.0)
		img.Set(int(place.coord.x), int(place.coord.y), color.RGBA{R: col_val, G: col_val, B: col_val, A: 255})
	}
	// draw path:
	var act = d.heatmap[d.end]
	for act != nil && act != d.heatmap[d.start] {
		img.Set(int(act.coord.x), int(act.coord.y), color.RGBA{R: 255, G: 0, B: 0, A: 255})

		var neighbours = make([]*Place, 0)
		var neighbour *Place

		// up
		neighbour = d.heatmap[Point{x: act.coord.x, y: act.coord.y - 1}]
		if neighbour != nil && neighbour.distance < act.distance {
			neighbours = append(neighbours, neighbour)
		}
		// down
		neighbour = d.heatmap[Point{x: act.coord.x, y: act.coord.y + 1}]
		if neighbour != nil && neighbour.distance < act.distance {
			neighbours = append(neighbours, neighbour)
		}

		// right
		neighbour = d.heatmap[Point{x: act.coord.x + 1, y: act.coord.y}]
		if neighbour != nil && neighbour.distance < act.distance {
			neighbours = append(neighbours, neighbour)
		}

		// left
		neighbour = d.heatmap[Point{x: act.coord.x - 1, y: act.coord.y}]
		if neighbour != nil && neighbour.distance < act.distance {
			neighbours = append(neighbours, neighbour)
		}

		if len(neighbours) > 0 {
			neighbour = neighbours[0]
			for _, n := range neighbours[1:] {
				if n.distance < neighbour.distance {
					neighbour = n
				}
			}
			act = neighbour
		} else {
			break
		}
	}

	out, err := os.Create(out_path)
	defer out.Close()
	lib.Check(err)
	png.Encode(out, img)
}
