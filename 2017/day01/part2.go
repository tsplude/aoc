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

	// Trim newline
	ints = ints[:len(ints)-1]

	// Count identical neighbors
	count := 0
	step := len(ints) / 2
	for i := range(ints[:len(ints) / 2]) {
		if ints[i] == ints[i + step] {
			count += ints[i] * 2
		}
	}

	// Print answer
	fmt.Println(count)
}
