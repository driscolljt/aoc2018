package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type instruction struct {
	id      int
	text    string
	preReqs []string
	step    string
}

type work struct {
	id   int
	time int
	step string
}

func main() {
	input := data()
	steps := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// for each step in steps, add missing steps with empty requirements
	for _, r := range []rune(steps) {
		step := string(r)
		found := false
		for _, s := range input {
			if s.step == step {
				found = true
			}
		}
		if !found {
			input = append(input, instruction{id: len(input), step: step})
		}
	}

	input = sortDataByStep(input)

	completeSteps := ""

	// 5 workers
	// 60 seconds + 1-26 based on letter (a = 1, b = 2, ...)
	// how long will it take to complete the instructions?
	workers := []work{work{id: 0, step: "", time: 0}, work{id: 1, step: "", time: 0}, work{id: 2, step: "", time: 0}, work{id: 3, step: "", time: 0}, work{id: 4, step: "", time: 0}}
	timer := 0

work:
	for idxW, worker := range workers {
		fmt.Printf("Current worker %v\n", worker)
		if worker.time == 0 {
			if worker.step != "" {
				completeSteps += worker.step
				// fmt.Println(len(input))
				if len(input) == 0 {
					break work
				}
				worker.step = ""
			}
		}
		if worker.step == "" {
			for idxI := range input {
				if len(input[idxI].preReqs) == 0 {
					fmt.Println("initializing worker with " + input[idxI].step)
					workers[idxW] = work{step: input[idxI].step, time: (60 + idxW + 1)}
					input = removePreReqs(input, workers[idxW].step)
					input = removeStep(input, workers[idxW].step)
					break
				}
			}
		}
		if worker.time > 0 {
			// fmt.Println("removing second from worker timer...")
			worker.time = worker.time - 1
		}
		if worker.id == 4 {
			timer = timer + 1
		}
		continue work
	}
	// fmt.Println("incrementing timer...")
	// increment timer?
}

// fmt.Printf("Completed in %v seconds\n", timer)

func removePreReqs(input []instruction, step string) []instruction {
	for idxI := range input {
		if input[idxI].step != step {
		reqs:
			for {
				// fmt.Println(len(input))
				// fmt.Println(idx)
				for idxP := range input[idxI].preReqs {
					if input[idxI].preReqs[idxP] == step {
						// fmt.Println(i.preReqs)
						input[idxI].preReqs = append(input[idxI].preReqs[:idxP], input[idxI].preReqs[idxP+1:]...)
						continue reqs
					}
				}
				break
			}
		}
	}
	fmt.Println("Removed from PreReqs " + step)
	printInstructions(input)
	return input
}

func removeStep(input []instruction, step string) []instruction {
	for idx := range input {
		// fmt.Println("Comparing input[idx].step " + input[idx].step)
		// fmt.Println("To step " + step)
		if input[idx].step == step {
			input = append(input[:idx], input[idx+1:]...)
			break
			// printInstructions(input)
		}
	}
	// fmt.Println("Removed from steps " + step)
	// printInstructions(input)
	return input
}

func sortDataByStep(instructs []instruction) []instruction {
	sort.Slice(instructs, func(i, j int) bool {
		return instructs[i].step < instructs[j].step
	})
	return instructs
}

func rmvFromSlice(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func printInstructions(instructs []instruction) {
	for _, i := range instructs {
		fmt.Printf("Step: '%v', PreReqs: '%v'\n", i.step, i.preReqs)
	}
}

func data() []instruction {
	rawData, err := ioutil.ReadFile("../input.txt")
	checkErr(err)
	data := string(rawData)

	ret := make([]instruction, 0)
	for i, s := range strings.Split(data, "\n") {
		id := i
		preReq := strings.Split(s, " ")[1]
		step := strings.Split(s, " ")[7]

		exists := false
		for idx, r := range ret {
			if r.step == step {
				r.preReqs = append(r.preReqs, preReq)
				ret[idx] = r
				exists = true
				break
			}
		}
		if !exists {
			instruct := instruction{id: id, text: s, preReqs: []string{preReq}, step: step}
			ret = append(ret, instruct)
		}
	}
	return ret
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
