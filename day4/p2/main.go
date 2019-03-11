package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type workLog struct {
	id          int
	dateTime    string
	description string
}

type guard struct {
	id          int
	sleepMatrix []int
}

func main() {
	data := sortDataByTime(data())

	guards := []guard{}
	g := guard{}
	// sleepMatrix := make([]int, 60)

	sMin := 0
	eMin := 0
	for idx, d := range data {
		if idx < 10 {
			fmt.Println(d)
		}
		// when current guard falls asleep
		if d.description == "falls asleep" {
			sm, err := strconv.Atoi(strings.Split(d.dateTime, ":")[1])
			checkErr(err)
			// track the minute they fell asleep
			sMin = sm
			// when current guard wakes up
		} else if d.description == "wakes up" {
			em, err := strconv.Atoi(strings.Split(d.dateTime, ":")[1])
			checkErr(err)
			eMin = em
			// update 'sleepMatrix' for each minute they were asleep
			for i := sMin; i < eMin; i++ {
				g.sleepMatrix[i]++
			}
			// when new guard comes on shift
		} else {
			found := false
			// if the guard is already in the slice
			for _, guard := range guards {
				if guard.id == g.id {
					// append the sleepMatrix with the new values
					for idx := range guard.sleepMatrix {
						guard.sleepMatrix[idx] += g.sleepMatrix[idx]
					}
					found = true
				}
			}

			// if the guard was not in the slice, append as a new guard
			if !found {
				guards = append(guards, g)
			}

			// re-init id and sleepMatrix
			var err error
			newGID := strings.Split(d.description, "#")[1]
			guardID := strings.Split(newGID, " ")[0]
			sleepMatrix := make([]int, 60)
			// init the new guard
			g = guard{}
			g.id, err = strconv.Atoi(guardID)
			checkErr(err)
			g.sleepMatrix = sleepMatrix
		}
	}

	// figure out which guard was asleep the most,
	// then figure out which minute they slept the most
	// across all of their shifts
	guards = sortGuardsByID(guards)

	hID := 0
	hM := 0
	hMins := 0
	for _, g := range guards {
		for i, sleptMins := range g.sleepMatrix {
			if sleptMins > hMins {
				hMins = sleptMins
				hM = i
				hID = g.id
			}
		}
	}

	fmt.Println(hID)
	fmt.Println(hM)
	fmt.Println(hID * hM)
}

func sortDataByTime(data []workLog) []workLog {
	sort.Slice(data, func(i, j int) bool {
		if data[i].dateTime < data[j].dateTime {
			return true
		}
		return false
	})
	return data
}

func sortGuardsByID(guards []guard) []guard {
	sort.Slice(guards, func(i, j int) bool {
		if guards[i].id < guards[j].id {
			return true
		}
		return false
	})
	return guards
}

func buildMatrix(w int, l int) [][]int {
	matrix := make([][]int, w)
	for i := range matrix {
		matrix[i] = make([]int, l)
	}
	return matrix
}

func data() []workLog {

	rawData, err := ioutil.ReadFile("../input.txt")
	checkErr(err)

	data := strings.Split(string(rawData), "\n")
	var parsedData []workLog
	for idx, d := range data {
		tmp := strings.Split(d, "]")[0]
		dateTime := strings.Split(tmp, "[")[1]

		description := strings.Split(d, "] ")[1]

		w := workLog{id: idx, dateTime: dateTime, description: description}
		parsedData = append(parsedData, w)
	}
	return parsedData
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
