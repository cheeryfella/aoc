package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

type Cave map[coord]rune

func (c Cave) drawRock(start, end coord) {
	current := coord{
		x: start.x,
		y: start.y,
	}
	for current.x != end.x || current.y != end.y {
		c[current] = '#'
		switch {
		case current.x < end.x:
			current.x++
		case current.y < end.y:
			current.y++
		case current.x > end.x:
			current.x--
		case current.y > end.y:
			current.y--
		}
	}
	c[end] = '#'
	boundary.maxX = int(
		math.Max(
			float64(end.x),
			math.Max(
				float64(boundary.maxX),
				float64(start.x))),
	)
	boundary.maxY = int(
		math.Max(
			float64(end.y),
			math.Max(
				float64(boundary.maxY),
				float64(start.y))),
	)
	boundary.minX = int(
		math.Min(
			float64(end.x),
			math.Min(
				float64(boundary.minX),
				float64(start.x))),
	)
	boundary.minY = int(
		math.Min(
			float64(end.y),
			math.Min(
				float64(boundary.minY),
				float64(start.y))),
	)
}

type CaveBoundary struct {
	minX int
	minY int
	maxX int
	maxY int
}

var boundary = CaveBoundary{minX: 1000, minY: 1000}

func main() {
	f, err := os.Open("./2022/14/14.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()

	cave := MapCave(f)

	//map1 := cave
	//
	//sand := poorSand(map1, Abyss)
	//
	//printMap(map1)
	//fmt.Printf("Grains: %v\nBoundary: %+v\n", sand-1, boundary)

	// day 2
	floorY := boundary.maxY + 2
	minX := 500 - floorY
	maxX := 500 + floorY
	map2 := cave
	map2.drawRock(coord{minX, floorY}, coord{maxX, floorY})

	sand2 := poorSand(map2, FullUp)
	printMap(map2)
	fmt.Printf("Grains: %v\nBoundary: %+v\n", sand2, boundary)
}

func poorSand(caveX Cave, condition func(cave Cave, point coord) bool) int {
	sand := 0
	down := coord{}
	downleft := coord{}
	downright := coord{}
QUIT:
	for {
		grain := coord{
			x: 500,
			y: 0,
		}
		sand++
		for {
			caveX[grain] = '▓'
			//printMap(caveX)
			//time.Sleep(50*time.Millisecond)
			if condition(caveX, grain) {
				//fmt.Printf("Boundary reached grain:%+v boundary:%+v\n", grain, boundary)
				break QUIT
			}
			down = coord{
				x: grain.x,
				y: grain.y + 1,
			}
			downleft = coord{
				x: grain.x - 1,
				y: grain.y + 1,
			}
			downright = coord{
				x: grain.x + 1,
				y: grain.y + 1,
			}
			if caveX[down] < '#' {
				caveX[grain] = 0
				grain.y += 1
				continue
			}
			if caveX[downleft] < '#' {
				caveX[grain] = 0
				grain.y += 1
				grain.x -= 1
				continue
			}
			if caveX[downright] < '#' {
				caveX[grain] = 0
				grain.y += 1
				grain.x += 1
				continue
			}
			break
		}
	}
	return sand
}

var Abyss = func(cave Cave, point coord) bool {
	if point.y >= boundary.maxY {
		return true
	}
	return false
}

var FullUp = func(cave Cave, point coord) bool {
	if point.x == 500 && point.y == 0 {
		pointL := coord{499, 1}
		pointR := coord{501, 1}
		if cave[pointL] == '▓' && cave[pointR] == '▓' && cave[pointR] == '▓' {
			return true
		}
	}
	return false
}

func MapCave(f io.Reader) Cave {
	scanner := bufio.NewScanner(f)
	cave := make(Cave)

	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), " -> ")
		for i := range coords[:len(coords)-1] {
			start := generateCoordinate(coords[i])
			end := generateCoordinate(coords[i+1])
			cave.drawRock(start, end)
		}
	}

	return cave
}

func generateCoordinate(point string) coord {
	values := strings.Split(point, ",")
	x, _ := strconv.Atoi(values[0])
	y, _ := strconv.Atoi(values[1])
	return coord{
		x: x,
		y: y,
	}
}

func printMap(cave Cave) {
	fmt.Print("\033[s")
	//fmt.Print("\033[H\033[2J")

	for y := 0; y < 3; y++ {
		fmt.Print("    ")
		for x := boundary.minX - 1; x <= boundary.maxX+1; x++ {
			numstr := strconv.Itoa(x)
			fmt.Printf(" %v", string(numstr[y]))
		}
		fmt.Println()
	}
	for y := 0; y <= boundary.maxY+3; y++ {
		fmt.Printf("%-3d ", y)
		for x := boundary.minX - 1; x <= boundary.maxX+1; x++ {
			if cave[coord{x, y}] == 0 {
				fmt.Print(" .")
			} else {
				fmt.Print(" " + string(cave[coord{x, y}]))
			}
		}
		fmt.Println()
	}
}
