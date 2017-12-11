package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, _ := ioutil.ReadFile("data.txt")

	var ints []int

	for _, char := range(string(data)) {
		ints = append(ints, int(char - '0'))
	}

	ints = ints[:len(ints)-1]

	// Count identical neighbors
	count := 0
	for i := range(ints[:len(ints)-1]) {
		if ints[i] == ints[i+1] {
			count += ints[i]
		}
	}

	// test for cycle
	if ints[0] == ints[len(ints)-1] {
		count += ints[0]
	}

	// Print answer
	fmt.Println(count)
}
