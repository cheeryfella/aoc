package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Item interface {
}

type coord struct {
	x int
	y int
}

type Beacon struct {
	coord
}

type Sensor struct {
	coord
	Beacon   Beacon
	distance int
	oor      map[coord]bool
}

type Grid map[coord]Item

type XPlot map[int]int
type Boundary struct {
	minX int
	minY int
	maxX int
	maxY int
}

var GridBoundary = Boundary{
	10000000,
	10000000,
	0,
	0,
}

func (g Grid) AddSensor(sensor *Sensor) {
	g[coord{sensor.x, sensor.y}] = *sensor
	g[coord{sensor.Beacon.x, sensor.Beacon.y}] = sensor.Beacon
	GridBoundary.maxX = int(
		math.Max(
			float64(GridBoundary.maxX),
			math.Max(
				float64(sensor.x),
				float64(sensor.Beacon.x))),
	)
	GridBoundary.maxY = int(
		math.Max(
			float64(GridBoundary.maxY),
			math.Max(
				float64(sensor.y),
				float64(sensor.Beacon.y))),
	)
	GridBoundary.minX = int(
		math.Min(
			float64(GridBoundary.minX),
			math.Min(
				float64(sensor.x),
				float64(sensor.Beacon.x))),
	)
	GridBoundary.minY = int(
		math.Min(
			float64(GridBoundary.minY),
			math.Min(
				float64(sensor.y),
				float64(sensor.Beacon.y))),
	)
}

var grid = Grid{}
var xplot = XPlot{}

// day 1
var line = make(map[int]bool)
var yLine = 2000000

// day 2
var candidateBeacon = make(map[coord]int)
var coordLimit = 4000000

func main() {
	f, err := os.Open("./2022/15/15.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		sensor := ParseData(scanner.Text())
		// day 1
		//day1(sensor)

		// day 2
		day2(sensor)
	}

	//day 1
	//fmt.Printf("line\n%+v\n\n\n", len(line))

	//drawGrid(grid)

	//day 2
	keys := make([]coord, 0, len(candidateBeacon))

	for key := range candidateBeacon {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return candidateBeacon[keys[i]] > candidateBeacon[keys[j]]
	})

	fmt.Printf("candidate\n%v -- %+v -- frequency %v\n", keys[0], candidateBeacon[keys[0]], frequency(keys[0]))
}

func day1(sensor *Sensor) {
	yDistance := sensor.y - yLine
	if sensor.y-yLine < 0 {
		yDistance *= -1
	}

	for i := 0; i <= sensor.distance-yDistance; i++ {
		line[sensor.x+i] = true
		line[sensor.x-i] = true
	}

	if sensor.Beacon.y == yLine {
		delete(line, sensor.Beacon.x)
	}
}

func day2(sensor *Sensor) {
	outOfRange := sensor.distance + 1
	for i := 0; i < outOfRange; i++ {
		if sensor.x+i > 0 && sensor.x+i < coordLimit {
			if (sensor.y+outOfRange)-i > 0 && (sensor.y+outOfRange)-i < coordLimit {
				candidateBeacon[coord{sensor.x + i, (sensor.y + outOfRange) - i}] += 1
			}
			if (sensor.y-outOfRange)+i > 0 && (sensor.y-outOfRange)+i < coordLimit {
				candidateBeacon[coord{sensor.x + i, (sensor.y - outOfRange) + i}] += 1
			}
		}
		if sensor.x-i > 0 && sensor.x-i < coordLimit {
			if (sensor.y+outOfRange)-i > 0 && (sensor.y+outOfRange)-i < coordLimit {
				candidateBeacon[coord{sensor.x - i, (sensor.y + outOfRange) - i}] += 1
			}
			if (sensor.y-outOfRange)+i > 0 && (sensor.y-outOfRange)+i < coordLimit {
				candidateBeacon[coord{sensor.x - i, (sensor.y - outOfRange) + i}] += 1
			}
		}
	}
}

func drawGrid(grid Grid) {
	fmt.Print("\033[s")
	fmt.Printf("Boundary: %+v\n\n", GridBoundary)

	//fmt.Print("\033[H\033[2J")
	minx := len(strconv.Itoa(GridBoundary.minX)) - 1
	maxx := len(strconv.Itoa(GridBoundary.maxX)) + 1
	xlen := int(math.Max(float64(minx), float64(maxx)))
	b := bytes.Buffer{}
	for i := 0; i < xlen; i++ {
		b.WriteString("            ")
		for x := GridBoundary.minX - 1; x <= GridBoundary.maxX+1; x++ {
			numstr := strconv.Itoa(x)
			if i < len(numstr) {
				b.WriteString(fmt.Sprintf("%v", string(numstr[i])))
			} else {
				b.WriteString(" ")
			}
		}
		b.WriteString("\n")
	}
	for i := GridBoundary.minY - 1; i <= GridBoundary.maxY+1; i++ {
		b.WriteString(fmt.Sprintf("%-11d ", i))
		for x := GridBoundary.minX - 1; x <= GridBoundary.maxX+1; x++ {
			switch grid[coord{x, i}].(type) {
			case Beacon:
				b.WriteString("B")
			case Sensor:
				b.WriteString("S")
			default:
				b.WriteString(".")

			}
		}
		b.WriteString("\n")
	}
	b.WriteString("\n")
	b.WriteTo(os.Stdout)
}

func ParseData(data string) *Sensor {
	regex := regexp.MustCompile(
		`Sensor at x=(?P<sx>-?\d+), y=(?P<sy>-?\d+): closest beacon is at x=(?P<bx>-?\d+), y=(?P<by>-?\d+)`,
	)

	sensor := Sensor{}
	matches := regex.FindStringSubmatch(data)
	oor := make(map[coord]bool)
	sensor.oor = oor

	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" {
			switch name {
			case "sx":
				sensor.x, _ = strconv.Atoi(matches[i])
			case "sy":
				sensor.y, _ = strconv.Atoi(matches[i])
			case "bx":
				sensor.Beacon.x, _ = strconv.Atoi(matches[i])
			case "by":
				sensor.Beacon.y, _ = strconv.Atoi(matches[i])
			}
		}
	}
	s := coord{sensor.x, sensor.y}
	b := coord{sensor.Beacon.x, sensor.Beacon.y}

	sensor.distance = distance(s, b)

	xplot[sensor.x] = sensor.y

	return &sensor
}

func distance(sensor, beacon coord) int {
	xdiff := sensor.x - beacon.x
	ydiff := sensor.y - beacon.y
	if xdiff < 0 {
		xdiff *= -1
	}
	if ydiff < 0 {
		ydiff *= -1
	}
	return xdiff + ydiff
}

func frequency(point coord) int {
	return (point.x * 4000000) + point.y
}
