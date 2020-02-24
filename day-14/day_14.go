package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

type md5hash struct {
	index int
	hash  string
}

func genHashes(base string) []*md5hash {

	hashes := make([]*md5hash, 0)
	index := 0

	for index < 20000 {
		var builder strings.Builder
		builder.WriteString(base)
		builder.WriteString(strconv.Itoa(index))
		data := []byte(builder.String())

		hash := md5.Sum(data)

		newhash := new(md5hash)
		newhash.index = index
		newhash.hash = fmt.Sprintf("%x", hash)

		hashes = append(hashes, newhash)
		index++
	}
	return hashes
}

func isTriplet(hash string) bool {
	for i := 0; i < len(hash)-3; i++ {
		var builder strings.Builder
		for j := 0; j < 3; j++ {
			builder.WriteByte(hash[i])
		}

		if strings.Count(hash, builder.String()) >= 1 {
			return true
		}
	}
	return false
}

func isFiver(hash string) bool {
	for i := 0; i < len(hash)-3; i++ {
		var builder strings.Builder
		for j := 0; j < 5; j++ {
			builder.WriteByte(hash[i])
		}

		if strings.Count(hash, builder.String()) >= 1 {
			return true
		}
	}
	return false
}

func main() {
	keys := make([]*md5hash, 0)
	hashes := genHashes("abc")
	for k, v := range hashes {
		if isTriplet(v.hash) {
			r := k + 1000
			if k+1000 > len(hashes)-1 {
				r = k + (len(hashes) - k)
			}
			for i := k + 1; i < r; i++ {
				if isFiver(hashes[i].hash) {
					keys = append(keys, v)
					break
				}
			}
		}
	}

	for k, v := range keys {
		fmt.Printf("%-5v %5v %50v\n", k, v.index, v.hash)
	}
}
