package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := data()

	chars := "abcdefghijklmnopqrstuvwxyz"
	runes := []rune(chars)
	shortest := 1000000
	for _, r := range runes {
		ch := string(r)
		newInput := strings.Replace(input, ch, "", -1)
		newInput = strings.Replace(newInput, strings.ToUpper(ch), "", -1)

		l := len(reactPolymers(newInput))
		if l < shortest {
			shortest = l
		}
	}
	fmt.Println(shortest)
}

func reactPolymers(data string) string {
L:
	for {
		runes := []rune(data)
		for i, r := range runes {
			if len(runes) > i+1 {
				a := string(r)
				b := string(runes[i+1])
				if strings.EqualFold(a, b) {
					if a != b {
						data = strings.Replace(data, a+b, "", -1)
						// if we update the string, we need to start over the loop
						continue L
					}
				}
			}
		}
		break
	}
	return data
}

func data() string {

	rawData, err := ioutil.ReadFile("../input.txt")
	checkErr(err)
	return string(rawData)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
