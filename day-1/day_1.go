package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

const (
	cw = iota
	ccw
)

type vec struct {
	x, y, v int
}

type cell struct {
	x, y int
}

//Rotate Vector 90 CW/CCW
func (v *vec) rotateVec90(direction int) *vec {
	r := new(vec)
	if direction == cw {
		r.x, r.y = -1*v.y, v.x

	} else if direction == ccw {
		r.x, r.y = v.y, v.x*-1
	}
	return r
}

//Get Abs value
func getAbs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

// Read input to fields
func readDirections(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	directions := strings.FieldsFunc(string(data), func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})
	return directions
}

func findDistanceToBunnyHQ(directions []string) {
	x, y := 0, 0
	var firstIntersection *cell = nil
	heading := &vec{x: 0, y: -1, v: 1}
	grid := make(map[cell]int)
	grid[cell{0, 0}] = 0

	for _, v := range directions {

		nextDirection := string(v[0])
		nextDistance := string(v[1:])

		switch nextDirection {
		case "L":
			heading = heading.rotateVec90(ccw)
		case "R":
			heading = heading.rotateVec90(cw)
		}
		heading.v, _ = strconv.Atoi(nextDistance)
		//vectors = append(vectors, heading)
		for i := 1; i <= heading.v; i++ {
			if heading.x != 0 {
				tick := cell{x + heading.x*i, y}
				if _, ok := grid[tick]; !ok {
					grid[tick] = getAbs(x) + getAbs(y)
				} else {
					if firstIntersection == nil {
						firstIntersection = &tick
					}
				}
			} else if heading.y != 0 {
				tick := cell{x, y + heading.y*i}
				if _, ok := grid[tick]; !ok {
					grid[tick] = getAbs(x) + getAbs(y)
				} else {
					if firstIntersection == nil {
						firstIntersection = &tick
					}
				}
			}
		}
		x += heading.v * heading.x
		y += heading.v * heading.y
	}
	fmt.Println("Distance to Bunny HQ ->\n\tPart I: ", getAbs(x)+getAbs(y), "\n\tPart II: ", getAbs(firstIntersection.x)+getAbs(firstIntersection.y))
}

//RunDay1 - Run solution for day 1
func main() {
	directions := readDirections("../inputs/day 1.txt")
	findDistanceToBunnyHQ(directions)

}
