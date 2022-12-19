package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type List []interface{}

func main() {
	fileContents, err := os.ReadFile("./2022/13/13.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}

	data := strings.Split(string(fileContents), "\n\n")

	var oks int
	for i := range data {
		packets := strings.Split(strings.Trim(data[i], "\n"), "\n")

		if len(packets) != 2 {
			panic("incorrect packet count")
		}

		left := BuildPackett(packets[0])
		right := BuildPackett(packets[1])

		res := PacketValidator(left, right)

		if res != 0 {
			oks += i + 1
		}
	}
	fmt.Printf("Sum : %v\n", oks)

	//day 2
	newPackets := strings.Replace(strings.Trim(string(fileContents),"\n"), "\n\n", "\n",-1)
	newData := strings.Split(newPackets, "\n")
	newData = append(newData, "[[2]]","[[6]]")
	packets := make([]List,len(newData))
	for j:=0;j<len(newData);j++{
		if newData[j] == "\n" {
			continue
		}
		packets[j] = BuildPackett(newData[j])
	}

	sort.Slice(packets, func(l, r int) bool {
		return PacketValidator(packets[l],packets[r]) > 0
	})

	index := []int{}
	for x := 0 ; x<len(packets); x++ {
		tmp,_ := json.Marshal(packets[x])
		if string(tmp) == "[[2]]" || string(tmp) == "[[6]]"{
			index = append(index,x+1)
		}
	}
	fmt.Printf("%v\nResult: %v\n",index, index[0]*index[1])
}

func BuildPackett(packet string) List {
	var p List
	err := json.Unmarshal([]byte(packet), &p)
	if err != nil {
		fmt.Printf("err: %e\nstring:\n%s \n", err, packet)
	}
	return p
}

func PacketValidator(l, r []interface{}) int {
	lSize := len(l)
	rSize := len(r)
	max := int(math.Max(float64(lSize), float64(rSize)))

	for i := 0; i < max; i++ {
		if i >= lSize {
			return 1
		}
		if i >= rSize {
			return 0
		}
		lVal, lok := l[i].(float64)
		rVal, rok := r[i].(float64)
		if lok && rok {
			if lVal < rVal {
				return 1
			}
			if lVal > rVal {
				return 0
			}
			if len(l)-1 == i && len(r)-1 > i {
				return 1
			}
			continue
		}
		if lok {
			newL := List{lVal}
			nr, _ := r[i].([]interface{})
			if res := PacketValidator(newL, nr); res != 2 {
				return res
			}
			continue
		}
		if rok {
			newR := List{rVal}
			nl, _ := l[i].([]interface{})
			if res := PacketValidator(nl, newR); res != 2 {
				return res
			}
			continue
		}

		subL, lok := l[i].([]interface{})
		subR, rok := r[i].([]interface{})

		if !lok || !rok {
			fmt.Printf("Error l: %v r: %v\n", lok, rok)
			continue
		}
		res := PacketValidator(subL, subR)
		if res != 2 {
			return res
		}
	}
	return 2
}
