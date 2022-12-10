package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var stack = map[int][]string{
	1: {"B", "Q", "C"},
	2: {"R", "Q", "W", "Z"},
	3: {"B", "M", "R", "L", "V"},
	4: {"C", "Z", "H", "V", "T", "W"},
	5: {"D", "Z", "H", "B", "N", "V", "G"},
	6: {"H", "N", "P", "C", "J", "F", "V", "Q"},
	7: {"D", "G", "T", "R", "W", "Z", "S"},
	8: {"C", "G", "M", "N", "B", "W", "Z", "P"},
	9: {"N", "J", "B", "M", "W", "Q", "F", "P"},
}

func main() {
	f, err := os.Open("./2022/5/5.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		no, from, to := parseMoves(scanner.Text())
		fmt.Printf("Move: %d from %d to %d\n", no, from, to)
		//for i := 0; i < no; i++ {
		//	fmt.Printf("Moving: %v from %v to %v\n", stack[from], stack[from], stack[to])
		//	stack[to] = append(stack[to], stack[from][len(stack[from])-1:]...)
		//	stack[from] = stack[from][:len(stack[from])-1]
		//}
		stack[to] = append(stack[to], stack[from][len(stack[from])-no:]...)
		stack[from] = stack[from][:len(stack[from])-no]
	}

	fmt.Printf("Top crates: %v%v%v%v%v%v%v%v%v\n",
		stack[1][len(stack[1])-1:],
		stack[2][len(stack[2])-1:],
		stack[3][len(stack[3])-1:],
		stack[4][len(stack[4])-1:],
		stack[5][len(stack[5])-1:],
		stack[6][len(stack[6])-1:],
		stack[7][len(stack[7])-1:],
		stack[8][len(stack[8])-1:],
		stack[9][len(stack[9])-1:],
	)
}

func parseMoves(str string) (no, from, to int) {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := re.FindStringSubmatch(str)
	no, _ = strconv.Atoi(matches[1])
	from, _ = strconv.Atoi(matches[2])
	to, _ = strconv.Atoi(matches[3])
	//fmt.Printf("Moving: %v from %v to %v\n%v = %v, %v = %v\n", no, from, to, from, stack[from], to, stack[to])

	return
}
