package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type char struct {
	ch    rune
	count int
}

func readDecodeInstructions(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func decodeAndValidate(instructions []string) {
	sum, northpoleSector := 0, 0
	for _, line := range instructions {
		sector, ok := getValidChecksumSector(line)
		sum += sector
		if ok {
			fmt.Println(decrypt(line, sector))
			if strings.Contains(decrypt(line, sector), "northpole object") {
				northpoleSector = sector
			}
		}
	}
	fmt.Println(sum, northpoleSector)
}

func getValidChecksumSector(line string) (int, bool) {

	charcount := make(map[rune]*char)
	lastDash := strings.LastIndex(line, "-")
	codeOnly := strings.Trim(line[:lastDash], " ")
	sector, _ := strconv.Atoi(line[lastDash+1 : strings.LastIndex(line, "[")])
	checksum := line[strings.Index(line, "[")+1 : strings.LastIndex(line, "]")]

	for _, c := range codeOnly {
		if unicode.IsLetter(c) {
			if _, ok := charcount[c]; !ok {
				charcount[c] = &char{c, 1}
			} else {
				charcount[c].count++
			}
		}
	}

	if isARealRoom(charcount, checksum) {
		return sector, true
	}
	return 0, false
}

func isARealRoom(charcount map[rune]*char, checksum string) bool {
	cmpCheckSum := ""
	occurrences := make([]char, 0)
	for _, v := range charcount {
		occurrences = append(occurrences, *v)
	}
	sort.Slice(occurrences, func(i, j int) bool {
		if occurrences[i].count > occurrences[j].count {
			return true
		}
		if occurrences[i].count < occurrences[j].count {
			return false
		}
		return occurrences[i].ch < occurrences[j].ch
	})

	for i := 0; i < len(checksum); i++ {
		cmpCheckSum += string(occurrences[i].ch)
	}

	return checksum == cmpCheckSum
}

func shiftRuneByN(ch rune, n int) rune {
	shifted := ch + rune(n%26)
	if shifted > 122 {
		return shifted - 26
	} else if shifted < 97 {
		return shifted + 26
	}
	return shifted
}

func decrypt(line string, sector int) string {
	decrypted := make([]rune, 0)
	for _, ch := range line {
		if unicode.IsLetter(ch) {
			decrypted = append(decrypted, shiftRuneByN(ch, sector))
		} else if unicode.IsPunct(ch) {
			decrypted = append(decrypted, ' ')
		}
	}
	return string(decrypted)
}

func main() {
	instructions := readDecodeInstructions("../inputs/day 4.txt")
	decodeAndValidate(instructions)
}
