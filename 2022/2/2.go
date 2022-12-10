package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shape struct {
	Score  int
	Result map[string]int
}

var result = map[string]Shape{
	"X": {
		1,
		map[string]int{
			"A": 3,
			"B": 0,
			"C": 6,
		},
	},
	"Y": {
		2,
		map[string]int{
			"A": 6,
			"B": 3,
			"C": 0,
		},
	},
	"Z": {
		3,
		map[string]int{
			"A": 0,
			"B": 6,
			"C": 3,
		},
	},
}

var scheme = map[string]Shape{
	"X": {
		0,
		map[string]int{
			"A": 3,
			"B": 1,
			"C": 2,
		},
	},
	"Y": {
		3,
		map[string]int{
			"A": 1,
			"B": 2,
			"C": 3,
		},
	},
	"Z": {
		6,
		map[string]int{
			"A": 2,
			"B": 3,
			"C": 1,
		},
	},
}

func main() {
	f, err := os.Open("./2022/2/2.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	myTotal := 0
	for scanner.Scan() {
		res := scanner.Text()
		turn := strings.Split(res, " ")
		myTotal += scheme[turn[1]].Score + scheme[turn[1]].Result[turn[0]]
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Total: %v\n", myTotal)
}
