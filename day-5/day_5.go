package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func genPassI(hash string) {
	var password strings.Builder
	index, count := 4515059, 0
	for count < 8 {
		bhash := []byte(fmt.Sprint(hash, index))
		newMD5 := fmt.Sprintf("%x", md5.Sum(bhash))
		if isuffix, err := strconv.Atoi(newMD5[:5]); err == nil && isuffix == 0 {
			password.WriteString(newMD5[5:6])
			fmt.Printf("%v Occurrence is: %v and password char: %v\n", index, newMD5, newMD5[5:6])
			count++
		}
		index++
	}
	fmt.Printf("\n%v\n", password.String())
}

func genPassII(hash string) {
	password := make([]string, 8)
	index := 4515059
	for hasEmptySlots(password) {
		bhash := []byte(fmt.Sprint(hash, index))
		newMD5 := fmt.Sprintf("%x", md5.Sum(bhash))
		if isuffix, err := strconv.Atoi(newMD5[:5]); err == nil && isuffix == 0 {
			pos, isposerr := strconv.Atoi(newMD5[5:6])
			val := newMD5[6:7]
			if isposerr == nil && pos <= 7 && pos >= 0 {
				if password[pos] != "" {
					index++
					continue
				}
				fmt.Printf("%v Occurrence is: %v and password char: %v\n", index, newMD5, newMD5[5:6])
				fmt.Println(pos, val, password)
				password[pos] = val
			}
		}
		index++
	}
	fmt.Printf("\n%v\n", strings.Join(password, ""))
}

func hasEmptySlots(password []string) bool {
	for _, v := range password {
		if v == "" {
			return true
		}
	}
	return false
}

func main() {
	genPassI("uqwqemis")
	genPassII("uqwqemis")
}
