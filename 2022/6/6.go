package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./2022/6/6.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var byteResult []byte
	for scanner.Scan() {
		byteResult = scanner.Bytes()
	}
	found := false
	markerLen := 14
loopend:
	for i := markerLen - 1; i < len(byteResult)-1; i++ {
		chars := byteResult[i-(markerLen-1) : i+1]
		fmt.Printf("check: %v\n", chars)
		found = false
	QUIT:
		for x := range chars {
			for j := range chars {
				//fmt.Printf("idx %v %v\n", x, j)
				if x == j {
					//fmt.Printf("skip %v %v\n", x, j)
					continue
				}
				//fmt.Printf("compare: %v %v\n", chars[x], chars[j])
				if chars[x] == chars[j] {
					fmt.Printf("matched %v %v\n", chars[x], chars[j])
					found = true
					break QUIT
				}
			}
		}
		if !found {
			fmt.Printf("Marker @ %v\n", i+1)
			break loopend
		}
	}
}
