package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func readRegisterInstructions(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func performOperations(instructions []string, overrideC bool) {
	registers := make(map[string]int)

	if overrideC {
		registers["c"] = 1
	}

	var i int
	for i < len(instructions) {
		instruction := instructions[i]
		params := strings.Fields(instruction)

		switch params[0] {
		case "cpy":
			v, err := strconv.Atoi(params[1])
			if err != nil {
				registers[params[2]] = registers[params[1]]
			} else {
				registers[params[2]] = v
			}
		case "inc":
			if _, ok := registers[params[1]]; !ok {
				registers[params[1]]++
			} else {
				registers[params[1]]++
			}
		case "dec":
			if _, ok := registers[params[1]]; !ok {
				registers[params[1]]--
			} else {
				registers[params[1]]--
			}
		case "jnz":
			v, _ := strconv.Atoi(params[1])
			if (unicode.IsLetter(rune(params[1][0])) && registers[params[1]] != 0) ||
				(unicode.IsNumber(rune(params[1][0])) && v != 0) {
				skip, _ := strconv.Atoi(params[2])
				i += skip
				continue
			}
		}
		i++
	}

	for k, v := range registers {
		fmt.Println(k, v)
	}
}

func main() {
	instructions := readRegisterInstructions("../inputs/day 12.txt")
	performOperations(instructions, true)
}
