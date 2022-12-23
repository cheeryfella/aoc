package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Valve struct {
	name         string
	flow         int
	open         bool
	Destinations []string
	visit int
}

func (v *Valve) PickDestination() string {
	//pressure := 0
	//for i := range v.Destinations {
	//	if !Valves[v.Destinations[i]].open && Valves[v.Destinations[i]].flow > pressure {
	//		pressure = Valves[v.Destinations[i]].flow
	//		desination = Valves[v.Destinations[i]]
	//	}
	//}
	if v.visit >= len(v.Destinations) {
		v.visit = 0
	}
	desination := v.visit
	v.visit +=1
	return v.Destinations[desination]
}

var Valves = make(map[string]*Valve)

func main() {
	f, err := os.Open("./2022/16/test.txt")
	//f, err := os.Open("./2022/16/16.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		valve := ParseValveData(scanner.Text())
		Valves[valve.name] = &valve
	}

	current := `AA`
	destination := ""
	action := ""
	pressure := 0
	for i := 0; i <= 30; i++ {
		pressure += CurrentRelief(Valves)
		action = ""
		v := Valves[current]
		switch true {
		case v.name != destination && destination != "":
			current = destination
			action += fmt.Sprintf("Moved to: %s ", destination)
		case v.flow > 0 && !v.open:
			v.open = true
			action += fmt.Sprintf("Opened valve: %s ", current)
			fallthrough
		case destination == "":
			destination = v.PickDestination()
			//action += fmt.Sprintf("Moving to: %s ", destination)
		}
		fmt.Printf("Minute: %d\n\tAction: %s\n\tPressure released: %d\n", i, action, pressure)
	}
}

func CurrentRelief(valves map[string]*Valve) int {
	rate := 0
	for i := range valves {
		if valves[i].open {
			rate += valves[i].flow
		}
	}
	return rate
}

func ParseValveData(text string) Valve {
	regex := regexp.MustCompile(
		`Valve (?P<id>[A-Z]{2}) has flow rate=(?P<flow>\d+); tunnel(s?) lead(s?) to valve(s?) (?P<dest>[A-Z, ]+)+`,
	)
	valve := Valve{}
	matches := regex.FindStringSubmatch(text)

	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" {
			switch name {
			case "id":
				valve.name = matches[i]
			case "flow":
				valve.flow, _ = strconv.Atoi(matches[i])
			case "dest":
				valve.Destinations = strings.Split(matches[i], ", ")
			}
		}
	}

	return valve
}
