package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction interface {
	Execute() int
	Cycles() int
}

type AddX struct {
	cycles int
	Value  int
}

func (a *AddX) Execute() int {
	Register += a.Value
	return 0
}

func (a *AddX) Cycles() int {
	c := a.cycles
	a.cycles = 0
	return c
}

type Noop struct {
	cycles int
}

func (n *Noop) Execute() int {
	return 0
}

func (n *Noop) Cycles() int {
	return n.cycles
}

var Register = 1

func main() {
	f, err := os.Open("./2022/10/10.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()

	Instructions := []Instruction{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := scanner.Text()
		instruction := strings.Split(data, " ")
		switch instruction[0] {
		case "addx":
			nmbr, _ := strconv.Atoi(instruction[1])
			Instructions = append(Instructions, &AddX{cycles: 2, Value: nmbr})
		case "noop":
			Instructions = append(Instructions, &Noop{cycles: 1})
		}
	}

	//working := make(chan bool, 1)

	cycle := 0
	command := 0
	signal := 0
	display := ""

	for {
		if command >= len(Instructions) { // if no commands left to execute, leave
			break
		}
		cycle += 1

		if (cycle-20)%40 == 0 {
			//fmt.Printf("++ Cycle %v Register %v Weight %v\n", cycle, Register, cycle*Register)
			signal += cycle * Register
		}

		cursorPos := (cycle - 1) % 40

		if cursorPos == Register-1 || cursorPos == Register || cursorPos == Register+1 {
			display += "#"
		} else {
			display += "."
		}

		cmmd := Instructions[command]

		if cmmd.Cycles() > 1 {
			continue
		}
		command += 1
		cmmd.Execute()
	}
	fmt.Printf("|| END %v Register %v SUM: %v\n\n", cycle, Register, signal)

	for i := 0; i < 6; i++ {
		fmt.Printf("%v\n", display[(40*i):(40*i)+40])
	}
}
