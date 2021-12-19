package main

import (
	"bufio"
	"os"
	"strings"
)

type Cave struct {
	connections []*Cave
	big         bool
}

func (c *Cave) forEachConnection(next func(cave *Cave) bool, done func()) {
	for _, cave := range c.connections {
		if next(cave) {
			cave.forEachConnection(next, done)
		}
	}

	done()
}

type Path []*Cave

func (p *Path) pop() (cave *Cave) {
	cave = (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return
}

func (p *Path) alreadyVisited(cave *Cave) bool {
	for _, part := range *p {
		if part == cave {
			return true
		}
	}
	return false
}

type CaveSystem map[string]*Cave

func (cs *CaveSystem) countPaths(extra bool) (count int) {
	start := (*cs)["start"]
	end := (*cs)["end"]
	path := Path{start}
	visitedSmallTwice := false
	next := func(cave *Cave) bool {
		switch {
		case cave == end:
			count++
			return false
		case cave == start:
			return false
		case !cave.big && path.alreadyVisited(cave):
			if !extra || visitedSmallTwice {
				return false
			}

			visitedSmallTwice = true
		}

		path = append(path, cave)
		return true
	}
	done := func() {
		cave := path.pop()
		if extra && !cave.big && path.alreadyVisited(cave) {
			visitedSmallTwice = false
		}
	}
	start.forEachConnection(next, done)

	return
}

func newCaveSystem(filename string) *CaveSystem {
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	var connection [2]*Cave
	system := make(CaveSystem)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for i, name := range strings.Split(scanner.Text(), "-") {
			cave, ok := system[name]
			if ok {
				connection[i] = cave
			} else {
				cave = &Cave{
					big: name[0] >= 'A' && name[0] <= 'Z',
				}
				system[name] = cave
				connection[i] = cave
			}
		}

		connection[0].connections = append(connection[0].connections, connection[1])
		connection[1].connections = append(connection[1].connections, connection[0])
	}

	return &system
}

func main() {
	for _, filename := range []string{"example1.txt", "example2.txt", "example3.txt", "input.txt"} {
		println(filename)

		system := newCaveSystem(filename)
		println("\tPart one:", system.countPaths(false))
		println("\tPart two:", system.countPaths(true))
	}
}
