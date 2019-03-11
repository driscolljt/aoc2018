package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type location struct {
	id int
	x  int
	y  int
}

func main() {
	input := data()
	data := sortInput(input)
	matrix := plotInput(data)

	fmt.Println(matrix)
}

func plotInput(locations []location) [][]string {
	gridMatrix := buildMatrix(500, 500)
	for _, l := range locations {
		gridMatrix[l.x][l.y] = strconv.Itoa(l.id)
		// fmt.Println(gridMatrix[l.x][l.y])
	}
	return gridMatrix
}

func buildMatrix(w int, l int) [][]string {
	matrix := make([][]string, w)
	for i := range matrix {
		matrix[i] = make([]string, l)
		for j := range matrix[i] {
			matrix[i][j] = "x"
		}
	}
	return matrix
}

func sortInput(locations []location) []location {
	sort.Slice(locations, func(i, j int) bool {
		if locations[i].x > locations[j].x {
			return false
		} else if locations[i].y > locations[j].y {
			return false
		}
		return true
	})
	// fmt.Println(locations)
	return locations
}

func data() []location {

	rawData, err := ioutil.ReadFile("../input.txt")
	checkErr(err)
	data := string(rawData)

	ret := make([]location, 50)
	for i, s := range strings.Split(data, "\n") {
		id := i
		x, _ := strconv.Atoi(strings.Split(s, ", ")[0])
		y, _ := strconv.Atoi(strings.Split(s, ", ")[1])
		l := location{id: id, x: x, y: y}
		ret[i] = l
	}
	return ret
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
