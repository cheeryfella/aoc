package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var priority = map[string]int{
	"?": 0,
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func main() {
	f, err := os.Open("./2022/3input/3.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	score := 0
	idx := 0
	res := []string{"", "", ""}
	for scanner.Scan() {

		res[idx] = scanner.Text()
		//	l := len(res) / 2
		//	pt1 := res[0:l]
		//	pt2 := res[l:]
		//	//fmt.Println(res)
		//quit:
		//	for i := range pt1 {
		//		if strings.Contains(pt2, string(pt1[i])) {
		//			fmt.Printf("found: %v\nscore: %v\n", string(pt1[i]), priority[string(pt1[i])])
		//			score += priority[string(pt1[i])]
		//			break quit
		//		}
		//	}

		if idx == 2 {
			idx = 0
			badge := findCommon(res)
			fmt.Printf("Badge: %v\n", badge)
			score += priority[badge]
			continue
		}
		idx += 1
	}

	fmt.Printf("Priority: %v\n", score)
}

func findCommon(strs []string) string {
	base := strs[0]
	for i := range base {
		if strings.Contains(strs[1], string(base[i])) && strings.Contains(strs[2], string(base[i])) {
			return string(base[i])
		}
	}
	return "?"
}
