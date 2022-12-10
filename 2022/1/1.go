package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Elf struct {
	Name  int
	Meals []Meal
	Total int
}

func (e *Elf) AddMeal(meal int) {
	e.Meals = append(e.Meals, Meal{meal})
}

func (e *Elf) GetCalories() int {
	tot := 0
	for i := range e.Meals {
		tot += e.Meals[i].Calories
	}
	return tot
}

type Meal struct {
	Calories int
}

var input []string

func main() {
	f, err := os.Open("./2022/1/1.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	Elves := []Elf{{}}
	idx := 0
	elf := Elf{Name: idx}
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			elf.Total = elf.GetCalories()
			fmt.Printf("Elf: %v\n", elf.Total)
			if len(Elves) == 0 {
				Elves[0] = elf
			} else {
			QUIT:
				for i := range Elves {
					if i == len(Elves)-1 {
						Elves = append(Elves, elf)
					}
					if Elves[i].Total < elf.Total {
						copy(Elves[(i+1):], Elves[i:])
						Elves[i] = elf
						break QUIT
					}
				}
			}
			idx += 1
			elf = Elf{Name: idx}
			continue
		}
		cals, err := strconv.Atoi(txt)
		if err != nil {
			fmt.Printf("Couldn't get calories: %+v\n", txt)
		}
		elf.AddMeal(cals)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Elves: %+v", Elves)
}
