package main

import (
	"fmt"
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
	data := sortInput(input)

	for i := range data {
		distances := make([]dist, len(data)-1)
		for j, d2 := range data {
			if len(data) > j+1 {
				distances[j] = calcMnhtnDist(data[i], d2)
			}
		}
		// do calculations at this point using the distances acquired
		// to determine if this set has a greater region than the previous
		// sort the largest region to the top and print
		var ul dist
		var ur dist
		var bl dist
		var br dist

		for _, di := range distances {
			u := false
			r := false
			l := false
			b := false
			if (di.xDist - data[i].x) > 0 {
				r = true
			} else if (di.xDist - data[i].x) < 0 {
				l = true
			}
			if (di.yDist - data[i].y) > 0 {
				u = true
			} else if (di.yDist - data[i].y) < 0 {
				b = true
			}

			var tmpX *int
			var tmpY *int

			if u && r {
				tmpX = &ur.xDist
				tmpY = &ur.yDist
			} else if u && l {
				tmpX = &ul.xDist
				tmpY = &ul.yDist
			} else if b && r {
				tmpX = &br.xDist
				tmpY = &br.yDist
			} else if b && l {
				tmpX = &bl.xDist
				tmpY = &bl.yDist
			}

			if tmpX != nil && tmpY != nil {
				if (*tmpX == 0 && *tmpY == 0) || (*tmpX+*tmpY) > (di.xDist+di.yDist) {
					*tmpX = di.xDist
					*tmpY = di.yDist
				}
			}
			// keep only the distances that have finite areas
			if ul.xDist != 0 && ul.yDist != 0 &&
				ur.xDist != 0 && ur.yDist != 0 &&
				bl.xDist != 0 && bl.yDist != 0 &&
				br.xDist != 0 && br.yDist != 0 {

				fmt.Print("l1")
				fmt.Print(di.l1)
				fmt.Println("")

				fmt.Print("l2")
				fmt.Print(di.l2)
				fmt.Println("")

				fmt.Print("ur")
				fmt.Print(ur)
				fmt.Println("")

				fmt.Print("ul")
				fmt.Print(ul)
				fmt.Println("")

				fmt.Print("br")
				fmt.Print(br)
				fmt.Println("")

				fmt.Print("bl")
				fmt.Print(bl)
				fmt.Println("")
			}
		}
	}
}

func calcMnhtnDist(v1, v2 location) dist {
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
