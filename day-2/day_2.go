package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	keyPadI = iota
	keyPadII
)

type direction struct {
	x, y int
}

var keycodeI [][]int = [][]int{
	[]int{1, 2, 3},
	[]int{4, 5, 6},
	[]int{7, 8, 9},
}
var keycodeII [][]int = [][]int{
	[]int{0, 0, 1, 0, 0},
	[]int{0, 2, 3, 4, 0},
	[]int{5, 6, 7, 8, 9},
	[]int{0, 10, 11, 12, 0},
	[]int{0, 0, 13, 0, 0},
}
var directions []direction = []direction{
	direction{0, -1}, //U
	direction{1, 0},  //R
	direction{0, 1},  //D
	direction{-1, 0}, //L
}

func readCodes(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func enterKeyCode(codes []string, keypadType int) {

	var x, y int
	if keypadType == keyPadI {
		x, y = 1, 1
	} else if keypadType == keyPadII {
		x, y = 0, 2
	}

	for _, code := range codes {
		dx, dy := 0, 0
		for _, codeRune := range code {
			switch codeRune {
			//U
			case 85:
				dx = x + directions[0].x
				dy = y + directions[0].y
			//R
			case 82:
				dx = x + directions[1].x
				dy = y + directions[1].y
			//D
			case 68:
				dx = x + directions[2].x
				dy = y + directions[2].y
			//L
			case 76:
				dx = x + directions[3].x
				dy = y + directions[3].y
			}

			if (dx >= 0 && dx < 3 && dy >= 0 && dy < 3) && keypadType == keyPadI {
				x, y = dx, dy
			} else if (dx >= 0 && dx < 5 && dy >= 0 && dy < 5) && keypadType == keyPadII {
				if keycodeII[dy][dx] != 0 {
					x, y = dx, dy
				} else {
					continue
				}
			}
		}
		if keypadType == keyPadI {
			fmt.Printf("%v", keycodeI[y][x])
		} else if keypadType == keyPadII {
			switch keycodeII[y][x] {
			case 10:
				fmt.Printf("%v", "A")
			case 11:
				fmt.Printf("%v", "B")
			case 12:
				fmt.Printf("%v", "C")
			case 13:
				fmt.Printf("%v", "D")
			default:
				fmt.Printf("%v", keycodeII[y][x])
			}
		}
	}
}

func main() {
	codes := readCodes("../inputs/day 2.txt")
	fmt.Println("\nPart I")
	enterKeyCode(codes, keyPadI)
	fmt.Println("\nPart II")
	enterKeyCode(codes, keyPadII)
	fmt.Println()
}
