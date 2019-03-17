package main

import (
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type location struct {
	id int
	x  int
	y  int
}

type dist struct {
	l1    int
	l2    int
	xDist int
	yDist int
}

func main() {
	input := data()
	input = sortInput(input)

	for i := range input {
		distances := make([]dist, len(input)-1)
		for j, d2 := range input {
			if len(input) > j+1 {
				distances[j] = calcDistance(input[i], d2)
			}
		}
		// at this point in the loop
		// i have the distances from the current
		// data point to every other point

		// if the current point has an infinite area,
		// continue the loop

		// otherwise, determine how many points are closer
		// to the current position than any other

		// keep only the highest and when the loop is complete,
		// print the point details and the area (including the point itself)
	}
}

func calcDistance(v1, v2 location) dist {
	dX := math.Abs(float64(v1.x) - float64(v2.x))
	dY := math.Abs(float64(v1.y) - float64(v2.y))

	ret := dist{xDist: int(dX), yDist: int(dY), l1: v1.id, l2: v2.id}
	// fmt.Println(ret)
	return ret
}

func sortDist(distances []dist) []dist {
	sort.Slice(distances, func(i, j int) bool {
		if distances[i].xDist > distances[j].xDist {
			return false
		}
		if distances[i].yDist > distances[j].yDist {
			return false
		}
		return true
	})
	return distances
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
