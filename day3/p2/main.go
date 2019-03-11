package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type claim struct {
	id   int64
	left int64
	top  int64
	w    int64
	h    int64
}

func main() {
	claims := data()
	matrix := buildMatrix(1000, 1000)

	for _, c := range claims {
		applyClaim(matrix, c)
	}

	count := 0
	for _, x := range matrix {
		for _, y := range x {
			if y > 1 {
				count++
			}
		}
	}

	for _, c := range claims {
		if checkClaim(matrix, c) {
			fmt.Println(c.id)
		}
	}
}

func checkClaim(matrix [][]int, c claim) bool {
	sX := int(c.left)
	sY := int(c.top)
	for x := 0; x < int(c.w); x++ {
		for y := 0; y < int(c.h); y++ {
			if matrix[sX+x][sY+y] > 1 {
				return false
			}
		}
	}
	return true
}

func applyClaim(matrix [][]int, c claim) {
	sX := int(c.left)
	sY := int(c.top)
	for x := 0; x < int(c.w); x++ {
		for y := 0; y < int(c.h); y++ {
			matrix[sX+x][sY+y]++
		}
	}
}

func buildMatrix(w int, l int) [][]int {
	matrix := make([][]int, w)
	for i := range matrix {
		matrix[i] = make([]int, l)
	}
	return matrix
}

func data() []claim {

	rawData, err := ioutil.ReadFile("../input.txt")
	checkErr(err)

	data := strings.Split(string(rawData), "\n")
	var parsedData []claim
	for _, d := range data {
		tmp := strings.Split(d, "#")[1]
		sID := strings.Split(tmp, " ")[0]
		id, _ := strconv.ParseInt(sID, 0, 64)

		tmp = strings.Split(d, ",")[0]
		sLeft := strings.Split(tmp, "@ ")[1]
		left, _ := strconv.ParseInt(sLeft, 0, 64)

		tmp = strings.Split(d, ":")[0]
		sTop := strings.Split(tmp, ",")[1]
		top, _ := strconv.ParseInt(sTop, 0, 64)

		tmp = strings.Split(d, "x")[0]
		sW := strings.Split(tmp, ": ")[1]
		w, _ := strconv.ParseInt(sW, 0, 64)

		tmp = strings.Split(d, "x")[1]
		h, _ := strconv.ParseInt(tmp, 0, 64)

		c := claim{id: id, left: left, top: top, w: w, h: h}
		parsedData = append(parsedData, c)
	}
	return parsedData
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
