package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"regexp"
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

func main() {
	//f, err := os.Open("./2022/15/test.txt")
	f, err := os.Open("./2022/15/15.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	line := make(map[int]bool)
	yLine := 2000000
	for scanner.Scan() {
		sensor := ParseData(scanner.Text())
		// day 1
		day1(sensor, yLine, line)

	}

	// day 1
	fmt.Printf("line\n%+v\n", len(line))
	//drawGrid(grid)
}

func day1(sensor *Sensor, yLine int, line map[int]bool) {
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

func drawGrid(grid Grid) {
	fmt.Print("\033[s")
	fmt.Printf("Boundary: %+v\n\n", GridBoundary)

	//fmt.Print("\033[H\033[2J")
	minx := len(strconv.Itoa(GridBoundary.minX))-1
	maxx := len(strconv.Itoa(GridBoundary.maxX))+1
	xlen := int(math.Max(float64(minx), float64(maxx)))
	b := bytes.Buffer{}
	for i := 0; i < xlen; i++ {
		b.WriteString("            ")
		for x := GridBoundary.minX-1; x <= GridBoundary.maxX+1; x++ {
			numstr := strconv.Itoa(x)
			if i < len(numstr) {
				b.WriteString(fmt.Sprintf("%v", string(numstr[i])))
			} else {
				b.WriteString(" ")
			}
		}
		b.WriteString("\n")
	}
	for i := GridBoundary.minY-1; i <= GridBoundary.maxY+1; i++ {
		b.WriteString(fmt.Sprintf("%-11d ", i))
		for x := GridBoundary.minX-1; x <= GridBoundary.maxX+1; x++ {
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

	xdiff := sensor.x-sensor.Beacon.x
	ydiff := sensor.y-sensor.Beacon.y
	if xdiff < 0 {
		xdiff *= -1
	}
	if ydiff < 0{
		ydiff *= -1
	}
	sensor.distance = xdiff + ydiff

	xplot[sensor.x] = sensor.y

	return &sensor
}
