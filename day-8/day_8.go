package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const LCD_WIDTH = 50
const LCD_HEIGHT = 6

const (
	rotate = "rotate"
	rect   = "rect"
)

func readTinlyLCDInstructions(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func initTinyLCD() [][]int {
	tinyLCD := make([][]int, LCD_HEIGHT)

	for r, rows := range tinyLCD {
		tinyLCD[r] = make([]int, LCD_WIDTH)
		for c := range rows {
			tinyLCD[r][c] = 0
		}
	}
	return tinyLCD
}

func onInstructionReceived(instructions []string, tinyLCD [][]int) {

	for _, instruction := range instructions {
		fields := strings.Fields(instruction)

		switch fields[0] {
		case rect:
			ops := strings.Split(fields[1], "x")
			column, _ := strconv.Atoi(ops[0])
			row, _ := strconv.Atoi(ops[1])

			for r := 0; r < row; r++ {
				for c := 0; c < column; c++ {
					tinyLCD[r][c] = 1
				}
			}
		case rotate:
			switch fields[1] {
			case "column":
				c, _ := strconv.Atoi(fields[2][2:])
				offset, _ := strconv.Atoi(fields[4])
				column := copyColumn(tinyLCD, c)

				resetColumn(tinyLCD, c)

				for r := 0; r < LCD_HEIGHT; r++ {
					nextIndex := (r + offset) % LCD_HEIGHT
					tinyLCD[nextIndex][c] = column[r]
				}
			case "row":
				r, _ := strconv.Atoi(fields[2][2:])
				offset, _ := strconv.Atoi(fields[4])
				row := copyRow(tinyLCD, r)

				resetRow(tinyLCD, r)

				for c := 0; c < LCD_WIDTH; c++ {
					nextIndex := (c + offset) % LCD_WIDTH
					tinyLCD[r][nextIndex] = row[c]
				}
			}
		}
	}
}

func copyColumn(tinyLCD [][]int, c int) []int {
	col := make([]int, LCD_HEIGHT)
	for r := 0; r < LCD_HEIGHT; r++ {
		col[r] = tinyLCD[r][c]
	}
	return col
}

func copyRow(tinyLCD [][]int, r int) []int {
	row := make([]int, LCD_WIDTH)
	for c := 0; c < LCD_WIDTH; c++ {
		row[c] = tinyLCD[r][c]
	}
	return row
}

func resetColumn(tinyLCD [][]int, c int) {
	for r := 0; r < LCD_HEIGHT; r++ {
		tinyLCD[r][c] = 0
	}
}

func resetRow(tinyLCD [][]int, r int) {
	for c := 0; c < LCD_WIDTH; c++ {
		tinyLCD[r][c] = 0
	}
}

func display(tinyLCD [][]int) {
	for _, rows := range tinyLCD {
		for _, col := range rows {
			switch col {
			case 0:
				fmt.Printf("\u001b[38;5;18m\u001b[48;5;17m%2v\u001b[0m", "◻")
			case 1:
				fmt.Printf("\u001b[38;5;196m\u001b[48;5;202m%2v\u001b[0m", "◼")
			}
		}
		fmt.Println()
	}
}

func main() {
	instructions := readTinlyLCDInstructions("../inputs/day 8.txt")
	lcd := initTinyLCD()
	onInstructionReceived(instructions, lcd)
	fmt.Println()
	display(lcd)
	fmt.Println()
}
