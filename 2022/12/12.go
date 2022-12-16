package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"sync"
	"time"
)

type Tile struct {
	Height int
	North  *Tile
	South  *Tile
	East   *Tile
	West   *Tile
}

var Elevation = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
	"i": 8,
	"j": 9,
	"k": 10,
	"l": 11,
	"m": 12,
	"n": 13,
	"o": 14,
	"p": 15,
	"q": 16,
	"r": 17,
	"s": 18,
	"t": 19,
	"u": 20,
	"v": 21,
	"w": 22,
	"x": 23,
	"y": 24,
	"z": 25,
}

type Grid map[int]map[int]*Tile

var start, end []int

type Journey struct {
	lock        sync.Mutex
	Path        map[int]map[int]int
	Destination []int
}

func NewJourney(destination []int) *Journey {
	path := make(map[int]map[int]int)
	j := Journey{
		lock:        sync.Mutex{},
		Path:        path,
		Destination: destination,
	}
	return &j
}

func (j *Journey) addMove(coords []int, step int) bool {
	j.lock.Lock()
	defer j.lock.Unlock()
	if _, ok := j.Path[coords[0]]; !ok {
		j.Path[coords[0]] = make(map[int]int)
	}
	if _, ok := j.Path[coords[0]][coords[1]]; !ok {
		j.Path[coords[0]][coords[1]] = step
		return true
	}
	if j.Path[coords[0]][coords[1]] <= step {
		return false
	}
	j.Path[coords[0]][coords[1]] = step
	return true
}

func (j *Journey) travel(from []int) {
	_ = j.addMove(from, 0)
	move(j, from, 0)
}

func move(j *Journey, from []int, steps int) {
	tile := grid[from[0]][from[1]]
	steps += 1
	if from[0] == j.Destination[0] && from[1] == j.Destination[1] {
		_ = j.addMove(from, steps)
		return
	}
	for i := range Moves {
		switch Moves[i] {
		case "n":
			if tile.North == nil {
				continue
			}

			if tile.North.Height-1 <= tile.Height {
				newCoords := []int{from[0] - 1, from[1]}
				if !j.addMove(newCoords, steps) {
					break
				}
				move(j, newCoords, steps)
			}
		case "s":
			if tile.South == nil {
				continue
			}

			if tile.South.Height-1 <= tile.Height {
				newCoords := []int{from[0] + 1, from[1]}
				if !j.addMove(newCoords, steps) {
					break
				}

				move(j, newCoords, steps)
			}
		case "e":
			if tile.East == nil {
				continue
			}

			if tile.East.Height-1 <= tile.Height {
				newCoords := []int{from[0], from[1] + 1}
				if !j.addMove(newCoords, steps) {
					break
				}
				move(j, newCoords, steps)
			}
		case "w":
			if tile.West == nil {
				continue
			}

			if tile.West.Height-1 <= tile.Height {
				newCoords := []int{from[0], from[1] - 1}
				if !j.addMove(newCoords, steps) {
					break
				}
				move(j, newCoords, steps)
			}
		}
	}
}

var Moves = []string{"n", "s", "e", "w"}

var grid = Grid{}

// day 2
var startPoints [][]int

func main() {
	//f, err := os.Open("./2022/12/test.txt")
	f, err := os.Open("./2022/12/12.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()

	grid = GenerateGrid(f)

	journey := NewJourney(end)
	journey.travel(start)

	fmt.Printf("Journey1 Steps: %v\n\n", journey.Path[end[0]][end[1]])

	//day 2
	var steps []int

	quit := make(chan bool)
	go func() {
		fmt.Print("\033[s")
		chars := []string{
			"", "|", "-|", "--|", "---|", "----|",
			"---||", "--|||", "-||||", "|||||",
			"||||-", "|||--", "||---", "|----", "-----",
			"----", "---", "--", "-",
		}
		i := 0
		for {
			select {
			case <-quit:
				fmt.Print("\033[1A\033[K")
				return
			default:
				no := i % len(chars)
				fmt.Print("\033[1A\033[K" + chars[no] + "\n")
				i += 1
				time.Sleep(250 * time.Millisecond)
			}
		}
	}()

	for i := 0; i < len(startPoints); i++ {
		j2 := NewJourney(end)
		j2.travel(startPoints[i])
		if j2.Path[end[0]][end[1]] > 0 {
			steps = append(steps, j2.Path[end[0]][end[1]])
		}
	}
	quit <- true
	sort.Slice(steps, func(i, j int) bool {
		return steps[i] <= steps[j]
	})
	fmt.Printf("\nJourney2 Steps: %v\n", steps[0])
}

func GenerateGrid(f io.Reader) Grid {
	scanner := bufio.NewScanner(f)
	grid := map[int]map[int]*Tile{}
	start = []int{}
	end = []int{}
	row := 0
	for scanner.Scan() {
		tiles := scanner.Text()
		grid[row] = map[int]*Tile{}
		for i := 0; i < len(tiles); i++ {
			char := string(tiles[i])
			if char == "S" {
				start = append(start, row, i)
				char = "a"
			}
			if char == "E" {
				end = append(end, row, i)
				char = "z"
			}
			// day 2
			if char == "a" {
				startPoints = append(startPoints, []int{row, i})
			}

			t := Tile{
				Height: Elevation[char],
			}
			if _, ok := grid[row-1][i]; ok {
				north := grid[row-1][i]
				t.North = north
				north.South = &t
			}
			if _, ok := grid[row][i-1]; ok {
				west := grid[row][i-1]
				t.West = west
				west.East = &t
			}

			grid[row][i] = &t
		}
		row += 1
	}

	return grid
}
