package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./2022/4/4.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	contained := 0
	overlap := 0
	for scanner.Scan() {
		res := scanner.Text()
		elf := strings.Split(res, ",")

		elf1 := strings.Split(elf[0], "-")
		elf2 := strings.Split(elf[1], "-")
		a1, _ := strconv.Atoi(elf1[0])
		a2, _ := strconv.Atoi(elf1[1])
		b1, _ := strconv.Atoi(elf2[0])
		b2, _ := strconv.Atoi(elf2[1])
		if (a1 <= b1 && a2 >= b2) || (b1 <= a1 && b2 >= a2) {
			contained += 1
		}
		if (a1 <= b1 && a2 >= b1) || (b1 <= a1 && b2 >= a1) {
			overlap += 1
		}
		fmt.Printf("Elf1: %v - %v\nElf2: %v - %v\n", a1, a2, b1, b2)
	}

	fmt.Printf("Contained: %v\nOverlap: %v", contained, overlap)
}
