package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Operation struct {
	operator string
	value    int
}

func (o *Operation) Operate(worry int) int {
	value := o.value
	if value == -1 {
		value = worry
	}
	switch o.operator {
	case "*":
		return worry * value
	case "+":
		return worry + value
	case "/":
		return worry / value
	case "-":
		return worry - value
	}
	return 0
}

type Test struct {
	value      int
	passTarget int
	failTarget int
}

func (t *Test) check(item int) int {
	if item%t.value == 0 {
		return t.passTarget
	}
	return t.failTarget
}

type Monkey struct {
	id        int
	items     []int
	Operation *Operation
	Test      *Test
	inspected int
}

func (m *Monkey) CatchItem(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) ThrowItem() {
	if len(m.items) == 0 {
		return
	}
	var thrown int
	thrown, m.items = m.items[0], m.items[1:]

	target := m.Test.check(thrown)
	Monkeys[target].CatchItem(thrown)
}

func (m *Monkey) Inspect() {
	m.inspected += 1
	if len(m.items) == 0 {
		return
	}
	m.items[0] = m.Operation.Operate(m.items[0])
}

func (m *Monkey) Tire() {
	if len(m.items) == 0 {
		return
	}

	// day 1
	//m.items[0] = int(math.Floor(float64(m.items[0] / 3)))

	// days 2
	if ChecksProduct < m.items[0] {
		m.items[0] = m.items[0] % ChecksProduct
	}
}

var Monkeys = make(map[int]*Monkey)

// day 2
var ChecksProduct = 1

func main() {

	fileContents, err := os.ReadFile("./2022/11/11.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}

	monkeyData := strings.Split(string(fileContents), "\n\n")

	for i := range monkeyData {
		monkey := ParseMonkeyData(monkeyData[i])
		Monkeys[monkey.id] = monkey
		//day 2
		ChecksProduct *= monkey.Test.value
	}

	maxRounds := 10000

	for round := 1; round <= maxRounds; round++ {
		for id := 0; id < len(Monkeys); id++ {
			monkey := Monkeys[id]
			for len(monkey.items) > 0 {
				monkey.Inspect()
				monkey.Tire()
				monkey.ThrowItem()
			}
		}
	}

	business := []int{}
	for i := range Monkeys {
		business = append(business, Monkeys[i].inspected)
		fmt.Printf("Monkey: %v %v %v\n", Monkeys[i].id, Monkeys[i].inspected, Monkeys[i].items)
	}

	sort.Slice(business, func(i, j int) bool {
		return business[i] >= business[j]
	})
	fmt.Printf("Monkey business: %v\n", business[0]*business[1])
}

func ParseMonkeyData(data string) *Monkey {
	regex := regexp.MustCompile(`Monkey (?P<id>\d+):\n.+Starting items: (?P<items>[\d, ]+)\n.+Operation: new = old (?P<operator>[\+\-\*\/]) (?P<opval>.+)\n.+Test: divisible by (?P<test>\d+)\n.+If true: throw to monkey (?P<testpass>\d+)\n.+If false: throw to monkey (?P<testfail>\d+)`)
	results := make(map[string]string)
	matches := regex.FindStringSubmatch(data)

	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" {
			results[name] = matches[i]
		}
	}

	items := strings.Split(results["items"], ",")
	var nums []int
	for j := range items {
		number, _ := strconv.Atoi(strings.TrimSpace(items[j]))
		nums = append(nums, number)
	}

	opval := -1
	if results["opval"] != "old" {
		opval, _ = strconv.Atoi(results["opval"])
	}
	o := &Operation{
		operator: results["operator"],
		value:    opval,
	}
	testVal, _ := strconv.Atoi(results["test"])
	testPass, _ := strconv.Atoi(results["testpass"])
	testFail, _ := strconv.Atoi(results["testfail"])
	t := &Test{
		value:      testVal,
		passTarget: testPass,
		failTarget: testFail,
	}
	id, _ := strconv.Atoi(results["id"])
	m := &Monkey{
		id:        id,
		items:     nums,
		Operation: o,
		Test:      t,
	}

	return m
}
