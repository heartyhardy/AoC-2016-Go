package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type bot struct {
	id     uint8
	values []uint8
	done   bool
}

type command struct {
	bot        uint8
	highTo     uint8
	highToType string
	lowTo      uint8
	lowToType  string
	done       bool
}

type output struct {
	id  uint8
	val uint8
}

type botmap map[uint8]*bot
type commandmap map[uint8]*command
type outputmap map[uint8]uint8

func readBotInstructions(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func initBots(instructions []string) (botmap, commandmap) {
	bots := make(botmap)
	commands := make(commandmap)

	for _, instruction := range instructions {
		fields := strings.Fields(instruction)
		if len(fields) < 7 {
			newBot := new(bot)
			v, _ := strconv.ParseUint(fields[1], 10, 8)
			bid, _ := strconv.ParseUint(fields[5], 10, 8)
			newBot.values = make([]uint8, 1)
			newBot.values[0] = uint8(v)
			newBot.id = uint8(bid)

			if _, ok := bots[uint8(bid)]; ok {
				bots[uint8(bid)].values = append(bots[uint8(bid)].values, uint8(v))
				continue
			}
			bots[uint8(bid)] = newBot

		} else {
			newCommand := new(command)
			bid, _ := strconv.ParseUint(fields[1], 10, 8)
			low, _ := strconv.ParseUint(fields[6], 10, 8)
			high, _ := strconv.ParseUint(fields[11], 10, 8)

			newCommand.bot = uint8(bid)
			newCommand.lowTo = uint8(low)
			newCommand.lowToType = fields[5]
			newCommand.highTo = uint8(high)
			newCommand.highToType = fields[10]
			commands[uint8(bid)] = newCommand
		}
	}
	return bots, commands
}

func solve(bots botmap, commands commandmap) {
	outputs := make(outputmap)

	next := findABotWithAPair(bots)

	for next != nil {
		passValue(bots, outputs, bots[next.id], commands[next.id])
		next = findABotWithAPair(bots)
	}

	for _, v := range bots {
		if (v.values[0] == 61 || v.values[0] == 17) &&
			(v.values[1] == 61 || v.values[1] == 17) {
			fmt.Println("Part I: ", v.id)
		}
	}

	fmt.Println("Part II: ", uint32(outputs[0])*uint32(outputs[1])*uint32(outputs[2]))

}

func getLow(values []uint8) uint8 {
	if values[0] < values[1] {
		return values[0]
	}
	return values[1]
}

func getHigh(values []uint8) uint8 {
	if values[0] > values[1] {
		return values[0]
	}
	return values[1]
}

func hasBotGotAPair(botid uint8, bots botmap) bool {
	return len(bots[botid].values) == 2
}

func findABotWithAPair(bots botmap) *bot {
	for _, v := range bots {
		if len(v.values) >= 2 && !v.done {
			return v
		}
	}
	return nil
}

func passValue(bots botmap, outputs outputmap, b *bot, command *command) {
	switch command.highToType {
	case "bot":
		if _, ok := bots[command.highTo]; !ok {
			newBot := new(bot)
			newBot.values = make([]uint8, 1)
			newBot.values[0] = getHigh(b.values)
			newBot.id = command.highTo
			bots[command.highTo] = newBot
		} else if ok {
			bots[command.highTo].values = append(bots[command.highTo].values, getHigh(b.values))
		}
	case "output":
		if _, ok := outputs[command.highTo]; !ok {
			outputs[command.highTo] = getHigh(b.values)
		} else {
			panic("Unintended value add!")
		}
	}

	switch command.lowToType {
	case "bot":
		if _, ok := bots[command.lowTo]; !ok {
			newBot := new(bot)
			newBot.values = make([]uint8, 1)
			newBot.values[0] = getLow(b.values)
			newBot.id = command.lowTo
			bots[command.lowTo] = newBot
		} else if ok {
			bots[command.lowTo].values = append(bots[command.lowTo].values, getLow(b.values))
		}
	case "output":
		if _, ok := outputs[command.lowTo]; !ok {
			outputs[command.lowTo] = getLow(b.values)
		} else {
			panic("Unintended value add!")
		}
	}
	b.done = true
	command.done = true
}

func main() {
	instructions := readBotInstructions("../inputs/day 10.txt")
	bots, commands := initBots(instructions)
	solve(bots, commands)
}
