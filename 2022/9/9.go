package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Knot struct {
	X        int
	Y        int
	History  map[string]bool
	observer []*Knot
}

func (k *Knot) register(Observer *Knot) {
	k.observer = append(k.observer, Observer)
}

func (k *Knot) notifyAll() {
	for i := range k.observer {
		k.observer[i].update(k.X, k.Y)
	}
}

func (k *Knot) update(x, y int) {
	xdiff := x - k.X
	ydiff := y - k.Y
	if math.Abs(float64(xdiff)) > 1 || math.Abs(float64(ydiff)) > 1 {
		dx := 0
		dy := 0
		if xdiff > 0 {
			dx = 1
		}
		if xdiff < 0 {
			dx = -1
		}
		if ydiff > 0 {
			dy = 1
		}
		if ydiff < 0 {
			dy = -1
		}
		k.step(dx, dy)
	}
}

func (k *Knot) step(x, y int) {
	k.X += x
	k.Y += y
	key := fmt.Sprintf("%d,%d", k.X, k.Y)
	if _, ok := k.History[key]; !ok {
		k.History[key] = true
	}
	k.notifyAll()
}

func (k *Knot) Move(direction string, distance int) {
	var x, y int
	x = 0
	y = 0
	switch direction {
	case "U":
		x = 1
	case "D":
		x = -1
	case "L":
		y = -1
	case "R":
		y = 1
	}
	for i := 0; i < distance; i++ {
		k.step(x, y)
	}
}

func main() {
	f, err := os.Open("./2022/9/9.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()

	head := &Knot{
		X:       0,
		Y:       0,
		History: map[string]bool{"0,0": true},
	}
	tail := &Knot{
		X:       0,
		Y:       0,
		History: map[string]bool{"0,0": true},
	}

	head.register(tail)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := scanner.Text()
		move := strings.Split(data, " ")
		direction := move[0]
		distance, _ := strconv.Atoi(move[1])
		head.Move(direction, distance)
	}

	fmt.Printf("Head %v %v\nTail %v %v\n", head.X, head.Y, tail.X, tail.Y)
	fmt.Printf("History: %v\n", tail.History)
	fmt.Printf("UNique: %v\n", len(tail.History))
}
