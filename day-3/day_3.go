package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readTrianglesList(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func groupTrianglesVertically(triangles []string) [][]string {
	verticalList := make([][]string, 3)
	verticalList[0] = make([]string, 0)
	verticalList[1] = make([]string, 0)
	verticalList[2] = make([]string, 0)

	for _, row := range triangles {
		fields := strings.Fields(row)

		for i, field := range fields {
			//n, _ := strconv.Atoi(field)
			verticalList[i] = append(verticalList[i], field)
		}
	}
	return verticalList
}

func validateTrianglesHorizontally(triangles []string) (count int) {
	for _, triangle := range triangles {
		sides := strings.Fields(triangle)
		if isvalid := validateTriangle(sides); isvalid {
			count++
		}
	}
	return
}

func validateTrianglesVertically(colums [][]string) (count int) {
	for i := 0; i < len(colums); i++ {
		for j := 0; j < len(colums[i]); j++ {
			if (j+1)%len(colums) == 0 {
				s := make([]string, 3)
				s[0], s[1], s[2] = colums[i][j], colums[i][j-1], colums[i][j-2]
				if isvalid := validateTriangle(s); isvalid {
					count++
				}
			}
		}
	}
	return
}

func validateTriangle(sides []string) bool {
	for i, side := range sides {
		l := sides[:i]
		r := sides[i+1:]
		merged := []string{}
		merged = append(merged, l...)
		merged = append(merged, r...)

		iside, _ := strconv.Atoi(side)
		sumOfSides := sum(merged)
		if iside >= sumOfSides {
			return false
		}
	}
	return true
}

func sum(nums []string) (res int) {
	for _, str := range nums {
		n, _ := strconv.Atoi(str)
		res += n
	}
	return
}

func main() {
	trianglesList := readTrianglesList("../inputs/day 3.txt")
	hcount := validateTrianglesHorizontally(trianglesList)
	fmt.Println("PART I: ", hcount)

	columns := groupTrianglesVertically(trianglesList)
	vcount := validateTrianglesVertically(columns)
	fmt.Println("PART II: ", vcount)

}
