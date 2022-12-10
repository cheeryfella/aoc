package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	Name   string
	Parent *Dir
	Dirs   map[string]Dir
	Files  map[string]File
}

func (d *Dir) CD(dname string) (*Dir, error) {
	if dname == ".." {
		if d.Parent == nil {
			return nil, errors.New("no parent")
		}
		return d.Parent, nil
	}
	if _, ok := d.Dirs[dname]; !ok {
		return nil, errors.New(" no such file or directory: " + dname)
	}
	dir := d.Dirs[dname]
	return &dir, nil
}

func (d *Dir) AddFile(name string, size int) {
	d.Files[name] = File{
		Name: name,
		Size: size,
	}
}

func (d *Dir) GetSize() int {
	size := 0
	for i := range d.Files {
		size += d.Files[i].Size
	}
	for i := range d.Dirs {
		dir := d.Dirs[i]
		size += dir.GetSize()
	}
	return size
}

type File struct {
	Name string
	Size int
}

const (
	DiskSpace         = 70000000
	RequiredFreeSpace = 30000000
)

var Dirs []Dir

var FS map[string]Dir

func main() {
	f, err := os.Open("./2022/7/7.txt")
	if err != nil {
		fmt.Print(err)
		panic("unable to open file")
	}
	defer f.Close()

	var Dirs []Dir
	fs := make(map[string]Dir, 1)

	fs["/"] = Dir{
		Name:  "/",
		Dirs:  map[string]Dir{},
		Files: map[string]File{},
	}

	current := fs["/"]
	scanner := bufio.NewScanner(f)
exitcode:
	for scanner.Scan() {
		res := scanner.Text()
		cmmd := strings.Split(res, " ")
		if string(cmmd[0]) == "$" {
			switch cmmd[1] {
			case "cd":
				if cmmd[2] == "/" {
					current = fs["/"]
					continue
				}

				target, err := current.CD(cmmd[2])

				if err != nil {
					fmt.Printf("Error: %v\n", err)
					break exitcode
				}
				current = *target
				continue
			case "ls":
				continue
			}
		}
		if string(cmmd[0]) == "dir" {
			if _, ok := current.Dirs[cmmd[1]]; ok {
				//errors.New(cmmd[1] + ": File exists")
				fmt.Printf("Error: %v\n", errors.New(cmmd[1]+": File (dir) exists"))
				break exitcode
			}
			parent := current
			nd := Dir{
				Name:   cmmd[1],
				Parent: &parent,
				Dirs:   map[string]Dir{},
				Files:  map[string]File{},
			}
			current.Dirs[cmmd[1]] = nd
			Dirs = append(Dirs, nd)
			continue
		}
		if _, ok := current.Files[cmmd[1]]; ok {
			fmt.Printf("Error: %v\n", errors.New(cmmd[1]+": File exists"))
		}
		fsize, _ := strconv.Atoi(cmmd[0])
		nfile := File{
			Name: cmmd[1],
			Size: fsize,
		}
		current.Files[cmmd[1]] = nfile
	}
	dfg := fs["/"]
	sizeTotal := 0
	for i := range Dirs {
		td := Dirs[i]
		size := td.GetSize()
		//fmt.Printf("Dir: %v Size: %v\n", td.Name, size)
		if size <= 100000 {
			sizeTotal += size
		}
	}
	fmt.Printf("sub 100000 size:\n%+v\n", sizeTotal)
	fmt.Printf("FS size:\n%+v\n", dfg.GetSize())
	free := DiskSpace - dfg.GetSize()
	fmt.Printf("FREE size:\n%+v\n", free)
	needed := RequiredFreeSpace - free
	fmt.Printf("Needed:\n%+v\n", needed)
	candidates := []Dir{}
	for i := range Dirs {
		td := Dirs[i]
		size := td.GetSize()
		//fmt.Printf("Dir: %v Size: %v\n", td.Name, size)
		if size >= needed {
			candidates = append(candidates, td)
		}
	}
	for i := range candidates {
		c := candidates[i]
		fmt.Printf("Size: %v\tCandidate: %v\n", c.GetSize(), c.Name)
	}
}
