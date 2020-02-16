package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readCompressed(filename string) string {
	data, _ := ioutil.ReadFile(filename)
	return string(data)
}

func uncompress(compressed string) {

	open := make([]int, 0)
	closed := make([]int, 0)
	text := make([]string, 0)
	instructions := make([]string, 0)

	i := 1
	for i < len(compressed) {
		switch compressed[i] {
		case '(':
			if compressed[i-1] != ')' {
				open = append(open, i)
			} else {
				i += strings.Index(compressed[i+1:], ")") + 1
			}
		case ')':
			closed = append(closed, i)
		}
		i++
	}

	// Add text elements inbetween
	for j := 0; j < len(open)-1; j++ {
		text = append(text, compressed[closed[j]+1:open[j+1]])
	}
	text = append(text, compressed[closed[len(closed)-1]+1:])

	//Add instructions
	for j := 0; j < len(open); j++ {
		instructions = append(instructions, compressed[open[j]+1:closed[j]])
	}

	fmt.Println(text, instructions)

	//uncompress
	var builder strings.Builder

	builder.WriteString(compressed[:open[0]])
	for j := 0; j < len(instructions); j++ {
		ops := strings.Split(instructions[j], "x")
		sublen, _ := strconv.Atoi(ops[0])
		n, _ := strconv.Atoi(ops[1])

		for k := 0; k < n; k++ {
			builder.WriteString(text[j][:sublen])
		}
		builder.WriteString(text[j][sublen:])

	}
	fmt.Println(builder.String())
}

func main() {
	compressed := readCompressed("../inputs/day 9.txt")
	uncompress(compressed)
}
