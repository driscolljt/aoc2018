package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := strings.Split(data(), "\r\n")
	var freq int64
	var freqList []int64
	found := false

	for !found {
		freq, freqList, found = processData(freq, freqList, data)
	}
	fmt.Println(freq)

}

func processData(freq int64, freqList []int64, data []string) (int64, []int64, bool) {

	found := false

	for d := range data {
		adj, _ := strconv.ParseInt(data[d], 0, 64)
		freq += adj
		if contains(freqList, freq) {
			found = true
			break
		}
		freqList = append(freqList, freq)
	}

	return freq, freqList, found
}

func data() string {
	return "+16\r\n-2\r\n-5\r\n+1\r\n-12\r\n-1\r\n-8\r\n+12\r\n+5\r\n-18\r\n-12\r\n-11\r\n-1\r\n+2\r\n-8\r\n-6\r\n-1\r\n-17\r\n-18\r\n+16\r\n-4\r\n+18\r\n+7\r\n-4\r\n+13\r\n+11\r\n-5\r\n+11\r\n-8\r\n+15\r\n+8\r\n-16\r\n+15\r\n+10\r\n-19\r\n-17\r\n+5\r\n-13\r\n+10\r\n+14\r\n+16\r\n+8\r\n+14\r\n-8\r\n+19\r\n-17\r\n+14\r\n+12\r\n+8\r\n+9\r\n-11\r\n+1\r\n+2\r\n-16\r\n+19\r\n+13\r\n-5\r\n+6\r\n-10\r\n-6\r\n+4\r\n-15\r\n+20\r\n+18\r\n-14\r\n+15\r\n+3\r\n+4\r\n-3\r\n-9\r\n+1\r\n+18\r\n+8\r\n+2\r\n-11\r\n+19\r\n-3\r\n-10\r\n+14\r\n+12\r\n+6\r\n-17\r\n-14\r\n-10\r\n-18\r\n+19\r\n+4\r\n+12\r\n-18\r\n-4\r\n-8\r\n-4\r\n-20\r\n-16\r\n-1\r\n-18\r\n+2\r\n-17\r\n-17\r\n+7\r\n-12\r\n-3\r\n+4\r\n+9\r\n+6\r\n+9\r\n+6\r\n+1\r\n-17\r\n+13\r\n+14\r\n-2\r\n+8\r\n+9\r\n-12\r\n+13\r\n-5\r\n-15\r\n+21\r\n+5\r\n+18\r\n-8\r\n-9\r\n+11\r\n+14\r\n+17\r\n+4\r\n-18\r\n+23\r\n+20\r\n-14\r\n+5\r\n-12\r\n+2\r\n+25\r\n+16\r\n-10\r\n-14\r\n+17\r\n-4\r\n-1\r\n-5\r\n+19\r\n+1\r\n+17\r\n+3\r\n-9\r\n+5\r\n+8\r\n-10\r\n+14\r\n+17\r\n-15\r\n+19\r\n+1\r\n-19\r\n+8\r\n+15\r\n-12\r\n+19\r\n-15\r\n+7\r\n+2\r\n-11\r\n-18\r\n+11\r\n+1\r\n+9\r\n+18\r\n+5\r\n+8\r\n-6\r\n-16\r\n-15\r\n-9\r\n+8\r\n-6\r\n-1\r\n-18\r\n-3\r\n-16\r\n+10\r\n-2\r\n-13\r\n+17\r\n+15\r\n-2\r\n-8\r\n-11\r\n-9\r\n-1\r\n+2\r\n+13\r\n+22\r\n-19\r\n+9\r\n+8\r\n-9\r\n-20\r\n+5\r\n-8\r\n-13\r\n-5\r\n+11\r\n+17\r\n+4\r\n+13\r\n+14\r\n+15\r\n+6\r\n+10\r\n-8\r\n+7\r\n+9\r\n+15\r\n+17\r\n+15\r\n+5\r\n+4\r\n+17\r\n-16\r\n-6\r\n+7\r\n+2\r\n+18\r\n+19\r\n-3\r\n-4\r\n-11\r\n+1\r\n-15\r\n+17\r\n-14\r\n+17\r\n+13\r\n-6\r\n-8\r\n+16\r\n+12\r\n+14\r\n-3\r\n-15\r\n-19\r\n-6\r\n-9\r\n-12\r\n-18\r\n+4\r\n-7\r\n+14\r\n-10\r\n-12\r\n+2\r\n+19\r\n+12\r\n-5\r\n+6\r\n+21\r\n-2\r\n-14\r\n+21\r\n-16\r\n+4\r\n+10\r\n+10\r\n+12\r\n+7\r\n+16\r\n-1\r\n+16\r\n+5\r\n+19\r\n+14\r\n+8\r\n+14\r\n-1\r\n-1\r\n-10\r\n-13\r\n-1\r\n-3\r\n+8\r\n+14\r\n-9\r\n+8\r\n+10\r\n-7\r\n+19\r\n-13\r\n+6\r\n+8\r\n+1\r\n+11\r\n-17\r\n+10\r\n+16\r\n+1\r\n-13\r\n+16\r\n+3\r\n+4\r\n-14\r\n+16\r\n+17\r\n-14\r\n+16\r\n+15\r\n+16\r\n+19\r\n-15\r\n+12\r\n+4\r\n-21\r\n+14\r\n-12\r\n-16\r\n-12\r\n-3\r\n+11\r\n-3\r\n+4\r\n-11\r\n+3\r\n+6\r\n+15\r\n-16\r\n-16\r\n-17\r\n-17\r\n+16\r\n-10\r\n+4\r\n+10\r\n-18\r\n+11\r\n-16\r\n-16\r\n-3\r\n-12\r\n+11\r\n-20\r\n+6\r\n-11\r\n-4\r\n-10\r\n-13\r\n-19\r\n+4\r\n-11\r\n-6\r\n+18\r\n-9\r\n+12\r\n+7\r\n+10\r\n-2\r\n-3\r\n+13\r\n-5\r\n-2\r\n-2\r\n-14\r\n+7\r\n-16\r\n-4\r\n-2\r\n-6\r\n+9\r\n+5\r\n-18\r\n-15\r\n+2\r\n+8\r\n-21\r\n+20\r\n-18\r\n-10\r\n+9\r\n+17\r\n-12\r\n-8\r\n-15\r\n+31\r\n-7\r\n-4\r\n-12\r\n-25\r\n-20\r\n-16\r\n-6\r\n+18\r\n+18\r\n-1\r\n+8\r\n+5\r\n-7\r\n-30\r\n-13\r\n-10\r\n-15\r\n-12\r\n-6\r\n+10\r\n+16\r\n-5\r\n-17\r\n-14\r\n+15\r\n+10\r\n+14\r\n-2\r\n+10\r\n+4\r\n+3\r\n-19\r\n+9\r\n+2\r\n-17\r\n+9\r\n-6\r\n+1\r\n+7\r\n+9\r\n+7\r\n-4\r\n+5\r\n+6\r\n+17\r\n-4\r\n-18\r\n+7\r\n-20\r\n-11\r\n-13\r\n-9\r\n-8\r\n+46\r\n-2\r\n+22\r\n+16\r\n-22\r\n-38\r\n-3\r\n+18\r\n+28\r\n+33\r\n+24\r\n+17\r\n+11\r\n-13\r\n-5\r\n+19\r\n-9\r\n+10\r\n+15\r\n+22\r\n+8\r\n+6\r\n-13\r\n-9\r\n-4\r\n-2\r\n+20\r\n-11\r\n-28\r\n+14\r\n+16\r\n+13\r\n+5\r\n+12\r\n+7\r\n+15\r\n+19\r\n-6\r\n+14\r\n+15\r\n-14\r\n+18\r\n+17\r\n+4\r\n-10\r\n+2\r\n-10\r\n-14\r\n+3\r\n-7\r\n+22\r\n-2\r\n+1\r\n+4\r\n-6\r\n-3\r\n+14\r\n+23\r\n-6\r\n-7\r\n-16\r\n-1\r\n+6\r\n+21\r\n+17\r\n+4\r\n-17\r\n+8\r\n-4\r\n+21\r\n-18\r\n-5\r\n+4\r\n+13\r\n+24\r\n+17\r\n+13\r\n+20\r\n+18\r\n+17\r\n+9\r\n+9\r\n+3\r\n+20\r\n+17\r\n-7\r\n+10\r\n+22\r\n-27\r\n+4\r\n-15\r\n+5\r\n-4\r\n-24\r\n+4\r\n-20\r\n-35\r\n+1\r\n-5\r\n+17\r\n-20\r\n+4\r\n+19\r\n+23\r\n-6\r\n+62\r\n+59\r\n-8\r\n-10\r\n+90\r\n+99\r\n-18\r\n-18\r\n+34\r\n-6\r\n+28\r\n-2\r\n+27\r\n-5\r\n+4\r\n+9\r\n+9\r\n-3\r\n+21\r\n+4\r\n-15\r\n+14\r\n-15\r\n+22\r\n-15\r\n-20\r\n+30\r\n+9\r\n-1\r\n+107\r\n-7\r\n-8\r\n-8\r\n+28\r\n+12\r\n-5\r\n-114\r\n+3\r\n+72\r\n-14\r\n-68\r\n-42\r\n+684\r\n+82290\r\n-12\r\n-19\r\n-3\r\n+6\r\n+4\r\n+3\r\n+16\r\n+6\r\n-15\r\n+18\r\n+3\r\n+13\r\n+10\r\n-4\r\n+10\r\n-17\r\n-7\r\n-18\r\n+2\r\n+18\r\n+13\r\n+13\r\n+6\r\n+8\r\n+18\r\n+11\r\n-13\r\n+8\r\n-7\r\n+11\r\n+4\r\n+5\r\n+7\r\n-2\r\n-16\r\n-19\r\n+14\r\n+19\r\n-6\r\n+11\r\n+8\r\n+12\r\n+7\r\n+11\r\n+5\r\n-18\r\n-4\r\n+1\r\n-7\r\n+17\r\n-8\r\n-8\r\n-5\r\n-10\r\n-16\r\n+3\r\n+15\r\n-17\r\n-9\r\n+14\r\n-9\r\n-12\r\n-4\r\n-15\r\n+7\r\n+3\r\n+17\r\n-3\r\n-19\r\n+3\r\n+6\r\n-2\r\n-17\r\n+18\r\n+4\r\n-8\r\n-8\r\n-12\r\n+10\r\n-18\r\n-20\r\n-5\r\n+15\r\n-13\r\n+20\r\n+4\r\n+9\r\n+1\r\n-7\r\n-6\r\n-5\r\n+10\r\n-9\r\n-2\r\n-21\r\n+14\r\n+21\r\n+11\r\n+4\r\n-14\r\n-5\r\n-16\r\n+1\r\n-15\r\n-10\r\n+5\r\n-1\r\n+13\r\n-11\r\n-16\r\n-16\r\n-15\r\n+13\r\n-2\r\n+5\r\n-14\r\n+4\r\n-8\r\n-3\r\n+10\r\n+5\r\n-3\r\n-16\r\n+5\r\n+16\r\n-20\r\n+11\r\n+16\r\n+6\r\n-2\r\n+6\r\n+3\r\n+21\r\n-15\r\n+1\r\n-9\r\n-2\r\n+12\r\n-3\r\n-19\r\n-9\r\n-16\r\n+11\r\n+4\r\n+18\r\n+9\r\n+26\r\n-28\r\n-12\r\n-22\r\n-18\r\n+4\r\n-7\r\n-2\r\n-5\r\n-16\r\n-19\r\n-14\r\n-3\r\n-5\r\n+14\r\n+17\r\n-4\r\n+7\r\n+4\r\n-9\r\n+14\r\n+3\r\n-11\r\n+15\r\n+10\r\n+13\r\n-10\r\n-15\r\n+18\r\n-13\r\n+1\r\n+2\r\n+3\r\n+6\r\n+10\r\n+1\r\n-18\r\n-9\r\n+3\r\n-22\r\n-13\r\n-16\r\n-15\r\n+3\r\n+6\r\n+1\r\n-12\r\n+3\r\n-17\r\n-1\r\n-16\r\n+4\r\n+9\r\n+8\r\n-18\r\n-15\r\n-18\r\n-8\r\n-13\r\n+2\r\n+15\r\n+8\r\n+3\r\n-16\r\n+2\r\n-16\r\n-12\r\n-5\r\n-3\r\n-12\r\n+19\r\n+14\r\n+20\r\n-12\r\n-11\r\n+9\r\n-10\r\n-6\r\n-10\r\n+14\r\n-17\r\n+12\r\n-14\r\n+16\r\n-6\r\n+8\r\n-10\r\n-17\r\n+4\r\n-6\r\n-1\r\n-18\r\n+2\r\n+4\r\n+9\r\n-5\r\n-13\r\n-7\r\n-16\r\n-14\r\n-8\r\n-10\r\n-4\r\n-19\r\n-19\r\n+20\r\n-4\r\n+13\r\n+19\r\n+9\r\n-17\r\n-4\r\n+11\r\n+18\r\n-1\r\n-3\r\n-16\r\n-14\r\n-17\r\n-11\r\n-6\r\n+14\r\n+6\r\n-7\r\n+5\r\n-7\r\n+14\r\n+6\r\n-3\r\n-19\r\n-17\r\n-15\r\n-12\r\n+4\r\n+5\r\n+7\r\n+5\r\n+1\r\n-8\r\n+3\r\n-11\r\n-13\r\n+4\r\n-17\r\n+15\r\n+3\r\n-11\r\n-17\r\n-4\r\n-7\r\n-1\r\n-5\r\n-7\r\n+3\r\n+5\r\n+15\r\n+21\r\n+4\r\n+17\r\n+2\r\n+12\r\n-1\r\n-1\r\n+5\r\n+17\r\n-8\r\n-3\r\n-2\r\n+7\r\n-13\r\n-19\r\n-2\r\n+3\r\n-16\r\n+11\r\n+7\r\n-14\r\n+19\r\n+3\r\n+11\r\n+14\r\n-12\r\n+2\r\n-14\r\n-18\r\n+1\r\n-27\r\n+13\r\n+3\r\n+8\r\n-14\r\n-40\r\n-1\r\n-10\r\n-4\r\n+8\r\n-31\r\n-10\r\n-3\r\n+20\r\n+6\r\n-64\r\n+7\r\n+21\r\n+52\r\n+16\r\n+32\r\n-7\r\n-3\r\n+2\r\n-91\r\n+24\r\n-94\r\n+35\r\n-117\r\n+18\r\n-19\r\n+8\r\n-12\r\n+19\r\n-12\r\n+8\r\n-24\r\n-19\r\n+14\r\n-24\r\n+7\r\n-11\r\n-19\r\n-2\r\n-16\r\n-2\r\n-6\r\n+10\r\n-12\r\n+1\r\n+5\r\n-20\r\n+10\r\n-15\r\n-3\r\n+7\r\n+12\r\n+8\r\n+11\r\n+13\r\n-11\r\n-11\r\n+13\r\n+20\r\n-18\r\n+9\r\n-16\r\n+18\r\n-10\r\n+20\r\n-13\r\n+15\r\n-18\r\n+8\r\n+14\r\n+18\r\n+4\r\n-6\r\n+16\r\n-2\r\n+9\r\n-2\r\n+33\r\n+4\r\n+6\r\n+19\r\n-11\r\n+14\r\n-13\r\n+4\r\n-2\r\n-4\r\n-21\r\n-10\r\n-82696"
}

func contains(arr []int64, val int64) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
