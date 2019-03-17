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
L:
	for {
		fmt.Println("Looking for next step...")
		for _, instruct := range input {
			// fmt.Printf("checking step %v\n", instruct.step)
			if len(instruct.preReqs) == 0 {
				fmt.Printf("completing step %v\n", instruct.step)
				completeSteps += instruct.step
				input = removePreReqs(input, instruct.step)

				input = removeStep(input, instruct.step)
				continue L
			}
		}
		break
	}

	printInstructions(input)
	fmt.Println(completeSteps)
}

func removePreReqs(input []instruction, step string) []instruction {
	for idx, i := range input {
		if i.step != step {
		reqs:
			for {
				fmt.Println(len(input))
				fmt.Println(idx)
				for idxP, p := range input[idx].preReqs {
					if p == step {
						// fmt.Println(i.preReqs)
						i.preReqs = append(i.preReqs[:idxP], i.preReqs[idxP+1:]...)
						// fmt.Println(i.preReqs)
						input[idx].preReqs = i.preReqs
						continue reqs
					}
				}
				break
			}
		}
	}
	return input
}

func removeStep(input []instruction, completeStep string) []instruction {
	// fmt.Printf("Starting Step List\n")
	printInstructions(input)
	for idx, i := range input {
		if i.step == completeStep {
			input = append(input[:idx], input[idx+1:]...)
		}
	}
	// fmt.Printf("New Step List\n")
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
