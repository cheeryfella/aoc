package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tree struct {
	Height int
	Up     *Tree
	Down   *Tree
	Left   *Tree
	Right  *Tree
}

func (t *Tree) getTarget(direction string) *Tree {
	switch direction {
	case "up":
		return t.Up
	case "down":
		return t.Down
	case "left":
		return t.Left
	case "right":
		return t.Right
	}
	return &Tree{}
}

func (t *Tree) Visible() bool {
	up := t.visible("up")
	down := t.visible("down")
	left := t.visible("left")
	right := t.visible("right")

	return up || down || left || right
}

func (t *Tree) visible(direction string) bool {
	if t.height(direction, t.Height) >= t.Height {
		return false
	}
	return true
}

func (t *Tree) height(direction string, height int) int {
	target := t.getTarget(direction)
	if target == nil {
		return -1
	}
	if target.Height >= height {
		return target.Height
	}
	return target.height(direction, height)
}

func (t *Tree) View() int {
	up := t.distance("up", t.Height)
	down := t.distance("down", t.Height)
	left := t.distance("left", t.Height)
	right := t.distance("right", t.Height)

	return up * down * left * right
}

func (t *Tree) distance(direction string, height int) int {
	target := t.getTarget(direction)
	if target == nil {
		return 0
	}
	if target.Height >= height {
		return 1
	}
	return 1 + target.distance(direction, height)
}

func main() {
	f, err := os.Open("./2022/8/8.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()
	grid := map[int]map[int]*Tree{}
	scanner := bufio.NewScanner(f)
	row := 0
	for scanner.Scan() {
		res := scanner.Text()
		grid[row] = map[int]*Tree{}
		for i := range res {
			height, _ := strconv.Atoi(string(res[i]))
			t := Tree{
				Height: height,
			}
			if _, ok := grid[row-1][i]; ok {
				up := grid[row-1][i]
				t.Up = up
				up.Down = &t
			}
			if _, ok := grid[row][i-1]; ok {
				left := grid[row][i-1]
				t.Left = left
				left.Right = &t
			}

			grid[row][i] = &t
		}
		row += 1
	}

	visible := 0
	distance := 0
	for i := 0; i < len(grid); i++ {
		for x := 0; x < len(grid[i]); x++ {
			t := grid[i][x]
			if t.Visible() {
				visible += 1
			}

			d := t.View()
			if d > distance {
				distance = d
			}
		}
	}

	fmt.Printf("Visble %v\n", visible)
	fmt.Printf("Distance %v\n", distance)
}
