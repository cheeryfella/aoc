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
	// day 1
	//tail := &Knot{
	//	X:       0,
	//	Y:       0,
	//	History: map[string]bool{"0,0": true},
	//}
	//head.register(tail)
	k1 := &Knot{
		X:       0,
		Y:       0,
		History: map[string]bool{"0,0": true},
	}
	k2 := &Knot{
		X:       0,
		Y:       0,
		History: map[string]bool{"0,0": true},
	}
	k3 := &Knot{
		X:       0,
		Y:       0,
		History: map[string]bool{"0,0": true},
	}
	k4 := &Knot{
		X:       0,
		Y:       0,
		History: map[string]bool{"0,0": true},
	}
	k5 := &Knot{
		X:       0,
		Y:       0,
		History: map[string]bool{"0,0": true},
	}
	k6 := &Knot{
		X:       0,
		Y:       0,
		History: map[string]bool{"0,0": true},
	}
	k7 := &Knot{
		X:       0,
		Y:       0,
		History: map[string]bool{"0,0": true},
	}
	k8 := &Knot{
		X:       0,
		Y:       0,
		History: map[string]bool{"0,0": true},
	}
	k9 := &Knot{
		X:       0,
		Y:       0,
		History: map[string]bool{"0,0": true},
	}
	head.register(k1)
	k1.register(k2)
	k2.register(k3)
	k3.register(k4)
	k4.register(k5)
	k5.register(k6)
	k6.register(k7)
	k7.register(k8)
	k8.register(k9)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := scanner.Text()
		move := strings.Split(data, " ")
		direction := move[0]
		distance, _ := strconv.Atoi(move[1])
		head.Move(direction, distance)
	}

	fmt.Printf("Head %v %v\nTail %v %v\n", head.X, head.Y, k9.X, k9.Y)
	fmt.Printf("UNique: %v\n", len(k9.History))
}
